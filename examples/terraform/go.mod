module github.com/bwplotka/mimic/examples/terraform

go 1.16

require (
	github.com/bwplotka/mimic v0.0.0-20190730202618-06ab9976e8ef
	github.com/kr/text v0.2.0 // indirect
	github.com/niemeyer/pretty v0.0.0-20200227124842-a10e7caefd8e // indirect
	gopkg.in/check.v1 v1.0.0-20200227125254-8fa46927fb4f // indirect
	gopkg.in/yaml.v3 v3.0.0-20200615113413-eeeca48fe776 // indirect
)

// This module is meant to be executed from repo root.
replace github.com/bwplotka/mimic => ../../
