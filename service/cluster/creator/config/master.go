package config

// Master configures the Kubernetes master nodes.
type Master struct {
	CPU     *CPU              `json:"cpu"`
	Labels  map[string]string `json:"labels"`
	Ram     *Ram              `json:"ram"`
	Storage *Storage          `json:"storage"`
}

// DefaultMaster provides a default master configuration by best effort.
func DefaultMaster() *Master {
	return &Master{
		CPU:     DefaultCPU(),
		Labels:  map[string]string{},
		Ram:     DefaultRam(),
		Storage: DefaultStorage(),
	}
}
