package config

// Ram configures the machine memory.
type Ram struct {
	Size string `json:"size"`
}

// DefaultRam provides a default ram configuration by best effort.
func DefaultRam() *Ram {
	return &Ram{
		Size: "",
	}
}
