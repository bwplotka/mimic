# Prometheus on Kubernetes

This example provides a reference implementation of how to create all configuration files needed to spin up a fully-featured Prometheus on Kubernetes.

Writing only Go code, we define and generate Kubernetes YAMLs and Prometheus configuration. 

We create the exact same output configuration files as created by the stock [helm prometheus chart](https://github.com/helm/charts/tree/master/stable/prometheus).

## How to run

```bash
go run main.go generate
``` 

Then view the output YAML files in `/gen`