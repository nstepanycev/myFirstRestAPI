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
	connectDb := fmt.Sprintf("DbHost=%s DbPort=%s DbUser=%s"+ "DbPass=%s DbName=%s ", 
	cfg.DbHost,cfg.DbPort,cfg.DbUser,cfg.DbPass,cfg.DbName)
	
	db, err := pgxpool.New(context.Background(),connectDb)
	if err != nil{
		slog.Info(err.Error())
		return nil, err
	}
	slog.Info("Database conneect success")

	return db, nil
}



