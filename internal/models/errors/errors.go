package models

import (
	"errors"
)


var (
	ErrURLNotFound = errors.New("URL not found")
	ErrURLExist = errors.New("URL exist")
)

type ResponseErrors struct{
	Status string  `json:"status"`
	Error  string  `json:"error,omitempty"`
}

const (
	StatusOK = "OK"
	StatusError = "Error"
)

