// Copyright (c) bwplotka/mimic Authors
// Licensed under the Apache License 2.0.

package volumes

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"

	"github.com/bwplotka/mimic"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type VolumeAndMount struct {
	corev1.VolumeMount
	// corev1.Volume has just Name and VolumeSource. A name field is already present in the VolumeMount, so we just add
	// the VolumeSource field here directly.
	VolumeSource corev1.VolumeSource
}

type VolumesAndMounts []VolumeAndMount

func (vam VolumeAndMount) Volume() corev1.Volume {
	return corev1.Volume{
		Name:         vam.Name,
		VolumeSource: vam.VolumeSource,
	}
}

func (vams VolumesAndMounts) Volumes() []corev1.Volume {
	volumes := make([]corev1.Volume, 0, len(vams))
	for _, vam := range vams {
		volumes = append(volumes, vam.Volume())
	}
	return volumes
}

func (vams VolumesAndMounts) VolumeMounts() []corev1.VolumeMount {
	mounts := make([]corev1.VolumeMount, 0, len(vams))
	for _, vam := range vams {
		mounts = append(mounts, vam.VolumeMount)
	}
	return mounts
}

type ConfigAndMount struct {
	metav1.ObjectMeta
	corev1.VolumeMount //nolint:govet
	Data               map[string]string
}

func (m ConfigAndMount) ConfigMap() corev1.ConfigMap {
	return corev1.ConfigMap{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "v1",
			Kind:       "ConfigMap",
		},
		ObjectMeta: m.ObjectMeta,
		Data:       m.Data,
	}
}

func (m ConfigAndMount) VolumeAndMount() VolumeAndMount {
	return VolumeAndMount{
		VolumeSource: corev1.VolumeSource{
			ConfigMap: &corev1.ConfigMapVolumeSource{
				LocalObjectReference: corev1.LocalObjectReference{Name: m.ObjectMeta.Name},
			},
		},
		VolumeMount: m.VolumeMount,
	}
}

func (m ConfigAndMount) HashEnv(name string) corev1.EnvVar {
	h := sha256.New()
	if err := json.NewEncoder(h).Encode(m.Data); err != nil {
		mimic.Panicf("failed to JSON encode & hash configMap data for %s, err: %v", m.VolumeMount.Name, err)
	}

	return corev1.EnvVar{
		Name:  name,
		Value: base64.URLEncoding.EncodeToString(h.Sum(nil)),
	}
}
