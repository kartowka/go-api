package main

import (
	"github.com/antfley/go-api/services/db"
	"gorm.io/gorm"
)

type Server struct {
	db *gorm.DB
}

func main(){
	DBConfig := db.DBConfig{
		DBHost: "localhost",
		DBUser: "root",
		DBName: "game",
		DBPort: "5432",
		DbPass: "root",
	}
	app := Server{db: db.LoadDB(DBConfig)}
	db.DBMigrate(app.db)
}
