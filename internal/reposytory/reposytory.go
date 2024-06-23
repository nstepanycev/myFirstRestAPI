package repos

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"test/internal/models/save"
)


type URL interface {
	SaveURL(request *models.Request) (*models.Request, error)
	GetURL(id int) ([]models.Request, error)
}

type Repos struct{
	URL
}

func NewReposytory(db *pgxpool.Pool) *Repos{
	return &Repos{
		URL: NewUrlDatabase(db),
	}
}

