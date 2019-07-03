# "Define your Deployments, Infrastructure and Configuration as a Golang Code"

[![Go Report Card](https://goreportcard.com/badge/github.com/bwplotka/gocodeit)](https://goreportcard.com/report/github.com/bwplotka/gocodeit) 
[![GoDoc](https://godoc.org/github.com/bwplotka/gocodeit?status.svg)](https://godoc.org/github.com/bwplotka/gocodeit)

`GoCodeIt` (`gci`): Golang module showcasing an awesome **concept** for:

* Defining Configuration (e.g Envoy, Prometheus, Alertmanager, Nginx etc)
* Defining Infrastructure (e.g Terraform, Ansible, Puppet, Chef, Prometheus Alerts, Rules, Grafana Dashaboards etc)
* Defining Deployments (e.g Docker compose, Kubernetes, etc)
* Defining any other file with whatever file format

...as a simple, templated, typed and testable Golang code!

## How to use it?

1. Start new project, create git repository.
1. Create a new golang file e.g `projects/example.go`
1. Import gocodeit using Golang 1.12+ via `go get github.com/bwplotka/gocodeit`.
1. Add initial hooks in your `main` package using `gocodeit` library. For [example](projects/example.go):

    ```go
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
    ```   
    
1. Run `go run projects/example.go generate`
1. You should see generated Kubernetes YAML file here: `cat gcigen/config/some-statefulset.yaml`
 
Now you are ready to start defining your own resources! 

`gocodeit` helps your template and maintain your definition but those not solve deploying. The basic flow would be to
 * clone your git project in some CI/dev machine
 * `go run <your project> generate`
 * Distributing config files where you want, e,g just `kubectl apply` resources if it's Kubernetes etc.

See other example or actually useful personal projects [here](projects)

## Why? 

Because we learnt that this approach is quite beneficial, the hard way. :rage4:

Why you should define your templates/infra/configs in Golang?

* Golang is a strongly **typed** language. This means that compiler and IDE of your choice will *massively* help you 
  find what config fields certain config allows, what values enum expects and what is the type of each field.
  
* Golang recommends [goDoc formatting](https://blog.golang.org/godoc-documenting-go-code), which means that you can leverage 
  native comments for each struct's fields to document behaviour or details related to the config field. Just go to config struct
  source code via IDE and check what each fields actually means! See [this great Kubernetes struct](https://github.com/kubernetes/apimachinery/blob/master/pkg/apis/meta/v1/types.go#L55) as an example.
   
* Reduce incompatibilities and unknowns to minimum. Golang allows versioned dependency management. This means that you if project you
  define configs or definitions against, is in Golang or uses configuration defined by protobuf, you can natively import such typed config 
  **directly** from source. To generate your configuration/definitions, you can fill exactly the same struct, as the project you configure will use 
  for unmarshal. No more blind searches and surprises. It cannot be safer or simpler than this.

  For example: you can get feedback early on, that the type you used in your Kubernetes stateful set definition’s field is wrong! 
  Or the field you used in the Prometheus configuration file does not exist anymore in version 2.0 of Prometheus. 
  
* Quick feedback loop. Catch most mistakes and incompatibilities in Golang compile time, before you even deploy it further. 
  As you probably know one of Golang goal is to have very fas compilation time, which feels like you are running a script.

* Unit/Integration test your configuration, infrastructure and deployment. 
    
    For example: Want to check if you PromQL queries in Prometheus alerts works as expected? Just write unit test for those using e.g [this](https://github.com/prometheus/prometheus/blob/f678e27eb62ecf56e2b0bad82345925a4d6162a2/cmd/promtool/unittest.go#L37)
    Want to check if your alertmanager routing works? Create unit test using native routing logic imported directly from github.com/prometheus/alertmananger

* Keep the set of the languages used in your organization to a minimum - just one: Golang, which one of the cleanest and easiest to read languages made.

* Associate things. If you create a Kubernetes Deployment that expects configMap A, it's sometimes easy to make a typo or forget to apply that configMap A.
  With Golang you can associate those two together either by common constant string, or by literally referencing `ConfigMap.Name` in your Kubernetes Deployment. 
  Catch the bugs early!

## What this project is: GoCodeIt

This projects is to show idea. An awesome pattern that we believe, everyone, especially backend engineers should use. 
It's not about using exactly this library that we aim for. The implementation might be not perfect and it was created from scratch based on the lessons we learnt. 

*Don't use this implementation* if you don't want to. Instead, **be inspired to create your own Golang helpers to generate Configuration, Infrastructure and Deployment definitions from Golang code!**

* Share the Go templates you create. 
* Share the Go onfiguration structs for non-Golang projects. 
* Share the Go unit/integration/acceptance tests against certain providers's definitions.
* Share best practices and your experience!

So what `GoCodeIt` Go module includes?

* [x] Minimal [providers](providers) package for config types that are not natively hosted as Golang code OR are not easily importable, yet.
* [x] Projects:
    * [Infra definitions for Prometheus remote read benchmarks on Kubernetes](projects/prom-remote-read-bench)
    * [(in progress) monioring for website using Dockercompose, Prometheus and Thanos](projects/infra-my-mon)
    * You want to add your own example here? Write to us on Slack or file GH issue!

`GoCodeIt` is used already on "production" for our personal projects infrastructure like [here](projects/infra-my-mon).

## What this project is NOT

* It is NOT implementing any deployment/distribution logic. It is NOT intended to trigger any changes. Use `kubectl apply`, ansible, puppet, chef, terraform, 
tool that fits the job for this. `GoCodeIt` only helps you to define, test and generate files to local filesystem, with all information needed by automation.

* It is NOT (yet) a registry for reusable templates, however we encourage to create public repositories for those!

## Important: Guide & best practices

The TL;DR is that defining your information in Golang code IS different then writing robust code for an application logic.
This code serves a certain goals like command-line only based file generation, configuring and defining files that will be 
consumed by other systems. This means that you need to switch context from other Golang code you might write for different purposes.

In details this means:

* Use panics as your error handling for unexpected input. Normally it's never acceptable but in this case, it gives you quick feedback on what
  is wrong with your configuration/generation and immdiately HALTS the generation which is desired.

* Don't use concurrency. Avoid unnecessary magic.

* Minimalise unnecessary abstractions. Your configs, infra, deployments should be as verbose as possible. AVOID using `if` 
  conditions, prefer consistent environments etc. It's *important* point. With powerful language it's tempting to add abstraction
  layers and thousands of special cases. This is your DEFINITIONS. It should be verbose. 
  
  **300 hundreds lines functions ARE acceptable here.**

* Particularly do NOT try to circumvent the need to understand the upstream product. If you use Golang to
  configure a product, you need to understand how to use that product just as if you were
  configuring it directly.
    
* Use native structs as much as possible as this e.g Kubernetes struts directly: 
  * Immediately maps with what you will see once file is deployed. (e.g if you do `kubectl get po <your-pod> -o yaml` someday)
    It helps to debug in future and allows others to quickly tweak it even if you are from different team.
  * Helps to reuse other templates/patterns from upstream in future

* Where the overhead is not prohibitive, config structure and values should be written with an
  appropriate Golang type. In particular, large string blobs and using strings to specify
  non-string (e.g. integer, boolean) options should be avoided.

* Associate and reference keys together. For example if you some Kubernetes deployment expects ConfigMap called "my-conf",
  consider putting "my-conf" in constant and reference it in deployment, or even better, reference ConfigMap.Name directly!

* Factor out and re-use a constant or parameter if and only if it is required by the system being configured
    
    * Do not factor out a common value just because the value occurs in multiple places.
    * If a system requires the same value at multiple locations to work correctly, factor that out into
      a constant or parameter.
    * If you factor out a value into a constant or parameter, that constant or parameter should have
      clearly defined semantics for the system being configured, usually reflected by the name given to
      the constant or parameter.
    * If a value (or substring) occurs in multiple locations with no connection between them:
      
      * If the values do not not need to be parametrised, use separate literals.
      * If the values need to be parameterised, use separate parameters.
    
    This is important to keep in mind in configuration-oriented code. Trying to refactor anything
    similar-looking will result in a spaghetti of dependencies where it is hard for maintainers to
    figure out which dependencies are required by the system and which happen to be an artifact of
    implementation. The more declarative nature of this code also means that many literals are described
    by the field they are being assigned to, and do not need further naming to clarify their purpose.
    
* Treat parameters and returned structs as immutable.

    I.e. you may update fields on a struct only in the function containing the literal that created the
    struct.

* Avoid basing configuration/behaviour on the environment or cluster

    I.e. avoid code like “if running in staging do X, if running in production to Y”. Refactor such code
    to take a parameter for this behaviour and have the definition of staging and production set the
    parameter appropriately.


* Try to NOT put secrets (tokens, private keys, etc.) in definition output (or anywhere else in your golang code!)

    I.e For kubernetes, put secrets via different mean as Kuberenetes Secrets. Alternatively define env variables and 
    use env vars substitution. For terraform, use the vault provider to have terraform pull secrets directly from vault.
    
    Ultimately use the best tools for this! See [vault](https://www.vaultproject.io/)
    
    Reason is that you want your secrets to be safely stored and rotated.

* Do not fetch inputs from outside.

    I.e. do not run external commands or make web requests.

* Apply YAGNI/KISS/etc.

    (“You ain’t gonna need it”, “Keep it simple, stupid”)

    Aim for simplicity, eg. do not create a new helper if you are going to use it only in one location.
    You can always factor it out later.

## Do you want to help us? Contribute?

Please do! 

First of all start defining your configuration, infra and deployment as Golang code! (: Share this concept, let's start to:
* Get rid of unreadable, untyped teamplating languages like jsonnet, m4 etc
* Test your configuration; FINALLY!

Please use GitHub issues and our slack `#gocodeit` for feedback and questions. Pull requests are VERY welcome as well!

## Problems:

* What if project has only json schema? or no schema at all, just project written in different language: 

    Answer: Generate it yourself from YAML (e.g using [this online tool](https://mengzhuo.github.io/yaml-to-go/)). 
    Answer2: At some point if this concept will be big enough anyone can maintain useful Go module with typed, 
    documented and testable config for some providers like we have [in providers package](providers) 

* Importing native Go structs is the dream however:

  * Not all project's config are prepared to be imported (config tied to implementation, huge deps, secret masked, no marshaling etc): 
  See: https://github.com/prometheus/alertmanager/pull/1804
  * While YAML schema should not change a lot (for stable APIs, 1+ semantic versions for projects), 
  internal underlying Golang struct's code can change a lot. If this concept will be widely popular it will encourage project to:
    * either maintain your config structs as Kubernetes does (it is an API and well documented types)
    * define configuration file via [protobuf]() e.g like envoy [here](https://github.com/envoyproxy/envoy/tree/507191a36958bbeb1b11143fe0acb149f3f2fb00/api/envoy/config)
    
## Other solutions

* https://github.com/cuelang/cue
