package service

import (
	"test/internal/models/save"
	"test/internal/services/repo"

)


type URLService struct{
	repo repos.URL
}

func NewUrlService(repo repos.URL) *URLService{
	return &URLService{repo: repo}
}

func (s *URLService) GetURLService(URLService int)([]models.Request, error){
	return s.repo.GetURL(URLService)
}

func (s *URLService) CreateURLService(URLService *models.Request)(*models.Request, error){
	return s.repo.SaveURL(URLService)
}