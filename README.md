# gocodeit

GoCodeIt (gci): Library showcasing a concept for defining:

* Configuration (e.g envoy, Prometheus, etc)
* Infrastructure (e.g terraform, ansible. etc)
* Deployments (e.g docker compose, Kubernetes, etc)

...as a simple, templated, typed and testable Golang code!

## Why

## How to use it?

## Problems:

* What if project has only json schema? or no schema at all, just project written in different language.
* Importing native Go structs is doable, but:
  * not all are prepared to be imported (comments, huge deps, secret masked etc): See: https://github.com/prometheus/alertmanager/pull/1804
  * While fields of YAML should not change (API), code can change a lot, so maybe it's better to just vendor your own version?
