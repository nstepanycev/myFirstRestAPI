package service

import (
	"test/internal/models/save"


)

type URLSave interface{
	GetURL(id int) (*models.UrlToSaveService, error)
	SaveURL(urlToSave *models.Request) (int, error)
}