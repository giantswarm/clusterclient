package config

import (
	"github.com/giantswarm/clusterclient/service/cluster/updater/config/aws"
	"github.com/giantswarm/clusterclient/service/cluster/updater/config/azure"
)

// Worker configures the Kubernetes worker nodes.
type Worker struct {
	CPU     CPU               `json:"cpu"`
	Labels  map[string]string `json:"labels"`
	Memory  Memory            `json:"memory"`
	AWS     aws.Worker        `json:"aws"`
	Azure   azure.Worker      `json:"azure"`
	Volumes Volumes           `json:"volumes"`
}

// DefaultWorker provides a default worker configuration by best effort.
func DefaultWorker() Worker {
	return Worker{
		CPU:     DefaultCPU(),
		Labels:  map[string]string{},
		Memory:  DefaultMemory(),
		AWS:     aws.DefaultWorker(),
		Volumes: DefaultVolumes(),
	}
}
