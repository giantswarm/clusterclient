package config

// Patch is the cluster specific configuration.
type Patch struct {
	Owner           string        `json:"owner"`
	Workers         []interface{} `json:"workers"`
	PatchAttributes []string      `json:"patch_attributes"`
}

// DefaultPatch provides a default patch by best effort.
func DefaultPatch() Patch {
	return Patch{
		Owner:           "",
		Workers:         make([]interface{}, 0),
		PatchAttributes: []string{},
	}
}
