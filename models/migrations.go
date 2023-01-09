package models

import "github.com/Shitikyan/Go-book-crud/config"

func Migrator() {
	db := config.GetDatabase()

	db.AutoMigrate(&Author{})
	db.AutoMigrate(&Book{})
}
