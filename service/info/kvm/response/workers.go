package response

type Workers struct {
	CountPerCluster CountPerCluster `json:"count_per_cluster"`

	// To be implemented:
	// cpu_cores (KVM Specific)
	// ram_size_gb (KVM Specific)

	Volumes Volumes
}

func DefaultWorkers() Workers {
	return Workers{
		CountPerCluster: DefaultCountPerCluster(),
		Volumes:         DefaultVolumes(),
	}
}
