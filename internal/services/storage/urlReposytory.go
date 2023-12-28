package storage

import (
	"test/internal/models/save"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Storage interface{
	SaveURL(URLSave *models.Request) (int, error)
	GetURL(URLSave *models.Request) (*models.Request, error)
}

type Repos struct{
	Storage
}

func NewReposytory(db *pgxpool.Pool) *Repos{
	return &Repos{
		Storage: NewUrlDatabase(db),
	}
}