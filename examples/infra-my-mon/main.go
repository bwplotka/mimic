package main

import (
	"github.com/bwplotka/gocodeit"
	"github.com/bwplotka/gocodeit/encoding"
	"github.com/bwplotka/gocodeit/providers/dockercompose"
	"github.com/bwplotka/gocodeit/providers/prometheus"
	"github.com/prometheus/common/config"
	"gopkg.in/alecthomas/kingpin.v2"
)

// This is not the best, but the simplest solution for secrets. See: README.md#Important: Guide & best practices
type Secrets struct {
	Alertmanager config.BasicAuth `yaml:"alertmanager"`
}

func main() {
	var secretFile *string
	gci := gocodeit.New(
		func(cmd *kingpin.CmdClause) {
			secretFile = cmd.Flag("secret-file", "YAML file with secrets").Required().String()
		},
	)

	// Make sure to generate at the very end.
	defer gci.Generate()

	var secrets Secrets
	gocodeit.UnmarshalSecretFile(*secretFile, &secrets)

	// Start generating stuff.
	genMyMonAll(gci, secrets)
}

func genMyMonAll(gci *gocodeit.Generator, secrets Secrets) {
	for _, env := range Environments {
		gci := gci.With(env.Name)
		for _, cl := range ClustersByEnv[env] {
			gci := gci.With(cl.Name)

			genMyMonDockerCompose(gci.With("deploy"))
			genMyMonPrometheusConfig(gci.With("configs"), secrets)
		}
	}
}

func genMyMonPrometheusConfig(gci *gocodeit.Generator, secrets Secrets) {
	cfg := prometheus.Config{
		AlertingConfig: prometheus.AlertingConfig{
			AlertmanagerConfigs: []*prometheus.AlertmanagerConfig{
				{
					HTTPClientConfig: config.HTTPClientConfig{
						BasicAuth: &secrets.Alertmanager,
					},
				},
			},
		},
		ScrapeConfigs: []*prometheus.ScrapeConfig{},
	}
	gci.Add("prometheus.yaml", encoding.YAML(cfg))
}

func genMyMonDockerCompose(gci *gocodeit.Generator) {
	const (
		prometheusDataVolume       = "prometheus-data"
		prometheusDockerVolumePath = "/docker-volumes/prometheus-data"

		prometheusConfigVolume           = "prometheus-config"
		prometheusDockerVolumeConfigPath = "/docker-volumes/prometheus-config"

		monitoringNet = "monitor-net"
	)

	// TODO(bwplotka): Add envoy, alertmanager, make sure docker is monitored as well.
	dpl := dockercompose.Config{
		Volumes: []dockercompose.VolumeConfig{
			{Name: prometheusDataVolume, Driver: "local"},
			{Name: prometheusConfigVolume, Driver: "local"},
		},

		Networks: dockercompose.NetworkConfigs{{Name: monitoringNet, Driver: "bridge"}},

		Services: dockercompose.Services{
			{
				Name:  "prometheus",
				Image: "quay.io/prometheus/prometheus:v2.6.1",
				Command: dockercompose.ShellCommand{
					"--config.file=" + prometheusDockerVolumeConfigPath + "prometheus.yml",
					"--storage.tsdb.retention.time=10d",
					"--storage.tsdb.path=" + prometheusDockerVolumePath,
					"--web.enable.lifecycle",
				},
				Restart: dockercompose.UnlessStopped_RestartServiceConfig,
				Ports: []dockercompose.ServicePortConfig{
					{
						Published: 9090,
						Target:    9090,
					},
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
			},
		},
	}

	gci.Add("mon-compose.yaml", encoding.YAML(dpl))
}
