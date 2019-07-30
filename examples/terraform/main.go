package main

import (
	"github.com/bwplotka/mimic"
	"github.com/bwplotka/mimic/encoding"
)

type Listener struct {
	InstancePort     int    `hcl:"instance_port"`
	InstanceProtocol string `hcl:"instance_protocol"`
	LBPort           int    `hcl:"lb_port"`
	LBProtocol       string `hcl:"lb_protocol"`
}

type AWSELB struct {
	Key       string   `hcl:",key"`
	Name      string   `hcl:"name"`
	Listener  Listener `hcl:"listener"`
	Instances []string `hcl:"instances"`
}

type AWSInstance struct {
	Key          string `hcl:",key"`
	Count        int    `hcl:"count"`
	AMI          string `hcl:"ami"`
	InstanceType string `hcl:"instance_type"`
}

func main() {
	generator := mimic.New()

	// Defer Generate to ensure we generate the output.
	defer generator.Generate()

	// Example taken from https://www.terraform.io/.
	instance := struct {
		AWSELB              AWSELB      `hcl:"resource \"aws_elb\""`
		AWSInstanceResource AWSInstance `hcl:"resource \"aws_instance\""`
	}{
		AWSELB: AWSELB{
			Key:  "frontend",
			Name: "frontend-load-balancer",
			Listener: Listener{
				InstancePort:     8080,
				InstanceProtocol: "http",
				LBPort:           80,
				LBProtocol:       "http",
			},
			Instances: []string{"${aws_instance.app.*.id}"},
		},
		AWSInstanceResource: AWSInstance{
			Key:          "app",
			Count:        5,
			AMI:          "ami-408c7f28",
			InstanceType: "t1.micro",
		},
	}

	// Now Add some-statefulset.yaml to the config folder.
	generator.With("terraform").Add("example.tf", encoding.HCL(instance))
}
