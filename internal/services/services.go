package service

import (
	"test/internal/models/save"
	"test/internal/services/repo"

)

type Storage interface{
	CreateURLService(service *models.Request) (*models.Request, error)
	GetURLService(service *models.Request) ([]models.Request, error)
}

type Service struct{
	Storage
}

func NewService(repos *repos.Repos) *Service{
	return &Service{
		Storage: NewUrlService(repos.URL),
	}
}