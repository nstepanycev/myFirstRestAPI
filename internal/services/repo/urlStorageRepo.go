package repos

import (
	"context"
	"log/slog"
	// "database/sql"
	// "log/slog"
	"test/internal/models/save"

	// "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	urlDb = "url"
)

type URL interface {
	SaveURL(request *models.Request) (*models.Request, error)
	GetURL(request *models.Request) ([]models.Request, error)
}

type URLStorage struct {
	db *pgxpool.Pool
}

func NewUrlDatabase(db *pgxpool.Pool) *URLStorage {
	return &URLStorage{db: db}
}

func (s *URLStorage) SaveURL(request *models.Request) (*models.Request, error) {
	
	rows, err := s.db.Query(context.Background(),
	"inser into public.url (id,alias,url) values ('$1','$2','$3')",request.Id,request.URL,request.Aliase)
	if err != nil{
		return nil, err
	}
	defer rows.Close()


	// query, args, err := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).
	// 	Insert(urlDb).Columns("id", "url", "aliase").
	// 	SetMap(squirrel.Eq{
	// 		"id": request.Id, "url": request.URL, "aliase": request.Aliase,
	// 	}).
	// 	Suffix("RETURNING id,url,aliace").ToSql()

	// if err != nil {
	// 	return nil, err
	// }
	// row := s.db.QueryRow(context.Background(), query, args...)
	// err = row.Scan(
	// 	&request.Id,
	// 	&request.URL,
	// 	&request.Aliase,
	// )
	// if err != nil {
	// 	return nil, err
	// }
	return request, nil
}

func (s *URLStorage) GetURL(request *models.Request) ([]models.Request, error) {
	// var qUrl *models.Request

	// sql, args, err := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).
	// 	Select("id, url, aliase").From(urlDb).Where(squirrel.Eq{"id": qUrl.Id}).ToSql()
	// if err != nil {
	// 	return nil, err
	// }
	// row := s.db.QueryRow(context.Background(), sql, args...)
	// err = row.Scan(&qUrl.Id, &qUrl.URL, &qUrl.Aliase)
	// if err != nil {
	// 	return nil, err
	// }


	rows, err := s.db.Query(context.Background(),"select * from public.url where id = $1", request.Id)
	if err != nil{
		return nil, err
	}
	defer rows.Close()
	qUrl := []models.Request{}

	for rows.Next(){
		u := models.Request{}
		err := rows.Scan(&u.Id, &u.URL, &u.Aliase)
		if err != nil{
			slog.Debug("rows:", err )
			continue
		}
		qUrl = append(qUrl,u)
	}
	return qUrl, nil
}
