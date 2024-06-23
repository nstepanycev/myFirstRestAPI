package repos

import (
	"context"
	"log/slog"
	"test/internal/models/save"
	"github.com/jackc/pgx/v5/pgxpool"
)


type URLStorage struct {
	db *pgxpool.Pool
}

func NewUrlDatabase(db *pgxpool.Pool) *URLStorage {
	return &URLStorage{db: db}
}

func (s *URLStorage) SaveURL(request *models.Request) (*models.Request, error) {
	
	rows, err := s.db.Query(context.Background(),
	`insert into public.url (id,url,alias) values ($1,$2,$3)`,request.Id,request.URL,request.Aliase)
	if err != nil{
		return nil, err
	}
	defer rows.Close()

	return request, nil
}

func (s *URLStorage) GetURL(id int) ([]models.Request, error) {

	rows, err := s.db.Query(context.Background(),`select * from public.url where id = $1`, id)
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
