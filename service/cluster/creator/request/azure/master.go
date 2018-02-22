package azure

// Master configures Azure-specific master node settings.
type Master struct {
	VmSize string `json:"vm_size"`
}

// DefaultMaster provides default Master.
func DefaultMaster() Master {
	return Master{
		VmSize: "",
	}
}
