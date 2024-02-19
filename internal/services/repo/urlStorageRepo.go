package repos

import (
	"context"
	"test/internal/models/save"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgxpool"
)


const (
	urlDb = "url"
)
type URL interface{
	SaveURL(request *models.Request) (*models.Request, error)
	GetURL(URLId int) (*models.Request, error)
}

type URLStorage struct{
	db *pgxpool.Pool
}

func NewUrlDatabase(db *pgxpool.Pool) *URLStorage{
	return &URLStorage{db:db}
}


func (s *URLStorage) SaveURL(request *models.Request) (*models.Request, error){
	query, args, err := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).
		Insert(urlDb).Columns("id","url", "aliase").
		SetMap(squirrel.Eq{
			"id":request.Id, "url":request.URL, "aliase":request.Aliase,
		}).
		Suffix("RETURNING id,url,aliace").ToSql()
		
		if err != nil{
		return nil, err
		}
	row := s.db.QueryRow(context.Background(), query, args...)
	err = row.Scan(
		&request.Id,
		&request.URL,
		&request.Aliase,
	)
	if err != nil{
		return nil, err
	}
	return request, nil
}

func (s *URLStorage) GetURL(URLId int) (*models.Request, error){
	var qUrl *models.Request

	sql, args, err := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).
		Select("id, url, aliase").From(urlDb).Where(squirrel.Eq{"id":qUrl.Id}).ToSql()
	if err != nil{
		return nil, err
	}
	row := s.db.QueryRow(context.Background(), sql, args...)
	err = row.Scan(&qUrl.Id, &qUrl.URL, &qUrl.Aliase)
	if err != nil{
		return nil, err
	}
	return qUrl, nil
}