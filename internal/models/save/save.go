package models

import (
	resp "test/internal/models/errors"
)

type Request struct{
	Id     int    `json:"id,omitempty" db:"id,omitempty"`
	URL    string `json:"url" db:"url"`
	Aliase string `json:"aliase,omitempty" db:"aliase"`
}

type Response struct{
	resp.ResponseErrors
	Alias string `json:"alias,omitempty"`
}
