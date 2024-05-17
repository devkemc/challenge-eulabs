package main

import (
	"database/sql"
	"github.com/devkemc/challenge-eulabs/internal/server/http"
	config2 "github.com/devkemc/challenge-eulabs/pkg/config"
	"github.com/devkemc/challenge-eulabs/pkg/db/product"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	config := config2.LoadConfig()
	db, err := sql.Open("mysql", config.DatabaseURI)
	if err != nil {
		panic(err)
	}
	query := product.New(db)
	server := http.NewServer(query)
	server.Start()
}
