package config

// Cluster is the cluster data structure.
type Cluster struct {
	ID string
}

// DefaultCluster returns a default Cluster object.
func DefaultCluster() *Cluster {
	return &Cluster{
		ID: "",
	}
}
