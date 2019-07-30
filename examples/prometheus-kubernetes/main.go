package main

import (
	"github.com/bwplotka/mimic"
	"github.com/bwplotka/mimic/encoding"
	"github.com/prometheus/common/model"
	rbacv1beta1"k8s.io/api/rbac/v1beta1"
	corev1 "k8s.io/api/core/v1"
	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"

	amConfig "github.com/prometheus/alertmanager/config"
	
)

const (
	namespace = "default"
	alertManagerPort = 9093
)

func main() {
	generator := mimic.New()
	defer generator.Generate()
	
	// Alertmanager
	
	alertmanagerCrb := rbacv1beta1.ClusterRoleBinding{
		ObjectMeta: metav1.ObjectMeta{
			Name: "release-name",
			Labels: map[string]string{
				"app": "prometheus",
				"component": "alertmanager",
			},
		},
		RoleRef: rbacv1beta1.RoleRef{
			APIGroup: "rbac.authorization.k8s.io",
			Kind: "ClusterRole",
			Name: "cluster-admin",
		},
		Subjects: []rbacv1beta1.Subject{
			{
				Kind: rbacv1beta1.ServiceAccountKind,
				Name: "release-name-prometheus-alertmanager",
				Namespace: namespace,
			},
		},
	}
	
	generator.Add("alertmanager-clusterrolebinding.yaml", encoding.YAML(alertmanagerCrb))

	fiveMinutes, err := model.ParseDuration("5m")
	if err != nil {
		panic(err)
	}
	tenSeconds, err := model.ParseDuration("10s")
	if err != nil {
		panic(err)
	}
	threeHours, err := model.ParseDuration("3h")
	if err != nil {
		panic(err)
	}
	
	alertManagerConfig := amConfig.Config{
		Receivers: []*amConfig.Receiver{
			{
				Name: "default-receiver",
			},
		},
		Route: &amConfig.Route{
			GroupInterval: &fiveMinutes,
			GroupWait: &tenSeconds,
			Receiver: "default-receiver",
			RepeatInterval: &threeHours,
			
		},
	}
	
	alertmanagerConfigMap := corev1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name: "release-name-prometheus-alertmanager",
			Labels: map[string]string{
				"app": "prometheus",
				"component": "alertmanager",
			},
		},
		Data: map[string]string{
			"alertmanager.yml": alertManagerConfig.String(),
		},
	}

	generator.Add("alertmanager-configmap.yaml", encoding.YAML(alertmanagerConfigMap))
	
	int32One := int32(1)
	
	alertManagerDeployment := appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: "release-name-prometheus-alertmanager",
			Labels: map[string]string{
				"app": "prometheus",
				"component": "alertmanager",
			},
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: &int32One,
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": "prometheus",
						"component": "alertmanager",
					},
				},
				Spec: corev1.PodSpec{
					ServiceAccountName: "prometheus-alertmanager",
					Containers: []corev1.Container{
						{
							Name: "prometheus-alertmanager",
							Image: "prom/alertmanager:v0.14.0",
							ImagePullPolicy: "IfNotPresent",
							Args: []string{
								"--config.file=/etc/config/alertmanager.yml",
								"--storage.path=/data",
								"--web.external-url=/",
							},
							Ports: []corev1.ContainerPort{
								{
									ContainerPort: alertManagerPort,
								},
							},
							ReadinessProbe: &corev1.Probe{
								Handler: corev1.Handler{
									HTTPGet: &corev1.HTTPGetAction{
										Path: "/#/status",
										Port: intstr.FromInt(alertManagerPort),
									},
								},
								InitialDelaySeconds: 30,
								TimeoutSeconds: 30,
							},
							VolumeMounts: []corev1.VolumeMount{
								{
									Name: "config-volume",
									MountPath: "/etc/config",
								},
								{
									Name: "storage-volume",
									MountPath: "/data",
									SubPath: "",
								},
							},
						},
						{
							Name: "prometheus-alertmanager-configmap-reload",
							Image: "jimmidyson/configmap-reload:v0.1",
							ImagePullPolicy: "IfNotPresent",
							Args: []string{
								"--volume-dir=/etc/config",
								"--webhook-url=http://localhost:9093/-/reload",
							},
							VolumeMounts: []corev1.VolumeMount{
								{
									Name: "config-volume",
									MountPath: "/etc/config",
									ReadOnly: true,
								},
							},
						},
					},
				},
			},
		},
	}

	generator.Add("alertmanager-deployment.yaml", encoding.YAML(alertManagerDeployment))
	
	// Kube-state-metrics
	
	// Node-exporter
	
	// Pushgateway
	
	// Server
	
}