package models


type UrlToSaveService struct{
	Id        *int   `json:"id" db:"id"`
	UrlToSave string `json:"url"`
	Aliace    string `json:"aliase"`
}
