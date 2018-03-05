package kvm

// Cluster configures KVM-specific cluster settings.
type Cluster struct {
	PortMappings []ProtocolPort `json:"port_mappings,omitempty"`
}

// DefaultCluster provides a default Cluster.
func DefaultCluster() Cluster {
	return Cluster{
		PortMappings: []ProtocolPort{},
	}
}
