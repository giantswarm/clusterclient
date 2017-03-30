package config

// Storage configures the machine storage.
type Storage struct {
	SizeGB int `json:"size_gb"`
}

// DefaultStorage provides a default storage configuration by best effort.
func DefaultStorage() *Storage {
	return &Storage{
		SizeGB: 0,
	}
}
