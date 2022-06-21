# Mimic - Define your Configuration, Infrastructure and Deployments as Go Code

[![Go Report Card](https://goreportcard.com/badge/github.com/bwplotka/mimic)](https://goreportcard.com/report/github.com/bwplotka/mimic) 
[![GoDoc](https://godoc.org/github.com/bwplotka/mimic?status.svg)](https://pkg.go.dev/github.com/bwplotka/mimic)

`mimic`: A Go module for defining, templating and generating configuration in Go:

* Define your Configuration (e.g Envoy, Prometheus, Alertmanager, Nginx, Prometheus Alerts, Rules, Grafana Dashaboards etc.)
* Define your Infrastructure (e.g Terraform, Ansible, Puppet, Chef, Kubernetes etc)
* Define any other file that you can imagine.

...using simple, typed and testable Go code!

_`mimic`: [Mimic is a super-human with the ability to copy the powers and abilities of others](https://marvel.fandom.com/wiki/Power_Mimicry)._

## Getting Started

1. Create a new `.go` file for your config e.g `config/example.go`
2. Import mimic using Go 1.17+ e.g `go get github.com/bwplotka/mimic@latest`
3. Instantiate mimic and defer generation in your `main` function using the `mimic` module:

  ```go
    package main

    import ( 
		"github.com/bwplotka/mimic"
    )

    func main() {
        generator := mimic.New() 
		
		// Defer Generate to ensure we generate the output.
        defer generator.Generate()
          
		//...
  ```

4. Add typed configuration in your `main` to each file using encoding you want using `With` method: `generator.With("config").Add("some.yaml", encoding.GhodssYAML(set))`
5. Run `go run config/example.go generate`
6. You will now see the generated Kubernetes YAML file here: `cat gen/config/some.yaml`

See full [example](examples/kubernetes-statefulset/example.go) here:
    
 ```go
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
    	generator := mimic.New()
    
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
    	generator.With("config").Add(name+".yaml", encoding.GhodssYAML(set))
    }
```   
 
Now you are ready to start defining your own resources! 

Other examples can be found in [here](examples).

## What is `mimic`?

`mimic` is a package that allows you to define your configuration using Go and generate this into configuration files 
your application and tooling understands. 

## Why was `mimic` created?

`mimic` has been built from our past experience using this concept to configure our applications and infrastructure.

It offers not only to show the concept and an implementation example but also to share what we have learned, best practice and patterns that we believe are valuable for everyone. 

## But Why Go? 

Why you should define your templates/infra/configs in Go?


* Configuration as code ... like actual code, not json, yaml or tf.

* Go is a strongly **typed** language. This means that compiler and IDE of your choice will *massively* help you 
  find what config fields are allowed, what values enum expects and what is the type of each field.
  
* Unit/Integration test your configuration, infrastructure and deployment.  
    For example: 
    * Test your PromQL queries in Prometheus alerts works as expected? Just write unit test for those using e.g [this](https://github.com/prometheus/prometheus/blob/f678e27eb62ecf56e2b0bad82345925a4d6162a2/cmd/promtool/unittest.go#L37)
    * Enforce conventions such as service naming conventions via tests.

* Import configuration structs directly from the project you are configurating for example bring in Kubernestes, Prometheus or
any other structs that are exported. Allowing you to ensure your config matches the project. 

    No more blind searches and surprises. It cannot be safer or simpler than this.

* Versioning and dependency management. Utilize go modules to ensure you are using the correct version of the configuration
for the project version you are running.

* Documentation of your config, Go recommends [goDoc formatting](https://blog.golang.org/godoc-documenting-go-code), so 
you can leverage native comments for each struct's fields to document behaviour or details related to the config field. 
Giving you visibility in your config of exactly what your defining. See [this great Kubernetes struct](https://github.com/kubernetes/apimachinery/blob/master/pkg/apis/meta/v1/types.go#L55) as an example.
  
* Quick feedback loop. Catch most mistakes and incompatibilities in Go compile time, before you even deploy it further. 
  As you probably know one of Go goal is to have very fas compilation time, which feels like you are running a script.

* Keep the set of the languages used in your organization to a minimum - just one: Go, one of the cleanest, simplest and developer friendly languages around.

## What `mimic` is **NOT**

* It does NOT implement any deployment/distribution logic. 
* It is NOT intended to trigger any changes. Instead use the right tool for the job e.g. `kubectl apply`, `ansible`, `puppet`, `chef`, `terraform`
* It is NOT (yet) a registry for reusable templates, however we encourage the community to create public repositories for those!

## What does `mimic` include?

* [x] [providers](providers) go structs representing configuration for applications, infrastructure and container management.
  * Included are a set of go providers that do not have native structs OR are not easily importable (yet).
* [x] [encoding](encoding) a way to transform your go config struts into specific file types.
  * Included are json, yaml and jsonpb.
* [x] [abstractions](abstractions) a way to abstract concepts to a higher level if really needed (see best practice). 
* [x] Examples:
    * [Infra definitions for Prometheus remote read benchmarks on Kubernetes](examples/prometheus-remote-read-benchmark)
    * You want to add your own example here? Write to us on Slack or file GH issue!

## Want to help us and Contribute?

Please do! 

First start defining your configuration, infra and deployment as Go code!

Share with the community:
* Share the Go templates you create. 
* Share the Go configuration structs for non-Go projects. 
* Share the Go unit/integration/acceptance tests against certain providers's definitions.
* Share best practices and your experience!

Please use GitHub issues and our slack `#mimic` for feedback and questions. 

As always pull requests are VERY welcome as well!

## Have a project written in Go?
  
If you maintain your own project using Go it would be great to help the effort of making config as go a reality by
exposing your configuration structs for users to import.

How:
    * Maintain and export your config structs like Kubernetes does (it is an API and well documented types)
    * Define configuration file via [protobuf]() e.g like envoy [here](https://github.com/envoyproxy/envoy/tree/507191a36958bbeb1b11143fe0acb149f3f2fb00/api/envoy/config)

## Problems:

* What if project has only json schema? or no schema at all, just project written in different language: 

    Answer: Generate it yourself from YAML (e.g using [this online tool](https://mengzhuo.github.io/yaml-to-go/)). 
    Answer2: At some point if this concept will be big enough anyone can maintain useful Go module with typed, 
    documented and testable config for some providers like we have [in providers package](providers) 

* Importing native Go structs is the dream, however:

  * Not all project's config are prepared to be imported (config tied to implementation, huge deps, secret masked, no marshaling etc): 
  See: https://github.com/prometheus/alertmanager/pull/1804
  * This is where providers come in and we can define a set of structs to meet the config specified for your needs.

## Documentation

* [Best Practice](docs/best_practice.md)

## Other solutions

* [Cue](https://github.com/cuelang/cue)
* mixins
* jsonnet
* Pullumi (Paid)
