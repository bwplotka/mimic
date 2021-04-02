module github.com/bwplotka/mimic/examples/kubernetes-statefulset

go 1.16

require (
	github.com/bwplotka/mimic v0.0.0-20190730202618-06ab9976e8ef
	github.com/go-openapi/swag v0.19.15
	k8s.io/api v0.20.5
	k8s.io/apimachinery v0.20.5
)

// This module is meant to be executed from repo root.
replace github.com/bwplotka/mimic => ../../
