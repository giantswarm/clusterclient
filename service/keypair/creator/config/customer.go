package config

type KeyPair struct {
	Description string `json:"description"`
	TTL         int    `json:"ttl"`
}

func DefaultKeyPair() *KeyPair {
	return &KeyPair{
		Description: "",
		TTL:         0,
	}
}
