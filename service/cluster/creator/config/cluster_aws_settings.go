package config

// ClusterAWSSettings configures AWS-specific cluster settings.
type ClusterAWSSettings struct {
	ResourceTags map[string]string `json:"resource_tags"`
}

// DefaultClusterAWSSettings provides a default ClusterAWSSettings.
func DefaultClusterAWSSettings() *ClusterAWSSettings {
	return &ClusterAWSSettings{
		ResourceTags: map[string]string{},
	}
}
