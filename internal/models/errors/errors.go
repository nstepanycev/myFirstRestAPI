package models


type ResponseErrors struct{
	Status string `json:"status"`
	Error string  `json:"error,omitempty"`
}

const (
	StatusOK = "OK"
	StatusError = "Error"
)