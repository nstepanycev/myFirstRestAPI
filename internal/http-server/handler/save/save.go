package save

import (
	resp "test/internal/models/response"
)

type Request struct{
	URL string `json:"url"`
	Aliase string `json:"aliase,omitempty"`
}

type Response struct{
	resp.Response
	Alias string `json:"alias"`
}