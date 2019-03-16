package dockercompose

import "fmt"
import "encoding/json"

// Steps how this file was constructed:
// * make gen-dockercompose-config
// * ConfigSchemaV37 -> Config
// * ConfigSchemaV37Volumes -> []Volume etc
type Config struct {
	// Configs corresponds to the JSON schema field "configs".
	Configs ConfigSchemaV37Configs `json:"configs,omitempty"`

	// Networks corresponds to the JSON schema field "networks".
	Networks []Network `json:"networks,omitempty"`

	// Secrets corresponds to the JSON schema field "secrets".
	Secrets []Secret `json:"secrets,omitempty"`

	// Services corresponds to the JSON schema field "services".
	Services []Service `json:"services,omitempty"`

	// Version corresponds to the JSON schema field "version".
	Version string `json:"version"`

	// Volumes corresponds to the JSON schema field "volumes".
	Volumes []Volume `json:"volumes,omitempty"`
}

type Constraints interface{}

type Deployment interface{}

type GenericResources []struct {
	// DiscreteResourceSpec corresponds to the JSON schema field
	// "discrete_resource_spec".
	DiscreteResourceSpec *GenericResourcesElemDiscreteResourceSpec `json:"discrete_resource_spec,omitempty"`
}

type GenericResourcesElemDiscreteResourceSpec struct {
	// Kind corresponds to the JSON schema field "kind".
	Kind *string `json:"kind,omitempty"`

	// Value corresponds to the JSON schema field "value".
	Value *float64 `json:"value,omitempty"`
}

type Healthcheck struct {
	// Disable corresponds to the JSON schema field "disable".
	Disable *bool `json:"disable,omitempty"`

	// Interval corresponds to the JSON schema field "interval".
	Interval *string `json:"interval,omitempty"`

	// Retries corresponds to the JSON schema field "retries".
	Retries *float64 `json:"retries,omitempty"`

	// StartPeriod corresponds to the JSON schema field "start_period".
	StartPeriod *string `json:"start_period,omitempty"`

	// Test corresponds to the JSON schema field "test".
	Test interface{} `json:"test,omitempty"`

	// Timeout corresponds to the JSON schema field "timeout".
	Timeout *string `json:"timeout,omitempty"`
}

type ListOfStrings []string

type ListOrDict interface{}

type Network interface{}

type Secret struct {
	// External corresponds to the JSON schema field "external".
	External interface{} `json:"external,omitempty"`

	// File corresponds to the JSON schema field "file".
	File *string `json:"file,omitempty"`

	// Labels corresponds to the JSON schema field "labels".
	Labels SecretLabels `json:"labels,omitempty"`

	// Name corresponds to the JSON schema field "name".
	Name *string `json:"name,omitempty"`
}

type SecretLabels interface{}

type Service struct {
	// Build corresponds to the JSON schema field "build".
	Build interface{} `json:"build,omitempty"`

	// CapAdd corresponds to the JSON schema field "cap_add".
	CapAdd []string `json:"cap_add,omitempty"`

	// CapDrop corresponds to the JSON schema field "cap_drop".
	CapDrop []string `json:"cap_drop,omitempty"`

	// CgroupParent corresponds to the JSON schema field "cgroup_parent".
	CgroupParent *string `json:"cgroup_parent,omitempty"`

	// Command corresponds to the JSON schema field "command".
	Command interface{} `json:"command,omitempty"`

	// Configs corresponds to the JSON schema field "configs".
	Configs []interface{} `json:"configs,omitempty"`

	// ContainerName corresponds to the JSON schema field "container_name".
	ContainerName *string `json:"container_name,omitempty"`

	// CredentialSpec corresponds to the JSON schema field "credential_spec".
	CredentialSpec *ServiceCredentialSpec `json:"credential_spec,omitempty"`

	// DependsOn corresponds to the JSON schema field "depends_on".
	DependsOn ListOfStrings `json:"depends_on,omitempty"`

	// Deploy corresponds to the JSON schema field "deploy".
	Deploy Deployment `json:"deploy,omitempty"`

	// Devices corresponds to the JSON schema field "devices".
	Devices []string `json:"devices,omitempty"`

	// Dns corresponds to the JSON schema field "dns".
	Dns ServiceDns `json:"dns,omitempty"`

	// DnsSearch corresponds to the JSON schema field "dns_search".
	DnsSearch ServiceDnsSearch `json:"dns_search,omitempty"`

	// Domainname corresponds to the JSON schema field "domainname".
	Domainname *string `json:"domainname,omitempty"`

	// Entrypoint corresponds to the JSON schema field "entrypoint".
	Entrypoint interface{} `json:"entrypoint,omitempty"`

	// EnvFile corresponds to the JSON schema field "env_file".
	EnvFile ServiceEnvFile `json:"env_file,omitempty"`

	// Environment corresponds to the JSON schema field "environment".
	Environment ServiceEnvironment `json:"environment,omitempty"`

	// Expose corresponds to the JSON schema field "expose".
	Expose []interface{} `json:"expose,omitempty"`

	// ExternalLinks corresponds to the JSON schema field "external_links".
	ExternalLinks []string `json:"external_links,omitempty"`

	// ExtraHosts corresponds to the JSON schema field "extra_hosts".
	ExtraHosts ServiceExtraHosts `json:"extra_hosts,omitempty"`

	// Healthcheck corresponds to the JSON schema field "healthcheck".
	Healthcheck *Healthcheck `json:"healthcheck,omitempty"`

	// Hostname corresponds to the JSON schema field "hostname".
	Hostname *string `json:"hostname,omitempty"`

	// Image corresponds to the JSON schema field "image".
	Image *string `json:"image,omitempty"`

	// Init corresponds to the JSON schema field "init".
	Init *bool `json:"init,omitempty"`

	// Ipc corresponds to the JSON schema field "ipc".
	Ipc *string `json:"ipc,omitempty"`

	// Isolation corresponds to the JSON schema field "isolation".
	Isolation *string `json:"isolation,omitempty"`

	// Labels corresponds to the JSON schema field "labels".
	Labels ServiceLabels `json:"labels,omitempty"`

	// Links corresponds to the JSON schema field "links".
	Links []string `json:"links,omitempty"`

	// Logging corresponds to the JSON schema field "logging".
	Logging *ServiceLogging `json:"logging,omitempty"`

	// MacAddress corresponds to the JSON schema field "mac_address".
	MacAddress *string `json:"mac_address,omitempty"`

	// NetworkMode corresponds to the JSON schema field "network_mode".
	NetworkMode *string `json:"network_mode,omitempty"`

	// Networks corresponds to the JSON schema field "networks".
	Networks interface{} `json:"networks,omitempty"`

	// Pid corresponds to the JSON schema field "pid".
	Pid interface{} `json:"pid,omitempty"`

	// Ports corresponds to the JSON schema field "ports".
	Ports []interface{} `json:"ports,omitempty"`

	// Privileged corresponds to the JSON schema field "privileged".
	Privileged *bool `json:"privileged,omitempty"`

	// ReadOnly corresponds to the JSON schema field "read_only".
	ReadOnly *bool `json:"read_only,omitempty"`

	// Restart corresponds to the JSON schema field "restart".
	Restart *string `json:"restart,omitempty"`

	// Secrets corresponds to the JSON schema field "secrets".
	Secrets []interface{} `json:"secrets,omitempty"`

	// SecurityOpt corresponds to the JSON schema field "security_opt".
	SecurityOpt []string `json:"security_opt,omitempty"`

	// ShmSize corresponds to the JSON schema field "shm_size".
	ShmSize interface{} `json:"shm_size,omitempty"`

	// StdinOpen corresponds to the JSON schema field "stdin_open".
	StdinOpen *bool `json:"stdin_open,omitempty"`

	// StopGracePeriod corresponds to the JSON schema field "stop_grace_period".
	StopGracePeriod *string `json:"stop_grace_period,omitempty"`

	// StopSignal corresponds to the JSON schema field "stop_signal".
	StopSignal *string `json:"stop_signal,omitempty"`

	// Sysctls corresponds to the JSON schema field "sysctls".
	Sysctls ServiceSysctls `json:"sysctls,omitempty"`

	// Tmpfs corresponds to the JSON schema field "tmpfs".
	Tmpfs ServiceTmpfs `json:"tmpfs,omitempty"`

	// Tty corresponds to the JSON schema field "tty".
	Tty *bool `json:"tty,omitempty"`

	// Ulimits corresponds to the JSON schema field "ulimits".
	Ulimits ServiceUlimits `json:"ulimits,omitempty"`

	// User corresponds to the JSON schema field "user".
	User *string `json:"user,omitempty"`

	// UsernsMode corresponds to the JSON schema field "userns_mode".
	UsernsMode *string `json:"userns_mode,omitempty"`

	// Volumes corresponds to the JSON schema field "volumes".
	Volumes []interface{} `json:"volumes,omitempty"`

	// WorkingDir corresponds to the JSON schema field "working_dir".
	WorkingDir *string `json:"working_dir,omitempty"`
}

type ServiceCredentialSpec struct {
	// File corresponds to the JSON schema field "file".
	File *string `json:"file,omitempty"`

	// Registry corresponds to the JSON schema field "registry".
	Registry *string `json:"registry,omitempty"`
}

type ServiceDns interface{}

type ServiceDnsSearch interface{}

type ServiceEnvFile interface{}

type ServiceEnvironment interface{}

type ServiceExtraHosts interface{}

type ServiceLabels interface{}

type ServiceLogging struct {
	// Driver corresponds to the JSON schema field "driver".
	Driver *string `json:"driver,omitempty"`

	// Options corresponds to the JSON schema field "options".
	Options ServiceLoggingOptions `json:"options,omitempty"`
}

type ServiceLoggingOptions map[string]interface{}

type ServiceSysctls interface{}

type ServiceTmpfs interface{}

type ServiceUlimits map[string]interface{}

type StringOrList interface{}

type Volume interface{}

// UnmarshalJSON implements json.Unmarshaler.
func (j *Config) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["version"]; !ok || v == nil {
		return fmt.Errorf("field version: required")
	}
	type Plain Config
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = Config(plain)
	return nil
}
