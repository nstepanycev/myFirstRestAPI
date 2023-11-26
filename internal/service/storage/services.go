package storage

import (
	"test/internal/models/save"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Storage interface{
	SaveURL(urlToSave *models.UrlToSaveService) (int, error)
	GetURL(id int) (*models.UrlToSaveService, error)
}

type Service struct{
	Storage
}

func NewReposytory(db *pgxpool.Pool) *Service{
	return &Service{
		Storage: NewClientStorage(db),
	}
}