package config

// CPU configures the machine CPU.
type CPU struct {
	Amount int `json:"amount"`
}

// DefaultCPU provides a default CPU configuration by best effort.
func DefaultCPU() *CPU {
	return &CPU{
		Amount: 0,
	}
}
