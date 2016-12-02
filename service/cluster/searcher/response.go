package searcher

import "time"

// Response is the return value of the service action.
type Response struct {
	APIEndpoint string    `json:"api_endpoint"`
	CreateDate  time.Time `json:"create_date"`
	ID          string    `json:"id"`
	Name        string    `json:"name"`
}

// DefaultResponse provides a default response by best effort.
func DefaultResponse() *Response {
	return &Response{
		APIEndpoint: "",
		CreateDate:  time.Time{},
		ID:          "",
		Name:        "",
	}
}
