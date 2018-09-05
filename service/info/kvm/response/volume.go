package response

type Volumes struct {
	Docker Volume `json:"docker"`
}

type Volume struct {
	SizeGB VolumeSize
}

type VolumeSize struct {
	Min     float64
	Max     float64
	Default float64
}

func DefaultVolumes() Volumes {
	return Volumes{
		Docker: Volume{
			SizeGB: VolumeSize{
				Min:     10,
				Max:     10000,
				Default: 100,
			},
		},
	}
}
