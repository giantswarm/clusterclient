package config

type Worker struct {
	CPUs   int               `json:"cpus"`
	Labels map[string]string `json:"labels"`
	Ram    string            `json:"ram"`
}

func DefaultWorker() *Worker {
	return &Worker{
		CPUs:   0,
		Labels: map[string]string{},
		Ram:    "",
	}
}
