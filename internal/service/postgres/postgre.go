package postgres

import (
	"context"
	"fmt"
	"test/internal/config"

	"github.com/jackc/pgx/v5/pgxpool"
)




func ConnectToDB(cfg config.StorageConfig) (*pgxpool.Pool, error){
	connectDb := fmt.Sprintf("DbHost=%s DbPort=%s DbUser=%s"+ "DbPass=%s DbName=%s ", 
	cfg.DbHost,cfg.DbPort,cfg.DbUser,cfg.DbPass,cfg.DbName)
	
	db, err := pgxpool.New(context.Background(),connectDb)
	if err != nil{
		fmt.Printf("Unable to connect to database: %v\n", err)
		return nil, err
	}
	fmt.Println("Connect to DB")

	return db, nil
}



