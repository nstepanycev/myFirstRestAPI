package service

import (
	"test/internal/models/save"
	"test/internal/services/storage"
)


type URLService struct{
	repos *storage.Repos
}

func NewUrlService(repos *storage.Repos) *URLService{
	return &URLService{repos: repos}
}

func (s *URLService) GetURLService(URLService *models.Request)(*models.Request, error){
	return s.repos.GetURL(URLService)
}

func (s *URLService) CreateURLService(URLService *models.Request)(int, error){
	return s.repos.SaveURL(URLService)
}