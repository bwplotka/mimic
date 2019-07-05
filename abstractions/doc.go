/*
Abstractions are used to (as the name suggests) abstract away the underlying config and structs to the caller allowing
complex configurations or concepts to be created with minimal code.

We advise that abstractions are kept to a minimum as they hide configuration details from the user and instead use native
go structs from the project you are configuring or use providers that are a defined against the configuration specification.

Abstraction packages should follow a folder structure that allows clear understanding of the abstraction.
```
abstractions
  - prometheus   // Abstraction for generation of Prometheus configuration for Prometheus to consume.
  - kubernetes
    - prometheus // Kubernetes abstractions for the deployment of Prometheus.
```
*/
package abstractions
