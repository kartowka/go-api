package main

import (
	"fmt"

	"github.com/antfley/go-api/routers"
	"github.com/antfley/go-api/services"
	"github.com/antfley/go-api/services/db"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Server struct {
	db *gorm.DB
	Router *gin.Engine
}
func (server *Server) InitServer(){
	server.Router = gin.Default()
	var Router = routers.RouteLoader{}
	for _, routes := range Router.LoadRoutes(){
		routes.Route(server.Router)
	}
}
func (server *Server) InjectDB(){
	services.InjectDBIntoServices(server.db)
}
func (server *Server) Run(){
	fmt.Println("Rise and shine! 🌞🌞🌞")
	fmt.Println("Listening on port : 5050")
	server.Router.Run("127.0.0.1:5050")
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
	app.InjectDB()
	app.InitServer()
	app.Run()
}
