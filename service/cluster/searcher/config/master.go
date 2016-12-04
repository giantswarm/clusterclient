package config

type Master struct {
	CPUs   int               `json:"cpus"`
	Labels map[string]string `json:"labels"`
	Ram    string            `json:"ram"`
}

func DefaultMaster() *Master {
	return &Master{
		CPUs:   0,
		Labels: map[string]string{},
		Ram:    "",
	}
}
