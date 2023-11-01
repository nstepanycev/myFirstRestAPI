package service

import (
	"test/internal/models"


)

// type Service struct{
// 	UrlShortner
// }

type URLSave interface{
	GetURL(id int) (*models.UrlToSaveService, error)
	SaveURL(urlToSave *models.UrlToSaveService) (int, error)
}