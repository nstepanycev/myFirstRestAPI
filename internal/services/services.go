package service

import (
	"test/internal/models/save"
	"test/internal/services/storage"

)

type Storage interface{
	CreateURLService(service *models.Request) (int, error)
	GetURLService(service *models.Request) (*models.Request, error)
}

type Service struct{
	Storage
}

func NewService(repos *storage.Repos) *Service{
	return &Service{
		Storage: NewUrlService(repos),
	}
}