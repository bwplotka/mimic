// Copyright (c) bwplotka/mimic Authors
// Licensed under the Apache License 2.0.

package main

import (
	"github.com/bwplotka/mimic"
	"github.com/bwplotka/mimic/encoding"
	"github.com/go-openapi/swag"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func main() {
	generator := mimic.New().WithTopLevelComment(mimic.GeneratedComment)

	// Defer Generate to ensure we generate the output.
	defer generator.Generate()

	// Hook in your config below.
	// As an example Kubernetes configuration!
	const name = "some-statefulset"

	// Create some containers ... (imagine for now).
	var container1, container2, container3 corev1.Container
	var volume1 corev1.Volume

	// Configure a statefulset using native Kubernetes structs.
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
	// Now Add some-statefulset.yaml to the config folder.
	generator.With("config").WithTopLevelComment("Represents a K8s StatefulSet \nwith two containers.").Add(name+".yaml", encoding.GhodssYAML(set))
}
