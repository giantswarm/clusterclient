package config

// Worker configures the Kubernetes worker nodes.
type Worker struct {
	CPU     *CPU              `json:"cpu"`
	Labels  map[string]string `json:"labels"`
	Ram     *Ram              `json:"ram"`
	Storage *Storage          `json:"storage"`
}

// DefaultWorker provides a default worker configuration by best effort.
func DefaultWorker() *Worker {
	return &Worker{
		CPU:     DefaultCPU(),
		Labels:  map[string]string{},
		Ram:     DefaultRam(),
		Storage: DefaultStorage(),
	}
}
