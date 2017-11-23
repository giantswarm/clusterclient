package response

type InstanceType struct {
	Options []string `json:"options"`
	Default string   `json:"default"`
}

func DefaultInstanceType() InstanceType {
	return InstanceType{
		Options: []string{},
		Default: "",
	}
}
