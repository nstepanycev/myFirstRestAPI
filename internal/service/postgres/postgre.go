package postgres

import (
	"context"
	"fmt"
	"os"
	"test/internal/config"

	"github.com/jackc/pgx/v5/pgxpool"
)


const (
	urlDb = "url"
)



func ConnectToDB(cfg config.StorageConfig) (*pgxpool.Pool, error){
	connectDb := fmt.Sprintf("DbHost=%s DbPort=%s DbUser=%s"+ "DbPass=%s DbName=%s ", 
	cfg.DbHost,cfg.DbPort,cfg.DbUser,cfg.DbPass,cfg.DbName)
	
	db, err := pgxpool.New(context.Background(),connectDb)
	if err != nil{
		fmt.Printf("Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Connect to DB")

	return db, nil
}



