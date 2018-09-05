package config

// Volumes defines all disks attached to the node
type Volumes struct {
	Docker Volume `json:"docker"`
}

// Volume defines a disk attached to the node
type Volume struct {
	SizeGB float64 `json:"size_gb"`
}

// DefaultVolumes provides a default volumes configuration for the nodes
func DefaultVolumes() Volumes {
	return Volumes{
		Docker: Volume{
			SizeGB: 100,
		},
	}
}
