package storage

import (
	"context"
	"fmt"
	"os"
	"test/internal/models/save"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	urlDb = "url"
)

type URLStorage struct{
	db *pgxpool.Pool
}

func NewClientStorage(db *pgxpool.Pool) *URLStorage{
	return &URLStorage{db:db}
}

func (s *URLStorage) SaveURL(request *models.Request) (int, error){
	query := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).
		Insert(urlDb).Columns("url", "aliase").
		Values(request.URL, request.Aliase).Suffix("RETURNING id")
	sql, args, err := query.ToSql()
	if err != nil{
		return 0, err
	}
	var id int
	row := s.db.QueryRow(context.Background(), sql, args...)
	err = row.Scan(&id)
	if err != nil{
		return 0, nil
	}
	return id, nil
}

func (s *URLStorage) GetURL(id int) (*models.Request, error){
	var qUrl *models.Request
	query := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).
		Select("id, url, aliase").From(urlDb).Where(squirrel.Eq{"id":id})
	sql, args, err := query.ToSql()
	if err != nil{
		return qUrl, err
	}
	row := s.db.QueryRow(context.Background(), sql, args...)
	err = row.Scan(&qUrl.Id, &qUrl.URL, &qUrl.Aliase)
	if err != nil{
		fmt.Println(err)
		os.Exit(1)
	}
	return qUrl, nil
}