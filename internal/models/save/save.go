package models

import (
	resp "test/internal/models/errors"
)

type Request struct{
	Id     *int   `json:"id, omitempty" db:"id"`
	URL    string `json:"url"`
	Aliase string `json:"aliase,omitempty"`
}

type Response struct{
	resp.ResponseErrors
	Alias string `json:"alias,omitempty"`
}
