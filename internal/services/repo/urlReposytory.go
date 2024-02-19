package repos

import (
	"github.com/jackc/pgx/v5/pgxpool"
)


type Repos struct{
	URL
}

func NewReposytory(db *pgxpool.Pool) *Repos{
	return &Repos{
		URL: NewUrlDatabase(db),
	}
}