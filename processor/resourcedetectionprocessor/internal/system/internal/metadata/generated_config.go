// Code generated by mdatagen. DO NOT EDIT.

package metadata

// ResourceAttributeConfig provides common config for a particular resource attribute.
type ResourceAttributeConfig struct {
	Enabled bool `mapstructure:"enabled"`
}

// ResourceAttributesConfig provides config for resourcedetectionprocessor/system resource attributes.
type ResourceAttributesConfig struct {
	HostID   ResourceAttributeConfig `mapstructure:"host.id"`
	HostName ResourceAttributeConfig `mapstructure:"host.name"`
	OsType   ResourceAttributeConfig `mapstructure:"os.type"`
}

func DefaultResourceAttributesConfig() ResourceAttributesConfig {
	return ResourceAttributesConfig{
		HostID: ResourceAttributeConfig{
			Enabled: false,
		},
		HostName: ResourceAttributeConfig{
			Enabled: true,
		},
		OsType: ResourceAttributeConfig{
			Enabled: true,
		},
	}
}