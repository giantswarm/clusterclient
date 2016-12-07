package config

// Storage configures the machine storage.
type Storage struct {
	Size string `json:"size"`
}

// DefaultStorage provides a default storage configuration by best effort.
func DefaultStorage() *Storage {
	return &Storage{
		Size: "",
	}
}
