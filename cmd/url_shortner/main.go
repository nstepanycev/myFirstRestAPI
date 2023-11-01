package main

import (
	"fmt"
	"os"
	"test/internal/config"
	"test/internal/storage/postgres"

	"golang.org/x/exp/slog"
)


const (
	envLocal = "local"
	envDev = "dev"
)



func main(){
	//init config
	cfg := config.MustLoad()
	fmt.Println(cfg)

	s := config.DbConfig()

	DataBaseConfig := config.StorageConfig{
		DbHost: s.DbHost,
		DbPort: s.DbPort,
		DbName: s.DbName,
		DbUser: s.DbUser,
		DbPass: s.DbPass,
	}
	//^TODO Fix

	log := setupLogger(cfg.Env)
	log = log.With(slog.String("env", cfg.Env))

	log.Info("initializing server", slog.String("addres", cfg.ConfigPath))
	log.Debug("Log config enable")
	
	db, err := postgres.ConnectToDB((*postgres.DbConfig)(&DataBaseConfig)) // TO DO Fix Я хз почему такая фигня вложенная разберусь
	if err != nil{
		os.Exit(1)
		log.Error("Error to connect to DB")
	}
	defer db.Close()

	//middleware

	//ruter
	// router := gin.NewRouter()

	// router.Use(middleware.RequestID())
	// router.Use(middleware.Logger())
	// router.Use(middleware.Recoverer())
	// router.Use(middleware.URLFormat())
	// r := gin.New()
	// r.Run(":8080")
	



}
	//init DB	
	//init loger
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
	//init storage

	

	//init router