// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// ServerApplyConfiguration represents a declarative configuration of the Server type for use
// with apply.
type ServerApplyConfiguration struct {
	Name          *string                          `json:"name,omitempty"`
	Zones         []string                         `json:"zones,omitempty"`
	ForwardPlugin *ForwardPluginApplyConfiguration `json:"forwardPlugin,omitempty"`
}

// ServerApplyConfiguration constructs a declarative configuration of the Server type for use with
// apply.
func Server() *ServerApplyConfiguration {
	return &ServerApplyConfiguration{}
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *ServerApplyConfiguration) WithName(value string) *ServerApplyConfiguration {
	b.Name = &value
	return b
}

// WithZones adds the given value to the Zones field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Zones field.
func (b *ServerApplyConfiguration) WithZones(values ...string) *ServerApplyConfiguration {
	for i := range values {
		b.Zones = append(b.Zones, values[i])
	}
	return b
}

// WithForwardPlugin sets the ForwardPlugin field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ForwardPlugin field is set to the value of the last call.
func (b *ServerApplyConfiguration) WithForwardPlugin(value *ForwardPluginApplyConfiguration) *ServerApplyConfiguration {
	b.ForwardPlugin = value
	return b
}
