package postgres

import (
	"context"
	"fmt"
	"os"
	"github.com/jackc/pgx/v5/pgxpool"
)


const (
	urlDb = "url"
)

type DbConfig struct {
	DbHost string
	DbPort string
	DbName string
	DbUser string
	DbPass string
}


func ConnectToDB(cfg *DbConfig) (*pgxpool.Pool, error){
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



