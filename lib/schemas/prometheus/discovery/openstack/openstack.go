// Copyright (c) bwplotka/mimic Authors
// Licensed under the Apache License 2.0.

// Copyright 2017 The Prometheus Authors
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package openstack

import (
	config_util "github.com/prometheus/common/config"
	"github.com/prometheus/common/model"
)

// SDConfig is the configuration for OpenStack based service discovery.
type SDConfig struct {
	IdentityEndpoint string                `yaml:"identity_endpoint"`
	Username         string                `yaml:"username"`
	UserID           string                `yaml:"userid"`
	Password         config_util.Secret    `yaml:"password"`
	ProjectName      string                `yaml:"project_name"`
	ProjectID        string                `yaml:"project_id"`
	DomainName       string                `yaml:"domain_name"`
	DomainID         string                `yaml:"domain_id"`
	Role             Role                  `yaml:"role"`
	Region           string                `yaml:"region"`
	RefreshInterval  model.Duration        `yaml:"refresh_interval,omitempty"`
	Port             int                   `yaml:"port"`
	AllTenants       bool                  `yaml:"all_tenants,omitempty"`
	TLSConfig        config_util.TLSConfig `yaml:"tls_config,omitempty"`
}

// OpenStackRole is role of the target in OpenStack.
type Role string

// The valid options for OpenStackRole.
const (
	// OpenStack document reference
	// https://docs.openstack.org/nova/pike/admin/arch.html#hypervisors
	OpenStackRoleHypervisor Role = "hypervisor"
	// OpenStack document reference
	// https://docs.openstack.org/horizon/pike/user/launch-instances.html
	OpenStackRoleInstance Role = "instance"
)
