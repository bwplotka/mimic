GO := go

GOJSONSCHEMA := ${GOBIN}/gojsonschema
GOIMPORTS := ${GOBIN}/goimports

all: format test

format:
	@echo ">> formatting code"
	@$(GOIMPORTS) -w .

gen-dockercompose-config: $(GOJSONSCHEMA)
	@echo ">> generating"
	@$(GOJSONSCHEMA) -o providers/dockercompose/config_v3_7.go -p dockercompose providers/dockercompose/config_schema_v3.7.json

test:
	@echo ">> building binaries"
	@$(GO) test ./...

$(GOIMPORTS):
ifeq (${GOBIN},)
	@echo >&2 "GOBIN environment variable is not defined, where to put installed binaries?"; exit 1
endif
	@echo ">> installing goimports"
	@GO111MODULE=on GOOS= GOARCH= $(GO) install golang.org/x/tools/cmd/goimports

$(GOJSONSCHEMA):
ifeq (${GOBIN},)
	@echo >&2 "GOBIN environment variable is not defined, where to put installed binaries?"; exit 1
endif
	@echo ">> installing gojsonschema"
	@GO111MODULE=on GOOS= GOARCH= $(GO) install github.com/atombender/go-jsonschema/cmd/gojsonschema