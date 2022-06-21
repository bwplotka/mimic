module github.com/bwplotka/mimic/examples/prometheus-rem-read-benchmark

go 1.16

require (
	github.com/bwplotka/mimic v0.0.0-20190730202618-06ab9976e8ef
	github.com/bwplotka/mimic/lib/schemas/prometheus v0.0.0-20210423144650-c81ccdd3db1c
	github.com/go-openapi/swag v0.19.15
	github.com/prometheus/common v0.20.0
	k8s.io/api v0.20.5
	k8s.io/apimachinery v0.20.5
)

// This module is meant to be executed from repo root.
replace github.com/bwplotka/mimic => ../../
