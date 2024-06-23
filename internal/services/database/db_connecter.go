package database

import (
	"context"
	"fmt"
	"log/slog"
	"test/internal/config"
	"github.com/jackc/pgx/v5/pgxpool"
)



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



