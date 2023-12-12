package main

import (
	"fmt"
	"os"
	"test/internal/config"
	httpserver "test/internal/http-server"
	handler "test/internal/http-server/handler"

	// "test/internal/http-server/middleware"
	// "test/internal/http-server/middleware/logger"
	"test/internal/service/storage"

	// "test/internal/http-server/middleware/logger"
	"test/internal/service/postgres"

	// "github.com/gin-gonic/gin"
	"golang.org/x/exp/slog"
)


const (
	envLocal = "local"
	envDev = "dev"
)



func main(){

	cfg := config.LoadConfig()
	fmt.Println(cfg)

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

	log.Info("initializing server", slog.String("addres", cfg.ConfigPath))
	log.Debug("Log config enable")
	
	//Database
	db, err := postgres.ConnectToDB(DataBaseConfig) 
	if err != nil{
		os.Exit(1)
		log.Error("Error to connect to DB")
	}
	defer db.Close()

	//
	service := storage.NewService(db)
	handlers := handler.NewHandler(service)


	//Router
	srv := new(httpserver.Server)
	if err := srv.Run(config.HTTPServer{}, handlers.InitRouter()); err != nil{ // fix config
		fmt.Println(err)
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