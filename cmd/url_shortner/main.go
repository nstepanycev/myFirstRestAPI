package main

import (
	"context"
	"os"
	"test/internal/config"
	httpserver "test/internal/http-server"
	handler "test/internal/http-server/handler"
	"test/internal/service/postgres"
	"test/internal/service/storage"
	// "github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/exp/slog"
)


const (
	envLocal = "local"
	envDev = "dev"
)



func main(){

	cfg := config.LoadConfig()

	DataBaseConfig := config.StorageConfig{
		DbHost: cfg.StorageConfig.DbHost,
		DbPort: cfg.StorageConfig.DbPort,
		DbName: cfg.StorageConfig.DbName,
		DbUser: cfg.StorageConfig.DbUser,
		DbPass: cfg.StorageConfig.DbPass,
	}
	//Logger
	log := setupLogger(cfg.Env)
	log = log.With(slog.String("env", cfg.Env))

	log.Info("initializing server", slog.String("host", cfg.HTTP_Server.Host))
	log.Debug("Log config enable")
	
	//Database
	db, err := postgres.ConnectToDB(DataBaseConfig) 
	if err != nil{
		slog.String("Err connect to database: %s", err.Error())
	}
	defer func(db *pgxpool.Pool, ctx context.Context){
		db.Close()
	}(db, context.Background())



	// router := gin.New()
	// router.Use(logger.NewMiddleware())
	// ShortnerRoute(router)

	//
	service := storage.NewService(db)
	handlers := handler.NewHandler(service)


	//Router
	srv := new(httpserver.Server)
	if err := srv.Run(cfg.HTTP_Server, handlers.InitRouter()); err != nil{
		slog.String("Error occurred while running http server: %s", err.Error())
	}

}

func setupLogger(env string) *slog.Logger{
	var log *slog.Logger

	switch env{
	case env:
		log = slog.New(slog.NewJSONHandler(os.Stdout,&slog.HandlerOptions{Level: slog.LevelDebug}))
	case envDev:
		log = slog.New(slog.NewJSONHandler(os.Stdout,&slog.HandlerOptions{Level: slog.LevelDebug}))
	}
	return log
}