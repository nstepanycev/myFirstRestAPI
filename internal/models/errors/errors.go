package models

import (
	"errors"
)
type ResponseErrors struct{
	Status string `json:"status"`
	Error string  `json:"error,omitempty"`
}

const (
	StatusOK = "OK"
	StatusError = "Error"
)

var (
	ErrURLNotFound = errors.New("URL not found")
	ErrURLExist = errors.New("URL exist")
)