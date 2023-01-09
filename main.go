package main

import (
	"github.com/Shitikyan/Go-book-crud/config"
	"github.com/Shitikyan/Go-book-crud/models"
	"github.com/Shitikyan/Go-book-crud/router"
)

func main() {

	r := router.Routes()
	config.DatabaseConnection()
	models.Migrator()

	r.Run()
}
