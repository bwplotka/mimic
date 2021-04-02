package main

import (
	"fmt"
	"io/ioutil"
	"path"

	"github.com/prometheus/common/model"

	"github.com/bwplotka/mimic/lib/abstr/kubernetes/volumes"
	"github.com/go-openapi/swag"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/apimachinery/pkg/util/intstr"

	"github.com/bwplotka/mimic"
	"github.com/bwplotka/mimic/encoding"
	"github.com/bwplotka/mimic/lib/schemas/prometheus"
)

const (
	selectorName = "app"
)

/*

Test procedure:

* Generate YAMLs from definitions:
  * `go run github.com/bwplotka/mimic/projects/prom-remote-read-bench generate`

* Apply baseline:
  * `kubectl apply -f gen/prom-rr-test.yaml`

* Forward gRPC sidecar port:
  * `kubectl port-forward pod/prom-rr-test-0 1234:19090`

* Perform tests using test.sh (modifying parameters in script itself - heavy queries!)
  * This performs heavy queries against Thanos gRPC Store.Series of sidecar which will proxy
requests as remote read to Prometheus
  * `bash ./projects/prom-remote-read-bench/test.sh localhost:1234`

*/

func main() {
	generator := mimic.New()

	// Make sure to generate at the very end.
	defer generator.Generate()

	// Generate resources for remote read tests.

	// Baseline.
	genRRTestPrometheus(
		generator,
		"prom-rr-test",
		"v2.11.0-rc.0-clear",
		"v0.5.0",
	)

	// Streamed.
	genRRTestPrometheus(
		generator,
		"prom-rr-test-streamed",
		"v2.11.0-rc.0-rr-streaming",
		"v0.5.0-rr-streamed2",
	)
}

func genRRTestPrometheus(generator *mimic.Generator, name string, promVersion string, thanosVersion string) {
	const (
		replicas = 1

		configVolumeName  = "prometheus-config"
		configVolumeMount = "/etc/prometheus"
		sharedDataPath    = "/data-shared"

		namespace = "rr-test"

		httpPort        = 9090
		httpSidecarPort = 19190
		grpcSidecarPort = 19090
		blockgenImage   = "improbable/blockgen:master-894c9481c4"
		// Generate 10k series.
		blockgenInput = `[{
  "type": "gauge",
  "jitter": 20,
  "max": 200000000,
  "min": 100000000,
  "result": {"multiplier":10000,"resultType":"vector","result":[{"metric":{"__name__":"kube_pod_container_resource_limits_memory_bytes","cluster":"eu1","container":"addon-resizer","instance":"172.17.0.9:8080","job":"kube-state-metrics","namespace":"kube-system","node":"minikube","pod":"kube-state-metrics-68f6cc566c-vp566"}}]}
}]`
	)
	var (
		promDataPath    = path.Join(sharedDataPath, "prometheus")
		prometheusImage = fmt.Sprintf("bplotka/prometheus:%s", promVersion)
		thanosImage     = fmt.Sprintf("improbable/thanos:%s", thanosVersion)
	)

	// Empty configuration, we don't need any scrape.
	cfgBytes, err := ioutil.ReadAll(encoding.YAML(prometheus.Config{
		GlobalConfig: prometheus.GlobalConfig{
			ExternalLabels: map[model.LabelName]model.LabelValue{
				"replica": "0",
			},
		},
	}))
	mimic.PanicIfErr(err)

	promConfigAndMount := volumes.ConfigAndMount{
		ObjectMeta: metav1.ObjectMeta{
			Name:      configVolumeName,
			Namespace: namespace,
			Labels: map[string]string{
				selectorName: name,
			},
		},
		VolumeMount: corev1.VolumeMount{
			Name:      configVolumeName,
			MountPath: configVolumeMount,
		},
		Data: map[string]string{
			"prometheus.yaml": string(cfgBytes),
		},
	}

	sharedVM := volumes.VolumeAndMount{
		VolumeMount: corev1.VolumeMount{
			Name:      name,
			MountPath: sharedDataPath,
		},
	}

	srv := corev1.Service{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Service",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
			Labels: map[string]string{
				selectorName: name,
			},
		},
		Spec: corev1.ServiceSpec{
			Type:      corev1.ServiceTypeClusterIP,
			ClusterIP: "None",
			Selector: map[string]string{
				selectorName: name,
			},
			Ports: []corev1.ServicePort{
				{
					Name:       "http",
					Port:       httpPort,
					TargetPort: intstr.FromInt(httpPort),
				},
				{
					Name:       "grpc-sidecar",
					Port:       grpcSidecarPort,
					TargetPort: intstr.FromInt(grpcSidecarPort),
				},
				{
					Name:       "http-sidecar",
					Port:       httpSidecarPort,
					TargetPort: intstr.FromInt(httpSidecarPort),
				},
			},
		},
	}

	blockgenInitContainer := corev1.Container{
		Name:    "blockgen",
		Image:   blockgenImage,
		Command: []string{"/bin/blockgen"},
		Args: []string{
			"synthetic",
			fmt.Sprintf("--input=%s", blockgenInput),
			fmt.Sprintf("--output-dir=%s", promDataPath),
			"--retention=24h",
		},
		VolumeMounts: []corev1.VolumeMount{sharedVM.VolumeMount},
		Resources: corev1.ResourceRequirements{
			Requests: corev1.ResourceList{
				corev1.ResourceCPU:    resource.MustParse("1"),
				corev1.ResourceMemory: resource.MustParse("10Gi"),
			},
			Limits: corev1.ResourceList{
				corev1.ResourceCPU:    resource.MustParse("1"),
				corev1.ResourceMemory: resource.MustParse("10Gi"),
			},
		},
	}

	prometheusContainer := corev1.Container{
		Name:  "prometheus",
		Image: prometheusImage,
		Args: []string{
			fmt.Sprintf("--config.file=%v/prometheus.yaml", configVolumeMount),
			"--log.level=info",
			// Unlimited RR.
			"--storage.remote.read-concurrent-limit=99999",
			"--storage.remote.read-sample-limit=9999999999999999",
			fmt.Sprintf("--storage.tsdb.path=%s", promDataPath),
			"--storage.tsdb.min-block-duration=2h",
			// Avoid compaction for less moving parts in results.
			"--storage.tsdb.max-block-duration=2h",
			"--storage.tsdb.retention.time=2d",
			"--web.enable-lifecycle",
			"--web.enable-admin-api",
		},
		Env: []corev1.EnvVar{
			{Name: "HOSTNAME", ValueFrom: &corev1.EnvVarSource{
				FieldRef: &corev1.ObjectFieldSelector{
					FieldPath: "metadata.name",
				},
			}},
			//{Name: "GODEBUG", Value:"madvdontneed=1"},
		},
		ImagePullPolicy: corev1.PullAlways,
		ReadinessProbe: &corev1.Probe{
			Handler: corev1.Handler{
				HTTPGet: &corev1.HTTPGetAction{
					Port: intstr.FromInt(int(httpPort)),
					Path: "-/ready",
				},
			},
			SuccessThreshold: 3,
		},
		Ports: []corev1.ContainerPort{
			{
				Name:          "http",
				ContainerPort: httpPort,
			},
		},
		VolumeMounts: volumes.VolumesAndMounts{promConfigAndMount.VolumeAndMount(), sharedVM}.VolumeMounts(),
		SecurityContext: &corev1.SecurityContext{
			RunAsNonRoot: swag.Bool(false),
			RunAsUser:    swag.Int64(1000),
		},
		Resources: corev1.ResourceRequirements{
			Requests: corev1.ResourceList{
				corev1.ResourceCPU:    resource.MustParse("1"),
				corev1.ResourceMemory: resource.MustParse("10Gi"),
			},
			Limits: corev1.ResourceList{
				corev1.ResourceCPU:    resource.MustParse("1"),
				corev1.ResourceMemory: resource.MustParse("10Gi"),
			},
		},
	}

	thanosSidecarContainer := corev1.Container{
		Name:            "thanos",
		Image:           thanosImage,
		Command:         []string{"thanos"},
		ImagePullPolicy: corev1.PullAlways,
		Args: []string{
			"sidecar",
			"--log.level=debug",
			"--debug.name=$(POD_NAME)",
			fmt.Sprintf("--http-address=0.0.0.0:%d", httpSidecarPort),
			fmt.Sprintf("--grpc-address=0.0.0.0:%d", grpcSidecarPort),
			fmt.Sprintf("--prometheus.url=http://localhost:%d", httpPort),
			fmt.Sprintf("--tsdb.path=%s", promDataPath),
		},
		Env: []corev1.EnvVar{
			{Name: "POD_NAME", ValueFrom: &corev1.EnvVarSource{
				FieldRef: &corev1.ObjectFieldSelector{
					FieldPath: "metadata.name",
				},
			}},
			//{Name: "GODEBUG", Value:"madvdontneed=1"},
		},
		Ports: []corev1.ContainerPort{
			{
				Name:          "m-sidecar",
				ContainerPort: httpSidecarPort,
			},
			{
				Name:          "grpc-sidecar",
				ContainerPort: grpcSidecarPort,
			},
		},
		ReadinessProbe: &corev1.Probe{
			Handler: corev1.Handler{
				HTTPGet: &corev1.HTTPGetAction{
					Port: intstr.FromInt(int(httpSidecarPort)),
					Path: "metrics",
				},
			},
		},
		VolumeMounts: volumes.VolumesAndMounts{sharedVM}.VolumeMounts(),
		Resources: corev1.ResourceRequirements{
			Requests: corev1.ResourceList{
				corev1.ResourceCPU:    resource.MustParse("1"),
				corev1.ResourceMemory: resource.MustParse("4Gi"),
			},
			Limits: corev1.ResourceList{
				corev1.ResourceCPU:    resource.MustParse("1"),
				corev1.ResourceMemory: resource.MustParse("4Gi"),
			},
		},
	}

	set := appsv1.StatefulSet{
		TypeMeta: metav1.TypeMeta{
			Kind:       "StatefulSet",
			APIVersion: "apps/v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
			Labels: map[string]string{
				selectorName: name,
			},
		},
		Spec: appsv1.StatefulSetSpec{
			Replicas:    swag.Int32(replicas),
			ServiceName: name,
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						selectorName: name,
						"version":    fmt.Sprintf("prometheus%s_thanos%s", promVersion, thanosVersion),
					},
				},
				Spec: corev1.PodSpec{
					InitContainers: []corev1.Container{blockgenInitContainer},
					Containers:     []corev1.Container{prometheusContainer, thanosSidecarContainer},
					Volumes:        volumes.VolumesAndMounts{promConfigAndMount.VolumeAndMount(), sharedVM}.Volumes(),
				},
			},
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					selectorName: name,
				},
			},
		},
	}

	generator.Add(name+".yaml", encoding.GhodssYAML(set, srv, promConfigAndMount.ConfigMap()))
}
