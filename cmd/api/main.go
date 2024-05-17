package main

import (
	"database/sql"
	_ "github.com/devkemc/challenge-eulabs/docs"
	"github.com/devkemc/challenge-eulabs/internal/server/http"
	config2 "github.com/devkemc/challenge-eulabs/pkg/config"
	"github.com/devkemc/challenge-eulabs/pkg/db/product"
	_ "github.com/go-sql-driver/mysql"
)

// @title			Challenge Eulabs API
// @version		1.0
// @description	API to validate knowledge with golang.
// @contact.name	Joaquim Kennedy Batista de Souza
// @contact.email	joaquimkbs@gmail.com
// @license.name	Apache 2.0
// @BasePath		/api/v1
// @schemes		http
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
