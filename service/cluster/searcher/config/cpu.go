package config

// CPU configures the machine CPU.
type CPU struct {
	Cores float64 `json:"cores"`
}

// DefaultCPU provides a default CPU configuration by best effort.
func DefaultCPU() *CPU {
	return &CPU{
		Cores: 0,
	}
}
