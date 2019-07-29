package main

import (
	"fmt"
	"time"

	"github.com/bwplotka/mimic"
	"github.com/bwplotka/mimic/encoding"
	"github.com/bwplotka/mimic/providers/dockercompose"
	"github.com/bwplotka/mimic/providers/prometheus"
	sdconfig "github.com/bwplotka/mimic/providers/prometheus/discovery/config"
	"github.com/bwplotka/mimic/providers/prometheus/discovery/dns"
	"github.com/bwplotka/mimic/providers/prometheus/discovery/targetgroup"
	"github.com/prometheus/common/config"
	"github.com/prometheus/common/model"
	"gopkg.in/alecthomas/kingpin.v2"
)

// This is not the best, but the simplest solution for secrets. See: README.md#Important: Guide & best practices
type Secrets struct {
	Alertmanager config.BasicAuth `yaml:"alertmanager"`
}

func main() {
	var secretFile *string
	generator := mimic.New(
		func(cmd *kingpin.CmdClause) {
			secretFile = cmd.Flag("secret-file", "YAML file with secrets").Required().String()
		},
	)

	// Make sure to generate at the very end.
	defer generator.Generate()

	var secrets Secrets
	mimic.UnmarshalSecretFile(*secretFile, &secrets)

	// Start generating stuff.
	genMyMonAll(generator, secrets)
}

func genMyMonAll(generator *mimic.Generator, secrets Secrets) {
	for _, env := range Environments {
		generator := generator.With(env.Name)
		for _, cl := range ClustersByEnv[env] {
			generator := generator.With(cl.Name)

			genMyMonDockerCompose(generator.With("deploy"), cl, secrets)
		}
	}
}

func genMyMonPrometheusConfig(
	generator *mimic.Generator,
	cl *Cluster,
	nodeExporterPort int,
	prometheusPort int,
	dockerdPort int,
	secrets Secrets,
) {
	promConfig := prometheus.Config{
		GlobalConfig: prometheus.GlobalConfig{
			ScrapeInterval: model.Duration(15 * time.Second),
			ScrapeTimeout:  model.Duration(15 * time.Second),
			ExternalLabels: model.LabelSet{
				"cluster": model.LabelValue(cl.Name),
				"env":     model.LabelValue(cl.Environment.Name),
				"replica": "mon-0",
			},
		},
		AlertingConfig: prometheus.AlertingConfig{
			AlertmanagerConfigs: []*prometheus.AlertmanagerConfig{
				{
					HTTPClientConfig: config.HTTPClientConfig{
						BasicAuth: &secrets.Alertmanager,
					},
				},
			},
		},
		// TODO: alerts: RuleFiles: []
		ScrapeConfigs: []*prometheus.ScrapeConfig{
			{
				JobName: "mon",
				ServiceDiscoveryConfig: sdconfig.ServiceDiscoveryConfig{
					StaticConfigs: []*targetgroup.Group{
						{Targets: []model.LabelSet{{model.AddressLabel: model.LabelValue(fmt.Sprintf("localhost:%d", prometheusPort))}}},
					},
				},
			},
			{
				JobName: "node",
				ServiceDiscoveryConfig: sdconfig.ServiceDiscoveryConfig{
					DNSSDConfigs: []*dns.SDConfig{
						{
							Names: []string{"tasks.node-exporter"},
							Type:  "A",
							Port:  nodeExporterPort,
						},
					},
				},
			},
			{
				JobName: "docker",
				ServiceDiscoveryConfig: sdconfig.ServiceDiscoveryConfig{
					StaticConfigs: []*targetgroup.Group{
						{Targets: []model.LabelSet{{model.AddressLabel: model.LabelValue(fmt.Sprintf("localhost:%d", dockerdPort))}}},
					},
				},
			},
		},
	}
	generator.Add("prometheus.yaml", encoding.YAML(promConfig))
}

func genMyMonDockerCompose(generator *mimic.Generator, cl *Cluster, secrets Secrets) {
	const (
		prometheusDataVolume       = "prometheus-data"
		prometheusDockerVolumePath = "/docker-volumes/prometheus-data"

		prometheusConfigVolume           = "prometheus-config"
		prometheusDockerVolumeConfigPath = "/docker-volumes/prometheus-config"

		monitoringNet = "monitor-net"

		promPort    = 9090
		nodeExpPort = 9100
		dockerdPort = 9323
	)

	var (
		promReplicas = uint64(1)
	)
	//Top-level object must be a mapping

	genMyMonPrometheusConfig(generator.With("configs"), cl, promPort, nodeExpPort, dockerdPort, secrets)

	// TODO(bwplotka): Add envoy, alertmanager, make sure docker is monitored as well etc.
	dpl := dockercompose.Config{
		Volumes: []dockercompose.VolumeConfig{
			{Name: prometheusDataVolume, Driver: "local"},
			{Name: prometheusConfigVolume, Driver: "local"},
		},
		Networks: dockercompose.NetworkConfigs{{Name: monitoringNet, Driver: "bridge"}},
		Services: dockercompose.Services{
			{
				Name:  "mon",
				Image: "quay.io/prometheus/prometheus:v2.6.1",
				Command: dockercompose.ShellCommand{
					"--config.file=" + prometheusDockerVolumeConfigPath + "prometheus.yml",
					"--storage.tsdb.retention.time=10d",
					"--storage.tsdb.path=" + prometheusDockerVolumePath,
					"--web.enable.lifecycle",
					// promPort,
				},
				Restart: dockercompose.UnlessStopped_RestartServiceConfig,
				Ports: []dockercompose.ServicePortConfig{
					{
						Published: promPort,
						Target:    promPort,
					},
				},
				Deploy: &dockercompose.DeployConfig{
					Replicas: &promReplicas,
				},
				Volumes: []dockercompose.ServiceVolumeConfig{
					{
						Type:        dockercompose.Volume_ServiceVolumeType,
						Source:      prometheusDataVolume,
						Target:      prometheusDockerVolumePath,
						ReadOnly:    true,
						Consistency: dockercompose.Consistent_ServiceVolumeConsistency,
						Volume: &dockercompose.ServiceVolumeVolume{
							NoCopy: true,
						},
					},
					{
						Type:        dockercompose.Volume_ServiceVolumeType,
						Source:      prometheusConfigVolume,
						Target:      prometheusDockerVolumeConfigPath,
						ReadOnly:    true,
						Consistency: dockercompose.Consistent_ServiceVolumeConsistency,
						Volume: &dockercompose.ServiceVolumeVolume{
							NoCopy: true,
						},
					},
				},
				Networks: map[string]*dockercompose.ServiceNetworkConfig{
					monitoringNet: {},
				},
			},
			{
				Name:  "node_exporter",
				Image: "prom/node-exporter:v0.17.0",
				Command: dockercompose.ShellCommand{
					"--path.procfs=/host/proc",
					"--path.rootfs=/rootfs",
					"--path.sysfs=/host/sys",
					"--collector.filesystem.ignored-mount-points=^/(sys|proc|dev|host|etc)($$|/)",
					// nodeExpPort
				},
				User:       "root",
				Privileged: true,
				Volumes: []dockercompose.ServiceVolumeConfig{
					{
						Type:     dockercompose.Bind_ServiceVolumeType,
						Source:   "/proc",
						Target:   "/host/proc",
						ReadOnly: true,
					},
					{
						Type:     dockercompose.Bind_ServiceVolumeType,
						Source:   "/sys",
						Target:   "/host/sys",
						ReadOnly: true,
					},
					{
						Type:     dockercompose.Bind_ServiceVolumeType,
						Source:   "/",
						Target:   "/rootfs",
						ReadOnly: true,
					},
				},
				Networks: map[string]*dockercompose.ServiceNetworkConfig{
					monitoringNet: {},
				},
				Deploy: &dockercompose.DeployConfig{
					Mode: "global",
				},
			},
			// Add  dockerd-exporter:
			//    image: stefanprodan/caddy
			//    networks:
			//      - net
			//    environment:
			//      - DOCKER_GWBRIDGE_IP=172.18.0.1
			//    configs:
			//      - source: dockerd_config
			//        target: /etc/caddy/Caddyfile
			//    deploy:
			//      mode: global
			//      resources:
			//        limits:
			//          memory: 128M
			//        reservations:
			//          memory: 64M
			// when I will have more than 2 docker node!
		},
	}

	generator.Add("mon-compose.yaml", encoding.YAML(dpl))
}
