package service

import (
	"test/internal/models/save"
	"test/internal/reposytory"

)

type Storage interface{
	CreateURLService(service *models.Request) (*models.Request, error)
	GetURLService(id int) ([]models.Request, error)
}

type Service struct{
	Storage
}

func NewService(repos *repos.Repos) *Service{
	return &Service{
		Storage: NewUrlService(repos.URL),
	}
}