package models

type GetRequest struct{}

type GetResponse struct {
	Date string `json:"date"`
	Err  error  `json:"err,omitempty"`
}

type ValidateRequest struct {
	Date string `json:"date"`
}
type ValidateResponse struct {
	Valid bool  `json:"valid"`
	Err   error `json:"err,omitempty"`
}

type StatusRequest struct{}

type StatusResponse struct {
	Status string `json:"status"`
}
