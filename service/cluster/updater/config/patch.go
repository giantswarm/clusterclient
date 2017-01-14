package config

// Patch is the cluster specific configuration.
type Patch struct {
	Owner string `json:"owner"`
}

// DefaultPatch provides a default patch by best effort.
func DefaultPatch() Patch {
	return Patch{
		Owner: "",
	}
}