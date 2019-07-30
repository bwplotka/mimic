module github.com/bwplotka/mimic

go 1.12

replace github.com/prometheus/alertmanager => github.com/bwplotka/alertmanager v0.1.0-mimic-am-config

require (
	github.com/alecthomas/template v0.0.0-20190718012654-fb15b899a751 // indirect
	github.com/alecthomas/units v0.0.0-20190717042225-c3de453c63f4 // indirect
	github.com/ghodss/yaml v0.0.0-20150909031657-73d445a93680
	github.com/go-kit/kit v0.9.0
	github.com/go-openapi/swag v0.19.4
	github.com/gogo/protobuf v1.2.1
	github.com/golang/protobuf v1.3.2
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/pkg/errors v0.8.1
	github.com/pmezard/go-difflib v1.0.0
	github.com/prometheus/alertmanager v0.18.0
	github.com/prometheus/common v0.6.0
	github.com/rodaine/hclencoder v0.0.0-20190213202847-fb9757bb536e
	github.com/stretchr/testify v1.3.0
	gopkg.in/alecthomas/kingpin.v2 v2.2.6
	gopkg.in/yaml.v2 v2.2.2
	k8s.io/api v0.0.0-20190722141453-b90922c02518
	k8s.io/apimachinery v0.0.0-20190719140911-bfcf53abc9f8
)
