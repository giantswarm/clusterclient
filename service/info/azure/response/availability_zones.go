package response

type AvailabilityZones struct {
	Default int      `json:"default"`
	Max     int      `json:"max"`
	Zones   []int `json:"zones,omitempty"`
}
