# Guide & Best Practices

**TL:DR; defining configuration in Go code *is different* than writing robust production go code for an application.**

This code serves a certain goals like command-line only based file generation, configuring and defining files that will be 
consumed by other systems. This means that you need to switch context from other Golang code you might write for different purposes.

In details this means:

* **DO NOT** try to circumvent the need to understand the upstream product. As with any configuration of a
   product, you need to understand how to use that product just as if you were
  configuring it directly.

* Use native structs as much as possible as this e.g Kubernetes struts directly: 
  * Immediately maps with what you will see once file is deployed. (e.g if you do `kubectl get po <your-pod> -o yaml` someday)
    It helps to debug in future and allows others to quickly tweak it even if you are from different team.
  * Helps to reuse other templates/patterns from upstream in future

* Use `panic` as your error handling for unexpected input. Normally it's never acceptable but in this case, it gives you quick feedback on what
  is wrong with your configuration/generation and immediately HALTS the generation which is desired.

* Don't use concurrency. Avoid unnecessary magic.

* Minimalise unnecessary abstractions. Your configs, infra, deployments should be as verbose as possible. AVOID using `if` 
  conditions, prefer consistent environments etc. It's *important* as with a powerful language it's easy and tempting to 
  add abstraction layers and special cases. This is your DEFINITIONS. It should be verbose. 
  
  **300 hundreds lines functions ARE acceptable here.**

* Where the overhead is not prohibitive, avoid primitive obsession, config structure and values should be written with an
  appropriate type. In particular, large string blobs and using strings to specify non-string (e.g. integer, boolean) 
  options should be avoided. This will also give you strong typing for your configuration.

* Associate and reference keys together. For example if you some Kubernetes deployment expects ConfigMap called "my-conf",
  consider putting "my-conf" in constant and reference it in deployment, or even better, reference ConfigMap.Name directly!

* Factor out and re-use a constant or parameter if and only if it is required by the system being configured
    
    * Do not refactor out a common value just because the value occurs in multiple places.
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

    i.e. you may update fields on a struct only in the function containing the literal that created the
    struct.

* Avoid basing configuration/behaviour on the environment or cluster

    i.e. avoid code like “if running in staging do X, if running in production to Y”. Refactor such code
    to take a parameter for this behaviour and have the definition of staging and production set the
    parameter appropriately.

* Avoid basing any output or configuration of your applications or infrastructure on your organisational structure ... this will 
  and does change often so don't bake it into your configuration. 

* **DO NOT** put secrets (tokens, private keys, etc.) in your config definition or output (or anywhere else in your code!)

    i.e For kubernetes you could put secrets in via Kuberenetes Secrets. Alternatively define env variables and 
    use env vars substitution.
    
    Ultimately use the best tools for this! See [vault](https://www.vaultproject.io/)

* Do not fetch inputs from outside.

    i.e. do not run external commands or make web requests.

* Apply YAGNI/KISS/etc.

    (“You ain’t gonna need it”, “Keep it simple, stupid”)

    Aim for simplicity, e.g. do not create a new helper if you are going to use it only in one location.
    You can always refactor it out later.
