package storage

import (
	"test/internal/models/save"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Storage interface{
	SaveURL(URLSave *models.Request) (int, error)
	GetURL(id int) (*models.Request, error)
}

type Service struct{
	Storage
}

func NewService(db *pgxpool.Pool) *Service{
	return &Service{
		Storage: NewClientStorage(db),
	}
}