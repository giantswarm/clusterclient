package config

// WorkerAWSSettings configures AWS-specific worker node settings.
type WorkerAWSSettings struct {
	InstanceType string `json:"instance_type"`
}

// DefaultWorkerAWSSettings provides default WorkerAWSSettings.
func DefaultWorkerAWSSettings() *WorkerAWSSettings {
	return &WorkerAWSSettings{
		InstanceType: "",
	}
}
