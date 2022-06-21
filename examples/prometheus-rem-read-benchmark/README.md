# Prometheus definitions for remote read.

This definition contains Prometheus statefulsets definitions set up to test remote read changes described [here](https://docs.google.com/document/d/1JqrU3NjM9HoGLSTPYOvR217f5HBKBiJTqikEB9UiJL0/edit#).

## Usage

Generate Kubernetes YAML:

`go run examples/prometheus-remote-read-benchmark/main.go generate`

See the `gen` directory. `ls gen` should give you `prom-rr-test-streamed.yaml  prom-rr-test.yaml`

Those are 2 Prometheus + Thanos. One is baseline, second is a version with modified remote read that allows streaming encoded chunks.
You can use `kubectl apply` to deploy those. 

Those resources are crafted for benchmark purposes -> they generate artificial metric data.

[@bwplotka](https://bwplotka.dev/) is using those to benchmark new remote read with Thanos sidecar on live Kubernetes cluster.  