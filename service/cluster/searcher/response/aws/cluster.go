package aws

// Cluster configures AWS-specific cluster settings.
type Cluster struct {
	AvailabilityZones []string          `json:"availability_zones,omitempty"`
	ResourceTags      map[string]string `json:"resource_tags"`
}

// DefaultCluster provides a default Cluster.
func DefaultCluster() Cluster {
	return Cluster{
		AvailabilityZones: []string{},
		ResourceTags:      map[string]string{},
	}
}
