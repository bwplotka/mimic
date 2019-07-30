package main

import (
	"github.com/bwplotka/mimic"
	"github.com/bwplotka/mimic/encoding"
	"github.com/prometheus/common/model"
	rbacv1beta1"k8s.io/api/rbac/v1beta1"
	corev1 "k8s.io/api/core/v1"
	appsv1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"

	amConfig "github.com/prometheus/alertmanager/config"
	
)

const (
	namespace = "default"
	alertManagerPort = 9093
	// This constant is not seemingly available in any of the k8s libraries
	imagePullPolicyIfNotPresent = imagePullPolicyIfNotPresent
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
							ImagePullPolicy: imagePullPolicyIfNotPresent,
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
							ImagePullPolicy: imagePullPolicyIfNotPresent,
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
	
	alertManagerPvc := corev1.PersistentVolumeClaim{
		ObjectMeta: metav1.ObjectMeta{
			Name: "release-name-prometheus-alertmanager",
			Labels: map[string]string{
				"app": "prometheus",
				"component": "alertmanager",
			},
		},
		Spec: corev1.PersistentVolumeClaimSpec{
			AccessModes: []corev1.PersistentVolumeAccessMode{
				corev1.ReadWriteOnce,
			},
			Resources: corev1.ResourceRequirements{
				Requests: corev1.ResourceList{
					corev1.ResourceStorage: resource.MustParse("2Gi"),
				},
			},
		},
	}

	generator.Add("alertmanager-pvc.yaml", encoding.YAML(alertManagerPvc))
	
	alertManagerService := corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name: "release-name-prometheus-alertmanager",
			Labels: map[string]string{
				"app": "prometheus",
				"component": "alertmanager",
			},
		},
		Spec: corev1.ServiceSpec{
			Ports: []corev1.ServicePort{
				{
					Name: "http",
					Port: 80,
					Protocol: "TCP",
					TargetPort: intstr.FromInt(alertManagerPort),
				},
			},
			Selector: map[string]string{
				"app": "prometheus",
				"component": "alertmanager",
			},
			Type: corev1.ServiceTypeClusterIP,
		},
	}

	generator.Add("alertmanager-service.yaml", encoding.YAML(alertManagerService))
	
	alertManagerServiceAccount := corev1.ServiceAccount{
		ObjectMeta: metav1.ObjectMeta{
			Name: "release-name-prometheus-alertmanager",
			Labels: map[string]string{
				"app": "prometheus",
				"component": "alertmanager",
			},
		},
	}

	generator.Add("alertmanager-serviceaccount.yaml", encoding.YAML(alertManagerServiceAccount))
	
	// Kube-state-metrics
	
	kubeStateMetricsClusterRole := rbacv1beta1.ClusterRole{
		ObjectMeta: metav1.ObjectMeta{
			Name: "release-name-prometheus-kube-state-metrics",
			Labels: map[string]string{
				"app": "prometheus",
				"component": "alertmanager",
			},
		},
		Rules: []rbacv1beta1.PolicyRule{
			{
				APIGroups: []string{""},
				Resources: []string{
					"namespaces",
					"nodes",
					"persistentvolumeclaims",
					"pods",
					"services",
					"resourcequotas",
					"replicationcontrollers",
					"limitranges",
					"persistentvolumeclaims",
					"persistentvolumes",
					"endpoints",
				},
				Verbs: []string{
					"list",
					"watch",
				},
			},
			{
				APIGroups: []string{"extensions"},
				Resources: []string{
					"daemonsets",
					"deployments",
					"replicasets",
				},
				Verbs: []string{
					"list",
					"watch",
				},
			},
			{
				APIGroups: []string{"apps"},
				Resources: []string{
					"statefulsets",
				},
				Verbs: []string{
					"list",
					"watch",
					"get",
				},
			},
			{
				APIGroups: []string{"batch"},
				Resources: []string{
					"cronjobs",
					"jobs",
				},
				Verbs: []string{
					"list",
					"watch",
				},
			},
			{
				APIGroups: []string{"autoscaling"},
				Resources: []string{
					"horizontalpodautoscalers",
				},
				Verbs: []string{
					"list",
					"watch",
				},
			},
		},
	}

	generator.Add("kube-state-metrics-clusterrole.yaml", encoding.YAML(kubeStateMetricsClusterRole))

	kubeStateMetricsClusterRoleBinding := rbacv1beta1.ClusterRoleBinding{
		ObjectMeta: metav1.ObjectMeta{
			Name: "release-name-prometheus-kube-state-metrics",
			Labels: map[string]string{
				"app": "prometheus",
				"component": "kube-state-metrics",
			},
		},
		RoleRef: rbacv1beta1.RoleRef{
			APIGroup: "rbac.authorization.k8s.io",
			Kind: "ClusterRole",
			Name: "release-name-prometheus-kube-state-metrics",
		},
		Subjects: []rbacv1beta1.Subject{
			{
				Kind: rbacv1beta1.ServiceAccountKind,
				Name: "release-name-prometheus-kube-state-metrics",
				Namespace: namespace,
			},
		},
	}

	generator.Add("kube-state-metrics-clusterrolebinding.yaml", encoding.YAML(kubeStateMetricsClusterRoleBinding))
	
	kubeStateMetricsDeployment := appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: "release-name-prometheus-kube-state-metrics",
			Labels: map[string]string{
				"app": "prometheus",
				"component": "kube-state-metrics",
			},
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: &int32One,
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": "prometheus",
						"component": "kube-state-metrics",
					},
				},
				Spec: corev1.PodSpec{
					ServiceAccountName: "release-name-prometheus-kube-state-metrics",
					Containers: []corev1.Container{
						{
							Name:            "prometheus-kube-state-metrics",
							Image:           "k8s.gcr.io/kube-state-metrics:v1.2.0",
							ImagePullPolicy: imagePullPolicyIfNotPresent,
							Ports: []corev1.ContainerPort{
								{
									ContainerPort: 8080,
									Name:          "metrics",
								},
							},
						},
					},
				},
			},
		},
	}

	generator.Add("kube-state-metrics-deployment.yaml", encoding.YAML(kubeStateMetricsDeployment))

	kubeStateMetricsService := corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name: "release-name-prometheus-kube-state-metrics",
			Labels: map[string]string{
				"app": "prometheus",
				"component": "kube-state-metrics",
			},
		},
		Spec: corev1.ServiceSpec{
			Ports: []corev1.ServicePort{
				{
					Name: "http",
					Port: 80,
					Protocol: "TCP",
					TargetPort: intstr.FromInt(8080),
				},
			},
			Selector: map[string]string{
				"app": "prometheus",
				"component": "kube-state-metrics",
			},
			Type: corev1.ServiceTypeClusterIP,
		},
	}

	generator.Add("kube-state-metrics-service.yaml", encoding.YAML(kubeStateMetricsService))

	kubeStateMetricsServiceAccount := corev1.ServiceAccount{
		ObjectMeta: metav1.ObjectMeta{
			Name: "release-name-prometheus-kube-state-metrics",
			Labels: map[string]string{
				"app": "prometheus",
				"component": "kube-state-metrics",
			},
		},
	}

	generator.Add("kube-state-metrics-serviceaccount.yaml", encoding.YAML(kubeStateMetricsServiceAccount))

	// Node-exporter
	
	nodeExporterClusterRoleBinding := rbacv1beta1.ClusterRoleBinding{
		ObjectMeta: metav1.ObjectMeta{
			Name: "release-name",
			Labels: map[string]string{
				"app": "prometheus",
				"component": "node-exporter",
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
				Name: "release-name-prometheus-node-exporter",
				Namespace: namespace,
			},
		},
	}

	generator.Add("node-exporter-clusterrolebinding.yaml", encoding.YAML(nodeExporterClusterRoleBinding))
	
	nodeExporterDaemonSet := appsv1.DaemonSet{
		ObjectMeta: metav1.ObjectMeta{
			Name: "release-name",
			Labels: map[string]string{
				"app": "prometheus",
				"component": "node-exporter",
			},
		},
		Spec:appsv1.DaemonSetSpec{
			UpdateStrategy:appsv1.DaemonSetUpdateStrategy{
				Type:appsv1.OnDeleteDaemonSetStrategyType,
			},
			Template:corev1.PodTemplateSpec{
				Spec:corev1.PodSpec{
					ServiceAccountName: "release-name-prometheus-node-exporter",
					Containers: []corev1.Container{
						{
							Name: "release-name-prometheus-node-exporter",
							Image: "prom/node-exporter:v0.15.2",
							ImagePullPolicy: imagePullPolicyIfNotPresent,
							Args: []string{
								"--path.procfs=/host/proc",
								"--path.sysfs=/host/sys",
							},
							Ports: []corev1.ContainerPort{
								{
									Name: "metrics",
									ContainerPort: 9100,
									HostPort: 9100,
								},
							},
							VolumeMounts: []corev1.VolumeMount{
								{
									Name: "proc",
									MountPath: "/host/proc",
									ReadOnly: true,
								},
								{
									Name: "sys",
									MountPath: "/host/sys",
									ReadOnly: true,
								},
							},
						},
					},
					HostNetwork:true,
					HostPID:true,
					Volumes:[]corev1.Volume{
						{
							Name: "proc",
							VolumeSource: corev1.VolumeSource{
								HostPath: &corev1.HostPathVolumeSource{
									Path: "/proc",
								},
							},
						},
						{
							Name: "sys",
							VolumeSource: corev1.VolumeSource{
								HostPath: &corev1.HostPathVolumeSource{
									Path: "/sys",
								},
							},
						},
					},
				},
			},
		},
	}

	generator.Add("node-exporter-daemonset.yaml", encoding.YAML(nodeExporterDaemonSet))

	nodeExporterService := corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name: "release-name-prometheus-node-exporter",
			Labels: map[string]string{
				"app": "prometheus",
				"component": "node-exporter",
			},
		},
		Spec: corev1.ServiceSpec{
			Ports: []corev1.ServicePort{
				{
					Name: "metrics",
					Port: 9100,
					Protocol: "TCP",
					TargetPort: intstr.FromInt(9100),
				},
			},
			Selector: map[string]string{
				"app": "prometheus",
				"component": "node-exporter",
			},
			Type: corev1.ServiceTypeClusterIP,
			ClusterIP: corev1.ClusterIPNone,
		},
	}

	generator.Add("node-exporter-service.yaml", encoding.YAML(nodeExporterService))

	nodeExporterServiceAccount := corev1.ServiceAccount{
		ObjectMeta: metav1.ObjectMeta{
			Name: "release-name-prometheus-node-exporter",
			Labels: map[string]string{
				"app": "prometheus",
				"component": "node-exporter",
			},
		},
	}

	generator.Add("node-exporter-serviceaccount.yaml", encoding.YAML(nodeExporterServiceAccount))
	
	// Pushgateway

	pushgatewayDeployment := appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: "release-name-prometheus-kube-state-metrics",
			Labels: map[string]string{
				"app": "prometheus",
				"component": "pushgateway",
			},
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: &int32One,
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": "prometheus",
						"component": "pushgateway",
					},
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:            "prometheus-pushgateway",
							Image:           "prom/pushgateway:v0.4.0",
							ImagePullPolicy: imagePullPolicyIfNotPresent,
							Ports: []corev1.ContainerPort{
								{
									ContainerPort: 9091,
								},
							},
							ReadinessProbe: &corev1.Probe{
								Handler: corev1.Handler{
									HTTPGet: &corev1.HTTPGetAction{
										Path: "/#/status",
										Port: intstr.FromInt(9091),
									},
								},
								InitialDelaySeconds: 10,
								TimeoutSeconds: 10,
							},
						},
					},
				},
			},
		},
	}

	generator.Add("pushgateway-deployment.yaml", encoding.YAML(pushgatewayDeployment))
	
	pushgatewayService := corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name: "release-name-prometheus-alertmanager",
			Labels: map[string]string{
				"app": "prometheus",
				"component": "pushgateway",
			},
		},
		Spec: corev1.ServiceSpec{
			Ports: []corev1.ServicePort{
				{
					Name: "http",
					Port: 9091,
					Protocol: "TCP",
					TargetPort: intstr.FromInt(9091),
				},
			},
			Selector: map[string]string{
				"app": "prometheus",
				"component": "pushgateway",
			},
			Type: corev1.ServiceTypeClusterIP,
		},
	}

	generator.Add("pushgateway-service.yaml", encoding.YAML(pushgatewayService))
	
	// Server

	serverClusterRole := rbacv1beta1.ClusterRole{
		ObjectMeta: metav1.ObjectMeta{
			Name: "release-name-prometheus-kube-state-metrics",
			Labels: map[string]string{
				"app": "prometheus",
				"component": "server",
			},
		},
		Rules: []rbacv1beta1.PolicyRule{
			{
				APIGroups: []string{""},
				Resources: []string{
					"nodes",
					"nodes/proxy",
					"services",
					"endpoints",
					"pods",
					"ingresses",
				},
				Verbs: []string{
					"list",
					"watch",
					"get",
				},
			},
			{
				APIGroups: []string{""},
				Resources: []string{
					"configmaps",
				},
				Verbs: []string{
					"get",
				},
			},
			{
				APIGroups: []string{"extensions"},
				Resources: []string{
					"ingresses/status",
					"ingresses",
				},
				Verbs: []string{
					"get",
					"list",
					"watch",
				},
			},
			{
				NonResourceURLs: []string{"/metrics"},
				Verbs: []string{"get"},
			},
		},
	}

	generator.Add("server-clusterole.yaml", encoding.YAML(serverClusterRole))\

	serverClusterRoleBinding := rbacv1beta1.ClusterRoleBinding{
		ObjectMeta: metav1.ObjectMeta{
			Name: "release-name-prometheus-kube-state-metrics",
			Labels: map[string]string{
				"app": "prometheus",
				"component": "server",
			},
		},
		RoleRef: rbacv1beta1.RoleRef{
			APIGroup: "rbac.authorization.k8s.io",
			Kind: "ClusterRole",
			Name: "release-name-prometheus-server",
		},
		Subjects: []rbacv1beta1.Subject{
			{
				Kind: rbacv1beta1.ServiceAccountKind,
				Name: "release-name-prometheus-server",
				Namespace: namespace,
			},
		},
	}

	generator.Add("server-clusterrolebinding.yaml", encoding.YAML(serverClusterRoleBinding))

	serverDeployment := appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: "release-name-prometheus-server",
			Labels: map[string]string{
				"app": "prometheus",
				"component": "server",
			},
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: &int32One,
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": "prometheus",
						"component": "server",
					},
				},
				Spec: corev1.PodSpec{
					ServiceAccountName: "release-name-prometheus-server",
					InitContainers: []corev1.Container{
						{
							Name: "init-chown-data",
							Image: "busybox:latest",
							ImagePullPolicy: imagePullPolicyIfNotPresent,
							Command: []string{
								// 65534 is the nobody user that prometheus uses.
								"chown", "-R", "65534:65534", "/data",
							},
							VolumeMounts: []corev1.VolumeMount{
								{
									Name: "storage-volume",
									MountPath: "/data",
									SubPath: "",
								},
							},
						},
					},
					Containers: []corev1.Container{
						{
							Name: "prometheus-server-configmap-reload",
							Image: "jimmidyson/configmap-reload:v0.1",
							ImagePullPolicy: imagePullPolicyIfNotPresent,
							Args: []string{
								"--volume-dir=/etc/config",
								"--webhook-url=http://localhost:9090/-/reload",
							},
							VolumeMounts: []corev1.VolumeMount{
								{
									Name: "config-volume",
									MountPath: "/etc/config",
									ReadOnly: true,
								},
							},
						},
						{
							Name: "prometheus-server",
							Image: "prom/prometheus:v2.2.1",
							ImagePullPolicy: imagePullPolicyIfNotPresent,
							Args: []string{
								"--config.file=/etc/config/prometheus.yml",
								"--storage.tsdb.path=/data",
								"--web.console.libraries=/etc/prometheus/console_libraries",
								"--web.console.templates=/etc/prometheus/consoles",
								"--web.enable-lifecycle",
							},
							Ports: []corev1.ContainerPort{
								{
									ContainerPort: 9090,
								},
							},
							ReadinessProbe: &corev1.Probe{
								Handler: corev1.Handler{
									HTTPGet: &corev1.HTTPGetAction{
										Path: "/-/ready",
										Port: intstr.FromInt(9090),
									},
								},
								InitialDelaySeconds: 30,
								TimeoutSeconds: 30,
							},
							LivenessProbe: &corev1.Probe{
								Handler: corev1.Handler{
									HTTPGet: &corev1.HTTPGetAction{
										Path: "/-/healthy",
										Port: intstr.FromInt(9090),
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
					},
				},
			},
		},
	}

	serverPvc := corev1.PersistentVolumeClaim{
		ObjectMeta: metav1.ObjectMeta{
			Name: "release-name-prometheus-server",
			Labels: map[string]string{
				"app": "prometheus",
				"component": "server",
			},
		},
		Spec: corev1.PersistentVolumeClaimSpec{
			AccessModes: []corev1.PersistentVolumeAccessMode{
				corev1.ReadWriteOnce,
			},
			Resources: corev1.ResourceRequirements{
				Requests: corev1.ResourceList{
					corev1.ResourceStorage: resource.MustParse("8Gi"),
				},
			},
		},
	}

	generator.Add("server-pvc.yaml", encoding.YAML(serverPvc))

	serverService := corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name: "release-name-prometheus-server",
			Labels: map[string]string{
				"app": "prometheus",
				"component": "server",
			},
		},
		Spec: corev1.ServiceSpec{
			Ports: []corev1.ServicePort{
				{
					Name: "http",
					Port: 80,
					Protocol: "TCP",
					TargetPort: intstr.FromInt(9090),
				},
			},
			Selector: map[string]string{
				"app": "prometheus",
				"component": "server",
			},
			Type: corev1.ServiceTypeClusterIP,
		},
	}

	generator.Add("server-service.yaml", encoding.YAML(serverService))

	serverServiceAccount := corev1.ServiceAccount{
		ObjectMeta: metav1.ObjectMeta{
			Name: "release-name-prometheus-server",
			Labels: map[string]string{
				"app": "prometheus",
				"component": "server",
			},
		},
	}

	generator.Add("server-serviceaccount.yaml", encoding.YAML(serverServiceAccount))
}