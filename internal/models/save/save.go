package models

import (
	resp "test/internal/models/errors"
)

type Request struct{
	URL string `json:"url"`
	Aliase string `json:"aliase,omitempty"`
}

type Response struct{
	resp.ResponseErrors
	Alias string `json:"alias,omitempty"`
}

type UrlToSaveService struct{
	Id        *int   `json:"id, omitempty" db:"id"`
	UrlToSave string `json:"url"`
	Aliace    string `json:"aliase"`
}
