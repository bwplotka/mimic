package main

import (
	"github.com/bwplotka/gocodeit"
	"github.com/bwplotka/gocodeit/encoding"
	"github.com/go-openapi/swag"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func main() {
	gci := gocodeit.New()

	// Make sure to Generate at the very end.
	defer gci.Generate()

	// Hook your definitions below.
	// For example Kubernetes configuration!
	const name = "some-statefulset"

	// Let's imagine we fill those...
	var container1, container2, container3 corev1.Container
	var volume1 corev1.Volume

	// Example statefulset using native Kubernetes structs.
	set := appsv1.StatefulSet{
		TypeMeta: metav1.TypeMeta{
			Kind:       "StatefulSet",
			APIVersion: "apps/v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
			Labels: map[string]string{
				"app": name,
			},
		},
		Spec: appsv1.StatefulSetSpec{
			Replicas:    swag.Int32(2),
			ServiceName: name,
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": name,
					},
				},
				Spec: corev1.PodSpec{
					InitContainers: []corev1.Container{container1},
					Containers:     []corev1.Container{container2, container3},
					Volumes:        []corev1.Volume{volume1},
				},
			},
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": name,
				},
			},
		},
	}
	// Generate file in config directory using chosen encoding.
	gci.With("config").Add(name+".yaml", encoding.GhodssYAML(set))
}
