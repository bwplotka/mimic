package dockercompose

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// Generated using:
// * cloning and building github.com/a-h/generate
// * ./schema-generate -p dockercompose -o providers/dockercompose/config_v3_7.go providers/dockercompose/config_schema_v3.7.json
// * all *_object refactored to no _object suffix.
// * Associate arrays together (e.g Secrets is now []*Secret)
// Made version hardcoded. We expect 3 and we generate always 3.

type Config struct {
	Configs  []*ConfigFile `json:"configs,omitempty"`
	Networks []*Network    `json:"networks,omitempty"`
	Secrets  []*Secret     `json:"secrets,omitempty"`
	Services []*Service    `json:"services,omitempty"`
	Volumes  []*Volume     `json:"volumes,omitempty"`
}

type ConfigFile struct {
	External interface{} `json:"external,omitempty"`
	File     string      `json:"file,omitempty"`
	Labels   interface{} `json:"labels,omitempty"`
	Name     string      `json:"name,omitempty"`
}

// ConfigItems
type ConfigItems struct {
	Subnet string `json:"subnet,omitempty"`
}

// CredentialSpec
type CredentialSpec struct {
	File     string `json:"file,omitempty"`
	Registry string `json:"registry,omitempty"`
}

// Deployment
type Deployment struct {
	EndpointMode   string          `json:"endpoint_mode,omitempty"`
	Labels         interface{}     `json:"labels,omitempty"`
	Mode           string          `json:"mode,omitempty"`
	Placement      *Placement      `json:"placement,omitempty"`
	Replicas       int             `json:"replicas,omitempty"`
	Resources      *Resources      `json:"resources,omitempty"`
	RestartPolicy  *RestartPolicy  `json:"restart_policy,omitempty"`
	RollbackConfig *RollbackConfig `json:"rollback_config,omitempty"`
	UpdateConfig   *UpdateConfig   `json:"update_config,omitempty"`
}

// DiscreteResourceSpec
type DiscreteResourceSpec struct {
	Kind  string  `json:"kind,omitempty"`
	Value float64 `json:"value,omitempty"`
}

// DriverOpts
type DriverOpts struct {
}

// External
type External struct {
	Name string `json:"name,omitempty"`
}

// GenericResourcesItems
type GenericResourcesItems struct {
	DiscreteResourceSpec *DiscreteResourceSpec `json:"discrete_resource_spec,omitempty"`
}

// Healthcheck
type Healthcheck struct {
	Disable     bool        `json:"disable,omitempty"`
	Interval    string      `json:"interval,omitempty"`
	Retries     float64     `json:"retries,omitempty"`
	StartPeriod string      `json:"start_period,omitempty"`
	Test        interface{} `json:"test,omitempty"`
	Timeout     string      `json:"timeout,omitempty"`
}

// Ipam
type Ipam struct {
	Config []*ConfigItems `json:"config,omitempty"`
	Driver string         `json:"driver,omitempty"`
}

// Limits
type Limits struct {
	Cpus   string `json:"cpus,omitempty"`
	Memory string `json:"memory,omitempty"`
}

// Logging
type Logging struct {
	Driver  string   `json:"driver,omitempty"`
	Options *Options `json:"options,omitempty"`
}

// Network
type Network struct {
	Attachable bool        `json:"attachable,omitempty"`
	Driver     string      `json:"driver,omitempty"`
	DriverOpts *DriverOpts `json:"driver_opts,omitempty"`
	External   interface{} `json:"external,omitempty"`
	Internal   bool        `json:"internal,omitempty"`
	Ipam       *Ipam       `json:"ipam,omitempty"`
	Labels     interface{} `json:"labels,omitempty"`
	Name       string      `json:"name,omitempty"`
}

// Options
type Options struct {
}

// Placement
type Placement struct {
	Constraints []string            `json:"constraints,omitempty"`
	Preferences []*PreferencesItems `json:"preferences,omitempty"`
}

// PreferencesItems
type PreferencesItems struct {
	Spread string `json:"spread,omitempty"`
}

// Reservations
type Reservations struct {
	Cpus             string                   `json:"cpus,omitempty"`
	GenericResources []*GenericResourcesItems `json:"generic_resources,omitempty"`
	Memory           string                   `json:"memory,omitempty"`
}

// Resources
type Resources struct {
	Limits       *Limits       `json:"limits,omitempty"`
	Reservations *Reservations `json:"reservations,omitempty"`
}

// RestartPolicy
type RestartPolicy struct {
	Condition   string `json:"condition,omitempty"`
	Delay       string `json:"delay,omitempty"`
	MaxAttempts int    `json:"max_attempts,omitempty"`
	Window      string `json:"window,omitempty"`
}

// RollbackConfig
type RollbackConfig struct {
	Delay           string  `json:"delay,omitempty"`
	FailureAction   string  `json:"failure_action,omitempty"`
	MaxFailureRatio float64 `json:"max_failure_ratio,omitempty"`
	Monitor         string  `json:"monitor,omitempty"`
	Order           string  `json:"order,omitempty"`
	Parallelism     int     `json:"parallelism,omitempty"`
}

// Secret
type Secret struct {
	External interface{} `json:"external,omitempty"`
	File     string      `json:"file,omitempty"`
	Labels   interface{} `json:"labels,omitempty"`
	Name     string      `json:"name,omitempty"`
}

// Service
type Service struct {
	Build           interface{}     `json:"build,omitempty"`
	CapAdd          []string        `json:"cap_add,omitempty"`
	CapDrop         []string        `json:"cap_drop,omitempty"`
	CgroupParent    string          `json:"cgroup_parent,omitempty"`
	Command         interface{}     `json:"command,omitempty"`
	Configs         []interface{}   `json:"configs,omitempty"`
	ContainerName   string          `json:"container_name,omitempty"`
	CredentialSpec  *CredentialSpec `json:"credential_spec,omitempty"`
	DependsOn       []string        `json:"depends_on,omitempty"`
	Deploy          interface{}     `json:"deploy,omitempty"`
	Devices         []string        `json:"devices,omitempty"`
	Dns             interface{}     `json:"dns,omitempty"`
	DnsSearch       interface{}     `json:"dns_search,omitempty"`
	Domainname      string          `json:"domainname,omitempty"`
	Entrypoint      interface{}     `json:"entrypoint,omitempty"`
	EnvFile         interface{}     `json:"env_file,omitempty"`
	Environment     interface{}     `json:"environment,omitempty"`
	Expose          []interface{}   `json:"expose,omitempty"`
	ExternalLinks   []string        `json:"external_links,omitempty"`
	ExtraHosts      interface{}     `json:"extra_hosts,omitempty"`
	Healthcheck     *Healthcheck    `json:"healthcheck,omitempty"`
	Hostname        string          `json:"hostname,omitempty"`
	Image           string          `json:"image,omitempty"`
	Init            bool            `json:"init,omitempty"`
	Ipc             string          `json:"ipc,omitempty"`
	Isolation       string          `json:"isolation,omitempty"`
	Labels          interface{}     `json:"labels,omitempty"`
	Links           []string        `json:"links,omitempty"`
	Logging         *Logging        `json:"logging,omitempty"`
	MacAddress      string          `json:"mac_address,omitempty"`
	NetworkMode     string          `json:"network_mode,omitempty"`
	Networks        interface{}     `json:"networks,omitempty"`
	Pid             interface{}     `json:"pid,omitempty"`
	Ports           []interface{}   `json:"ports,omitempty"`
	Privileged      bool            `json:"privileged,omitempty"`
	ReadOnly        bool            `json:"read_only,omitempty"`
	Restart         string          `json:"restart,omitempty"`
	Secrets         []interface{}   `json:"secrets,omitempty"`
	SecurityOpt     []string        `json:"security_opt,omitempty"`
	ShmSize         interface{}     `json:"shm_size,omitempty"`
	StdinOpen       bool            `json:"stdin_open,omitempty"`
	StopGracePeriod string          `json:"stop_grace_period,omitempty"`
	StopSignal      string          `json:"stop_signal,omitempty"`
	Sysctls         interface{}     `json:"sysctls,omitempty"`
	Tmpfs           interface{}     `json:"tmpfs,omitempty"`
	Tty             bool            `json:"tty,omitempty"`
	Ulimits         *Ulimits        `json:"ulimits,omitempty"`
	User            string          `json:"user,omitempty"`
	UsernsMode      string          `json:"userns_mode,omitempty"`
	Volumes         []interface{}   `json:"volumes,omitempty"`
	WorkingDir      string          `json:"working_dir,omitempty"`
}

// Ulimits
type Ulimits struct {
}

// UpdateConfig
type UpdateConfig struct {
	Delay           string  `json:"delay,omitempty"`
	FailureAction   string  `json:"failure_action,omitempty"`
	MaxFailureRatio float64 `json:"max_failure_ratio,omitempty"`
	Monitor         string  `json:"monitor,omitempty"`
	Order           string  `json:"order,omitempty"`
	Parallelism     int     `json:"parallelism,omitempty"`
}

// Volume
type Volume struct {
	Driver     string      `json:"driver,omitempty"`
	DriverOpts *DriverOpts `json:"driver_opts,omitempty"`
	External   interface{} `json:"external,omitempty"`
	Labels     interface{} `json:"labels,omitempty"`
	Name       string      `json:"name,omitempty"`
}

// Volumes
type Volumes struct {
}

func (strct *ConfigItems) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// Marshal the "subnet" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"subnet\": ")
	if tmp, err := json.Marshal(strct.Subnet); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *ConfigItems) UnmarshalJSON(b []byte) error {
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "subnet":
			if err := json.Unmarshal([]byte(v), &strct.Subnet); err != nil {
				return err
			}
		default:
			return fmt.Errorf("additional property not allowed: \"" + k + "\"")
		}
	}
	return nil
}

func (strct *ConfigFile) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *ConfigFile) UnmarshalJSON(b []byte) error {
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, _ := range jsonMap {
		switch k {
		default:
			return fmt.Errorf("additional property not allowed: \"" + k + "\"")
		}
	}
	return nil
}

func (strct *Deployment) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// Marshal the "endpoint_mode" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"endpoint_mode\": ")
	if tmp, err := json.Marshal(strct.EndpointMode); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "labels" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"labels\": ")
	if tmp, err := json.Marshal(strct.Labels); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "mode" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"mode\": ")
	if tmp, err := json.Marshal(strct.Mode); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "placement" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"placement\": ")
	if tmp, err := json.Marshal(strct.Placement); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "replicas" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"replicas\": ")
	if tmp, err := json.Marshal(strct.Replicas); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "resources" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"resources\": ")
	if tmp, err := json.Marshal(strct.Resources); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "restart_policy" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"restart_policy\": ")
	if tmp, err := json.Marshal(strct.RestartPolicy); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "rollback_config" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"rollback_config\": ")
	if tmp, err := json.Marshal(strct.RollbackConfig); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "update_config" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"update_config\": ")
	if tmp, err := json.Marshal(strct.UpdateConfig); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *Deployment) UnmarshalJSON(b []byte) error {
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "endpoint_mode":
			if err := json.Unmarshal([]byte(v), &strct.EndpointMode); err != nil {
				return err
			}
		case "labels":
			if err := json.Unmarshal([]byte(v), &strct.Labels); err != nil {
				return err
			}
		case "mode":
			if err := json.Unmarshal([]byte(v), &strct.Mode); err != nil {
				return err
			}
		case "placement":
			if err := json.Unmarshal([]byte(v), &strct.Placement); err != nil {
				return err
			}
		case "replicas":
			if err := json.Unmarshal([]byte(v), &strct.Replicas); err != nil {
				return err
			}
		case "resources":
			if err := json.Unmarshal([]byte(v), &strct.Resources); err != nil {
				return err
			}
		case "restart_policy":
			if err := json.Unmarshal([]byte(v), &strct.RestartPolicy); err != nil {
				return err
			}
		case "rollback_config":
			if err := json.Unmarshal([]byte(v), &strct.RollbackConfig); err != nil {
				return err
			}
		case "update_config":
			if err := json.Unmarshal([]byte(v), &strct.UpdateConfig); err != nil {
				return err
			}
		default:
			return fmt.Errorf("additional property not allowed: \"" + k + "\"")
		}
	}
	return nil
}

func (strct *DiscreteResourceSpec) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// Marshal the "kind" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"kind\": ")
	if tmp, err := json.Marshal(strct.Kind); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "value" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"value\": ")
	if tmp, err := json.Marshal(strct.Value); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *DiscreteResourceSpec) UnmarshalJSON(b []byte) error {
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "kind":
			if err := json.Unmarshal([]byte(v), &strct.Kind); err != nil {
				return err
			}
		case "value":
			if err := json.Unmarshal([]byte(v), &strct.Value); err != nil {
				return err
			}
		default:
			return fmt.Errorf("additional property not allowed: \"" + k + "\"")
		}
	}
	return nil
}

func (strct *GenericResourcesItems) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// Marshal the "discrete_resource_spec" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"discrete_resource_spec\": ")
	if tmp, err := json.Marshal(strct.DiscreteResourceSpec); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *GenericResourcesItems) UnmarshalJSON(b []byte) error {
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "discrete_resource_spec":
			if err := json.Unmarshal([]byte(v), &strct.DiscreteResourceSpec); err != nil {
				return err
			}
		default:
			return fmt.Errorf("additional property not allowed: \"" + k + "\"")
		}
	}
	return nil
}

func (strct *Healthcheck) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// Marshal the "disable" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"disable\": ")
	if tmp, err := json.Marshal(strct.Disable); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "interval" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"interval\": ")
	if tmp, err := json.Marshal(strct.Interval); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "retries" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"retries\": ")
	if tmp, err := json.Marshal(strct.Retries); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "start_period" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"start_period\": ")
	if tmp, err := json.Marshal(strct.StartPeriod); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "test" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"test\": ")
	if tmp, err := json.Marshal(strct.Test); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "timeout" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"timeout\": ")
	if tmp, err := json.Marshal(strct.Timeout); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *Healthcheck) UnmarshalJSON(b []byte) error {
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "disable":
			if err := json.Unmarshal([]byte(v), &strct.Disable); err != nil {
				return err
			}
		case "interval":
			if err := json.Unmarshal([]byte(v), &strct.Interval); err != nil {
				return err
			}
		case "retries":
			if err := json.Unmarshal([]byte(v), &strct.Retries); err != nil {
				return err
			}
		case "start_period":
			if err := json.Unmarshal([]byte(v), &strct.StartPeriod); err != nil {
				return err
			}
		case "test":
			if err := json.Unmarshal([]byte(v), &strct.Test); err != nil {
				return err
			}
		case "timeout":
			if err := json.Unmarshal([]byte(v), &strct.Timeout); err != nil {
				return err
			}
		default:
			return fmt.Errorf("additional property not allowed: \"" + k + "\"")
		}
	}
	return nil
}

func (strct *Ipam) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// Marshal the "config" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"config\": ")
	if tmp, err := json.Marshal(strct.Config); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "driver" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"driver\": ")
	if tmp, err := json.Marshal(strct.Driver); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *Ipam) UnmarshalJSON(b []byte) error {
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "config":
			if err := json.Unmarshal([]byte(v), &strct.Config); err != nil {
				return err
			}
		case "driver":
			if err := json.Unmarshal([]byte(v), &strct.Driver); err != nil {
				return err
			}
		default:
			return fmt.Errorf("additional property not allowed: \"" + k + "\"")
		}
	}
	return nil
}

func (strct *Limits) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// Marshal the "cpus" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"cpus\": ")
	if tmp, err := json.Marshal(strct.Cpus); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "memory" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"memory\": ")
	if tmp, err := json.Marshal(strct.Memory); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *Limits) UnmarshalJSON(b []byte) error {
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "cpus":
			if err := json.Unmarshal([]byte(v), &strct.Cpus); err != nil {
				return err
			}
		case "memory":
			if err := json.Unmarshal([]byte(v), &strct.Memory); err != nil {
				return err
			}
		default:
			return fmt.Errorf("additional property not allowed: \"" + k + "\"")
		}
	}
	return nil
}

func (strct *Logging) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// Marshal the "driver" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"driver\": ")
	if tmp, err := json.Marshal(strct.Driver); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "options" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"options\": ")
	if tmp, err := json.Marshal(strct.Options); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *Logging) UnmarshalJSON(b []byte) error {
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "driver":
			if err := json.Unmarshal([]byte(v), &strct.Driver); err != nil {
				return err
			}
		case "options":
			if err := json.Unmarshal([]byte(v), &strct.Options); err != nil {
				return err
			}
		default:
			return fmt.Errorf("additional property not allowed: \"" + k + "\"")
		}
	}
	return nil
}

func (strct *Network) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// Marshal the "attachable" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"attachable\": ")
	if tmp, err := json.Marshal(strct.Attachable); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "driver" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"driver\": ")
	if tmp, err := json.Marshal(strct.Driver); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "driver_opts" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"driver_opts\": ")
	if tmp, err := json.Marshal(strct.DriverOpts); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "external" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"external\": ")
	if tmp, err := json.Marshal(strct.External); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "internal" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"internal\": ")
	if tmp, err := json.Marshal(strct.Internal); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "ipam" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"ipam\": ")
	if tmp, err := json.Marshal(strct.Ipam); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "labels" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"labels\": ")
	if tmp, err := json.Marshal(strct.Labels); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "name" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"name\": ")
	if tmp, err := json.Marshal(strct.Name); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *Network) UnmarshalJSON(b []byte) error {
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "attachable":
			if err := json.Unmarshal([]byte(v), &strct.Attachable); err != nil {
				return err
			}
		case "driver":
			if err := json.Unmarshal([]byte(v), &strct.Driver); err != nil {
				return err
			}
		case "driver_opts":
			if err := json.Unmarshal([]byte(v), &strct.DriverOpts); err != nil {
				return err
			}
		case "external":
			if err := json.Unmarshal([]byte(v), &strct.External); err != nil {
				return err
			}
		case "internal":
			if err := json.Unmarshal([]byte(v), &strct.Internal); err != nil {
				return err
			}
		case "ipam":
			if err := json.Unmarshal([]byte(v), &strct.Ipam); err != nil {
				return err
			}
		case "labels":
			if err := json.Unmarshal([]byte(v), &strct.Labels); err != nil {
				return err
			}
		case "name":
			if err := json.Unmarshal([]byte(v), &strct.Name); err != nil {
				return err
			}
		default:
			return fmt.Errorf("additional property not allowed: \"" + k + "\"")
		}
	}
	return nil
}

func (strct *Placement) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// Marshal the "constraints" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"constraints\": ")
	if tmp, err := json.Marshal(strct.Constraints); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "preferences" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"preferences\": ")
	if tmp, err := json.Marshal(strct.Preferences); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *Placement) UnmarshalJSON(b []byte) error {
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "constraints":
			if err := json.Unmarshal([]byte(v), &strct.Constraints); err != nil {
				return err
			}
		case "preferences":
			if err := json.Unmarshal([]byte(v), &strct.Preferences); err != nil {
				return err
			}
		default:
			return fmt.Errorf("additional property not allowed: \"" + k + "\"")
		}
	}
	return nil
}

func (strct *PreferencesItems) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// Marshal the "spread" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"spread\": ")
	if tmp, err := json.Marshal(strct.Spread); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *PreferencesItems) UnmarshalJSON(b []byte) error {
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "spread":
			if err := json.Unmarshal([]byte(v), &strct.Spread); err != nil {
				return err
			}
		default:
			return fmt.Errorf("additional property not allowed: \"" + k + "\"")
		}
	}
	return nil
}

func (strct *Reservations) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// Marshal the "cpus" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"cpus\": ")
	if tmp, err := json.Marshal(strct.Cpus); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "generic_resources" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"generic_resources\": ")
	if tmp, err := json.Marshal(strct.GenericResources); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "memory" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"memory\": ")
	if tmp, err := json.Marshal(strct.Memory); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *Reservations) UnmarshalJSON(b []byte) error {
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "cpus":
			if err := json.Unmarshal([]byte(v), &strct.Cpus); err != nil {
				return err
			}
		case "generic_resources":
			if err := json.Unmarshal([]byte(v), &strct.GenericResources); err != nil {
				return err
			}
		case "memory":
			if err := json.Unmarshal([]byte(v), &strct.Memory); err != nil {
				return err
			}
		default:
			return fmt.Errorf("additional property not allowed: \"" + k + "\"")
		}
	}
	return nil
}

func (strct *Resources) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// Marshal the "limits" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"limits\": ")
	if tmp, err := json.Marshal(strct.Limits); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "reservations" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"reservations\": ")
	if tmp, err := json.Marshal(strct.Reservations); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *Resources) UnmarshalJSON(b []byte) error {
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "limits":
			if err := json.Unmarshal([]byte(v), &strct.Limits); err != nil {
				return err
			}
		case "reservations":
			if err := json.Unmarshal([]byte(v), &strct.Reservations); err != nil {
				return err
			}
		default:
			return fmt.Errorf("additional property not allowed: \"" + k + "\"")
		}
	}
	return nil
}

func (strct *RestartPolicy) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// Marshal the "condition" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"condition\": ")
	if tmp, err := json.Marshal(strct.Condition); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "delay" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"delay\": ")
	if tmp, err := json.Marshal(strct.Delay); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "max_attempts" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"max_attempts\": ")
	if tmp, err := json.Marshal(strct.MaxAttempts); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "window" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"window\": ")
	if tmp, err := json.Marshal(strct.Window); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *RestartPolicy) UnmarshalJSON(b []byte) error {
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "condition":
			if err := json.Unmarshal([]byte(v), &strct.Condition); err != nil {
				return err
			}
		case "delay":
			if err := json.Unmarshal([]byte(v), &strct.Delay); err != nil {
				return err
			}
		case "max_attempts":
			if err := json.Unmarshal([]byte(v), &strct.MaxAttempts); err != nil {
				return err
			}
		case "window":
			if err := json.Unmarshal([]byte(v), &strct.Window); err != nil {
				return err
			}
		default:
			return fmt.Errorf("additional property not allowed: \"" + k + "\"")
		}
	}
	return nil
}

func (strct *RollbackConfig) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// Marshal the "delay" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"delay\": ")
	if tmp, err := json.Marshal(strct.Delay); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "failure_action" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"failure_action\": ")
	if tmp, err := json.Marshal(strct.FailureAction); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "max_failure_ratio" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"max_failure_ratio\": ")
	if tmp, err := json.Marshal(strct.MaxFailureRatio); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "monitor" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"monitor\": ")
	if tmp, err := json.Marshal(strct.Monitor); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "order" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"order\": ")
	if tmp, err := json.Marshal(strct.Order); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "parallelism" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"parallelism\": ")
	if tmp, err := json.Marshal(strct.Parallelism); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *RollbackConfig) UnmarshalJSON(b []byte) error {
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "delay":
			if err := json.Unmarshal([]byte(v), &strct.Delay); err != nil {
				return err
			}
		case "failure_action":
			if err := json.Unmarshal([]byte(v), &strct.FailureAction); err != nil {
				return err
			}
		case "max_failure_ratio":
			if err := json.Unmarshal([]byte(v), &strct.MaxFailureRatio); err != nil {
				return err
			}
		case "monitor":
			if err := json.Unmarshal([]byte(v), &strct.Monitor); err != nil {
				return err
			}
		case "order":
			if err := json.Unmarshal([]byte(v), &strct.Order); err != nil {
				return err
			}
		case "parallelism":
			if err := json.Unmarshal([]byte(v), &strct.Parallelism); err != nil {
				return err
			}
		default:
			return fmt.Errorf("additional property not allowed: \"" + k + "\"")
		}
	}
	return nil
}

func (strct *Config) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// Marshal the "configs" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"configs\": ")
	if tmp, err := json.Marshal(strct.Configs); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "networks" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"networks\": ")
	if tmp, err := json.Marshal(strct.Networks); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "secrets" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"secrets\": ")
	if tmp, err := json.Marshal(strct.Secrets); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "services" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"services\": ")
	if tmp, err := json.Marshal(strct.Services); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "Version" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "version" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"version\": ")
	buf.Write([]byte(fmt.Sprintf("%d", 3)))
	comma = true
	// Marshal the "volumes" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"volumes\": ")
	if tmp, err := json.Marshal(strct.Volumes); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *Config) UnmarshalJSON(b []byte) error {
	var version string
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "configs":
			if err := json.Unmarshal([]byte(v), &strct.Configs); err != nil {
				return err
			}
		case "networks":
			if err := json.Unmarshal([]byte(v), &strct.Networks); err != nil {
				return err
			}
		case "secrets":
			if err := json.Unmarshal([]byte(v), &strct.Secrets); err != nil {
				return err
			}
		case "services":
			if err := json.Unmarshal([]byte(v), &strct.Services); err != nil {
				return err
			}
		case "version":
			version = string(v)
		case "volumes":
			if err := json.Unmarshal([]byte(v), &strct.Volumes); err != nil {
				return err
			}
		default:
			return fmt.Errorf("additional property not allowed: \"" + k + "\"")
		}
	}
	// check if version (a required property) was received
	if version != "3" {
		return fmt.Errorf("expected \"version\" 3 but found %v", version)
	}
	return nil
}

func (strct *Secret) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// Marshal the "external" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"external\": ")
	if tmp, err := json.Marshal(strct.External); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "file" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"file\": ")
	if tmp, err := json.Marshal(strct.File); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "labels" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"labels\": ")
	if tmp, err := json.Marshal(strct.Labels); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "name" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"name\": ")
	if tmp, err := json.Marshal(strct.Name); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *Secret) UnmarshalJSON(b []byte) error {
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "external":
			if err := json.Unmarshal([]byte(v), &strct.External); err != nil {
				return err
			}
		case "file":
			if err := json.Unmarshal([]byte(v), &strct.File); err != nil {
				return err
			}
		case "labels":
			if err := json.Unmarshal([]byte(v), &strct.Labels); err != nil {
				return err
			}
		case "name":
			if err := json.Unmarshal([]byte(v), &strct.Name); err != nil {
				return err
			}
		default:
			return fmt.Errorf("additional property not allowed: \"" + k + "\"")
		}
	}
	return nil
}

func (strct *Service) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// Marshal the "build" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"build\": ")
	if tmp, err := json.Marshal(strct.Build); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "cap_add" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"cap_add\": ")
	if tmp, err := json.Marshal(strct.CapAdd); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "cap_drop" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"cap_drop\": ")
	if tmp, err := json.Marshal(strct.CapDrop); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "cgroup_parent" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"cgroup_parent\": ")
	if tmp, err := json.Marshal(strct.CgroupParent); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "command" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"command\": ")
	if tmp, err := json.Marshal(strct.Command); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "configs" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"configs\": ")
	if tmp, err := json.Marshal(strct.Configs); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "container_name" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"container_name\": ")
	if tmp, err := json.Marshal(strct.ContainerName); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "credential_spec" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"credential_spec\": ")
	if tmp, err := json.Marshal(strct.CredentialSpec); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "depends_on" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"depends_on\": ")
	if tmp, err := json.Marshal(strct.DependsOn); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "deploy" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"deploy\": ")
	if tmp, err := json.Marshal(strct.Deploy); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "devices" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"devices\": ")
	if tmp, err := json.Marshal(strct.Devices); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "dns" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"dns\": ")
	if tmp, err := json.Marshal(strct.Dns); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "dns_search" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"dns_search\": ")
	if tmp, err := json.Marshal(strct.DnsSearch); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "domainname" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"domainname\": ")
	if tmp, err := json.Marshal(strct.Domainname); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "entrypoint" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"entrypoint\": ")
	if tmp, err := json.Marshal(strct.Entrypoint); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "env_file" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"env_file\": ")
	if tmp, err := json.Marshal(strct.EnvFile); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "environment" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"environment\": ")
	if tmp, err := json.Marshal(strct.Environment); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "expose" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"expose\": ")
	if tmp, err := json.Marshal(strct.Expose); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "external_links" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"external_links\": ")
	if tmp, err := json.Marshal(strct.ExternalLinks); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "extra_hosts" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"extra_hosts\": ")
	if tmp, err := json.Marshal(strct.ExtraHosts); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "healthcheck" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"healthcheck\": ")
	if tmp, err := json.Marshal(strct.Healthcheck); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "hostname" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"hostname\": ")
	if tmp, err := json.Marshal(strct.Hostname); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "image" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"image\": ")
	if tmp, err := json.Marshal(strct.Image); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "init" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"init\": ")
	if tmp, err := json.Marshal(strct.Init); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "ipc" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"ipc\": ")
	if tmp, err := json.Marshal(strct.Ipc); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "isolation" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"isolation\": ")
	if tmp, err := json.Marshal(strct.Isolation); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "labels" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"labels\": ")
	if tmp, err := json.Marshal(strct.Labels); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "links" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"links\": ")
	if tmp, err := json.Marshal(strct.Links); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "logging" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"logging\": ")
	if tmp, err := json.Marshal(strct.Logging); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "mac_address" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"mac_address\": ")
	if tmp, err := json.Marshal(strct.MacAddress); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "network_mode" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"network_mode\": ")
	if tmp, err := json.Marshal(strct.NetworkMode); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "networks" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"networks\": ")
	if tmp, err := json.Marshal(strct.Networks); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "pid" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"pid\": ")
	if tmp, err := json.Marshal(strct.Pid); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "ports" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"ports\": ")
	if tmp, err := json.Marshal(strct.Ports); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "privileged" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"privileged\": ")
	if tmp, err := json.Marshal(strct.Privileged); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "read_only" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"read_only\": ")
	if tmp, err := json.Marshal(strct.ReadOnly); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "restart" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"restart\": ")
	if tmp, err := json.Marshal(strct.Restart); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "secrets" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"secrets\": ")
	if tmp, err := json.Marshal(strct.Secrets); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "security_opt" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"security_opt\": ")
	if tmp, err := json.Marshal(strct.SecurityOpt); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "shm_size" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"shm_size\": ")
	if tmp, err := json.Marshal(strct.ShmSize); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "stdin_open" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"stdin_open\": ")
	if tmp, err := json.Marshal(strct.StdinOpen); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "stop_grace_period" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"stop_grace_period\": ")
	if tmp, err := json.Marshal(strct.StopGracePeriod); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "stop_signal" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"stop_signal\": ")
	if tmp, err := json.Marshal(strct.StopSignal); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "sysctls" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"sysctls\": ")
	if tmp, err := json.Marshal(strct.Sysctls); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "tmpfs" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"tmpfs\": ")
	if tmp, err := json.Marshal(strct.Tmpfs); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "tty" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"tty\": ")
	if tmp, err := json.Marshal(strct.Tty); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "ulimits" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"ulimits\": ")
	if tmp, err := json.Marshal(strct.Ulimits); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "user" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"user\": ")
	if tmp, err := json.Marshal(strct.User); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "userns_mode" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"userns_mode\": ")
	if tmp, err := json.Marshal(strct.UsernsMode); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "volumes" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"volumes\": ")
	if tmp, err := json.Marshal(strct.Volumes); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "working_dir" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"working_dir\": ")
	if tmp, err := json.Marshal(strct.WorkingDir); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *Service) UnmarshalJSON(b []byte) error {
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "build":
			if err := json.Unmarshal([]byte(v), &strct.Build); err != nil {
				return err
			}
		case "cap_add":
			if err := json.Unmarshal([]byte(v), &strct.CapAdd); err != nil {
				return err
			}
		case "cap_drop":
			if err := json.Unmarshal([]byte(v), &strct.CapDrop); err != nil {
				return err
			}
		case "cgroup_parent":
			if err := json.Unmarshal([]byte(v), &strct.CgroupParent); err != nil {
				return err
			}
		case "command":
			if err := json.Unmarshal([]byte(v), &strct.Command); err != nil {
				return err
			}
		case "configs":
			if err := json.Unmarshal([]byte(v), &strct.Configs); err != nil {
				return err
			}
		case "container_name":
			if err := json.Unmarshal([]byte(v), &strct.ContainerName); err != nil {
				return err
			}
		case "credential_spec":
			if err := json.Unmarshal([]byte(v), &strct.CredentialSpec); err != nil {
				return err
			}
		case "depends_on":
			if err := json.Unmarshal([]byte(v), &strct.DependsOn); err != nil {
				return err
			}
		case "deploy":
			if err := json.Unmarshal([]byte(v), &strct.Deploy); err != nil {
				return err
			}
		case "devices":
			if err := json.Unmarshal([]byte(v), &strct.Devices); err != nil {
				return err
			}
		case "dns":
			if err := json.Unmarshal([]byte(v), &strct.Dns); err != nil {
				return err
			}
		case "dns_search":
			if err := json.Unmarshal([]byte(v), &strct.DnsSearch); err != nil {
				return err
			}
		case "domainname":
			if err := json.Unmarshal([]byte(v), &strct.Domainname); err != nil {
				return err
			}
		case "entrypoint":
			if err := json.Unmarshal([]byte(v), &strct.Entrypoint); err != nil {
				return err
			}
		case "env_file":
			if err := json.Unmarshal([]byte(v), &strct.EnvFile); err != nil {
				return err
			}
		case "environment":
			if err := json.Unmarshal([]byte(v), &strct.Environment); err != nil {
				return err
			}
		case "expose":
			if err := json.Unmarshal([]byte(v), &strct.Expose); err != nil {
				return err
			}
		case "external_links":
			if err := json.Unmarshal([]byte(v), &strct.ExternalLinks); err != nil {
				return err
			}
		case "extra_hosts":
			if err := json.Unmarshal([]byte(v), &strct.ExtraHosts); err != nil {
				return err
			}
		case "healthcheck":
			if err := json.Unmarshal([]byte(v), &strct.Healthcheck); err != nil {
				return err
			}
		case "hostname":
			if err := json.Unmarshal([]byte(v), &strct.Hostname); err != nil {
				return err
			}
		case "image":
			if err := json.Unmarshal([]byte(v), &strct.Image); err != nil {
				return err
			}
		case "init":
			if err := json.Unmarshal([]byte(v), &strct.Init); err != nil {
				return err
			}
		case "ipc":
			if err := json.Unmarshal([]byte(v), &strct.Ipc); err != nil {
				return err
			}
		case "isolation":
			if err := json.Unmarshal([]byte(v), &strct.Isolation); err != nil {
				return err
			}
		case "labels":
			if err := json.Unmarshal([]byte(v), &strct.Labels); err != nil {
				return err
			}
		case "links":
			if err := json.Unmarshal([]byte(v), &strct.Links); err != nil {
				return err
			}
		case "logging":
			if err := json.Unmarshal([]byte(v), &strct.Logging); err != nil {
				return err
			}
		case "mac_address":
			if err := json.Unmarshal([]byte(v), &strct.MacAddress); err != nil {
				return err
			}
		case "network_mode":
			if err := json.Unmarshal([]byte(v), &strct.NetworkMode); err != nil {
				return err
			}
		case "networks":
			if err := json.Unmarshal([]byte(v), &strct.Networks); err != nil {
				return err
			}
		case "pid":
			if err := json.Unmarshal([]byte(v), &strct.Pid); err != nil {
				return err
			}
		case "ports":
			if err := json.Unmarshal([]byte(v), &strct.Ports); err != nil {
				return err
			}
		case "privileged":
			if err := json.Unmarshal([]byte(v), &strct.Privileged); err != nil {
				return err
			}
		case "read_only":
			if err := json.Unmarshal([]byte(v), &strct.ReadOnly); err != nil {
				return err
			}
		case "restart":
			if err := json.Unmarshal([]byte(v), &strct.Restart); err != nil {
				return err
			}
		case "secrets":
			if err := json.Unmarshal([]byte(v), &strct.Secrets); err != nil {
				return err
			}
		case "security_opt":
			if err := json.Unmarshal([]byte(v), &strct.SecurityOpt); err != nil {
				return err
			}
		case "shm_size":
			if err := json.Unmarshal([]byte(v), &strct.ShmSize); err != nil {
				return err
			}
		case "stdin_open":
			if err := json.Unmarshal([]byte(v), &strct.StdinOpen); err != nil {
				return err
			}
		case "stop_grace_period":
			if err := json.Unmarshal([]byte(v), &strct.StopGracePeriod); err != nil {
				return err
			}
		case "stop_signal":
			if err := json.Unmarshal([]byte(v), &strct.StopSignal); err != nil {
				return err
			}
		case "sysctls":
			if err := json.Unmarshal([]byte(v), &strct.Sysctls); err != nil {
				return err
			}
		case "tmpfs":
			if err := json.Unmarshal([]byte(v), &strct.Tmpfs); err != nil {
				return err
			}
		case "tty":
			if err := json.Unmarshal([]byte(v), &strct.Tty); err != nil {
				return err
			}
		case "ulimits":
			if err := json.Unmarshal([]byte(v), &strct.Ulimits); err != nil {
				return err
			}
		case "user":
			if err := json.Unmarshal([]byte(v), &strct.User); err != nil {
				return err
			}
		case "userns_mode":
			if err := json.Unmarshal([]byte(v), &strct.UsernsMode); err != nil {
				return err
			}
		case "volumes":
			if err := json.Unmarshal([]byte(v), &strct.Volumes); err != nil {
				return err
			}
		case "working_dir":
			if err := json.Unmarshal([]byte(v), &strct.WorkingDir); err != nil {
				return err
			}
		default:
			return fmt.Errorf("additional property not allowed: \"" + k + "\"")
		}
	}
	return nil
}

func (strct *UpdateConfig) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// Marshal the "delay" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"delay\": ")
	if tmp, err := json.Marshal(strct.Delay); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "failure_action" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"failure_action\": ")
	if tmp, err := json.Marshal(strct.FailureAction); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "max_failure_ratio" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"max_failure_ratio\": ")
	if tmp, err := json.Marshal(strct.MaxFailureRatio); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "monitor" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"monitor\": ")
	if tmp, err := json.Marshal(strct.Monitor); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "order" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"order\": ")
	if tmp, err := json.Marshal(strct.Order); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "parallelism" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"parallelism\": ")
	if tmp, err := json.Marshal(strct.Parallelism); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *UpdateConfig) UnmarshalJSON(b []byte) error {
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "delay":
			if err := json.Unmarshal([]byte(v), &strct.Delay); err != nil {
				return err
			}
		case "failure_action":
			if err := json.Unmarshal([]byte(v), &strct.FailureAction); err != nil {
				return err
			}
		case "max_failure_ratio":
			if err := json.Unmarshal([]byte(v), &strct.MaxFailureRatio); err != nil {
				return err
			}
		case "monitor":
			if err := json.Unmarshal([]byte(v), &strct.Monitor); err != nil {
				return err
			}
		case "order":
			if err := json.Unmarshal([]byte(v), &strct.Order); err != nil {
				return err
			}
		case "parallelism":
			if err := json.Unmarshal([]byte(v), &strct.Parallelism); err != nil {
				return err
			}
		default:
			return fmt.Errorf("additional property not allowed: \"" + k + "\"")
		}
	}
	return nil
}

func (strct *Volume) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// Marshal the "driver" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"driver\": ")
	if tmp, err := json.Marshal(strct.Driver); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "driver_opts" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"driver_opts\": ")
	if tmp, err := json.Marshal(strct.DriverOpts); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "external" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"external\": ")
	if tmp, err := json.Marshal(strct.External); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "labels" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"labels\": ")
	if tmp, err := json.Marshal(strct.Labels); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "name" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"name\": ")
	if tmp, err := json.Marshal(strct.Name); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *Volume) UnmarshalJSON(b []byte) error {
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "driver":
			if err := json.Unmarshal([]byte(v), &strct.Driver); err != nil {
				return err
			}
		case "driver_opts":
			if err := json.Unmarshal([]byte(v), &strct.DriverOpts); err != nil {
				return err
			}
		case "external":
			if err := json.Unmarshal([]byte(v), &strct.External); err != nil {
				return err
			}
		case "labels":
			if err := json.Unmarshal([]byte(v), &strct.Labels); err != nil {
				return err
			}
		case "name":
			if err := json.Unmarshal([]byte(v), &strct.Name); err != nil {
				return err
			}
		default:
			return fmt.Errorf("additional property not allowed: \"" + k + "\"")
		}
	}
	return nil
}

func (strct *Volumes) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *Volumes) UnmarshalJSON(b []byte) error {
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, _ := range jsonMap {
		switch k {
		default:
			return fmt.Errorf("additional property not allowed: \"" + k + "\"")
		}
	}
	return nil
}
