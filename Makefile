GOBIN ?= ${GOPATH}/bin
GO ?= $(shell which go)

FILES_TO_FMT      ?= $(shell find . -path ./vendor -prune -o -name '*.go' -print)

GOJSONSCHEMA         := $(GOBIN)/gojsonschema

GOIMPORTS_VERSION    ?= 9d4d845e86f14303813298ede731a971dd65b593
GOIMPORTS            ?= $(GOBIN)/goimports-$(GOIMPORTS_VERSION)
GOLANGCILINT_VERSION ?= v1.17.1
GOLANGCILINT         ?= $(GOBIN)/golangci-lint-$(GOLANGCILINT_VERSION)
LICHE_VERSION        ?= 2a2e6e56f6c615c17b2e116669c4cdb31b5453f3
LICHE                ?= $(GOBIN)/liche-$(LICHE_VERSION)

GO111MODULE       ?= on
export GO111MODULE

.PHONY: all
all: format test

.PHONY: update-go-deps
update-go-deps:
	@echo ">> updating Go dependencies"
	@for m in $$($(GO) list -mod=readonly -m -f '{{ if and (not .Indirect) (not .Main)}}{{.Path}}{{end}}' all); do \
		$(GO) get $$m; \
	done
	@$(GO) mod tidy

# check-docs checks if documentation have discrepancy with flags and if the links are valid.
.PHONY: check-docs
check-docs: $(LICHE)
	@$(LICHE) --document-root . *.md

.PHONY: check-go-mod
check-go-mod:
	@go mod verify

.PHONY: format
format: $(GOIMPORTS)
	@echo ">> formatting code"
	@$(GOIMPORTS) -w $(FILES_TO_FMT)

gen-dockercompose-config: $(GOJSONSCHEMA)
	@echo ">> generating"
	@$(GOJSONSCHEMA) -o providers/dockercompose/config_v3_7.go -p dockercompose providers/dockercompose/config_schema_v3.7.json

.PHONY: lint
lint: $(GOLANGCILINT)
	@echo ">> linting all of the Go files"
	@$(GOLANGCILINT) run --disable-all -E goimports ./...
	@$(GOLANGCILINT) run ./...

.PHONY: test
test:
	@echo ">> testing binaries"
	@go test ./...
	@cd examples/kubernetes-statefulset && go run example.go generate
	@cd examples/prometheus-remote-read-benchmark && go run main.go generate
	@cd examples/terraform && go run main.go generate

# $(1): Go install path. (e.g github.com/campoy/embedmd)
# $(2): Tag.
define fetch_go_bin_version
	@echo ">> installing $(1) at $(2)"
	@GO111MODULE=on GOOS= GOARCH= go get $(1)@$(2)
	@mv -- '$(GOBIN)/$(shell basename $(1))' '$(GOBIN)/$(shell basename $(1))-$(2)'
	@go mod tidy
endef

$(GOIMPORTS):
	$(call fetch_go_bin_version,golang.org/x/tools/cmd/goimports,$(GOIMPORTS_VERSION))

$(GOJSONSCHEMA):
	@echo ">> installing gojsonschema"
	@GO111MODULE=on GOOS= GOARCH= $(GO) install github.com/atombender/go-jsonschema/cmd/gojsonschema

$(GOLANGCILINT):
	$(call fetch_go_bin_version,github.com/golangci/golangci-lint/cmd/golangci-lint,$(GOLANGCILINT_VERSION))

$(LICHE):
	$(call fetch_go_bin_version,github.com/raviqqe/liche,$(LICHE_VERSION))