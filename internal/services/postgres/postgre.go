package postgres

import (
	"context"
	"fmt"
	"log/slog"
	"test/internal/config"

	// "test/internal/http-server/middleware/logger"

	"github.com/jackc/pgx/v5/pgxpool"
)

type URLStorage struct{
	db *pgxpool.Pool
}

func NewClientStorage(db *pgxpool.Pool) *URLStorage{
	return &URLStorage{db:db}
}


func ConnectToDB(cfg config.StorageConfig) (*pgxpool.Pool, error){
	connectDb := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
	cfg.DbUser,cfg.DbPass,cfg.DbHost,cfg.DbPort,cfg.DbName)

	db, err := pgxpool.New(context.Background(),connectDb)
	if err != nil{
		slog.Debug("Unable connact", err)
		return nil, err
	}
	slog.Info("Database conneect success")

	return db, nil
}



