package routers

import (
	"github.com/antfley/go-api/controller"
	"github.com/gin-gonic/gin"
)

type Match struct {}

func (m *Match) Route(route *gin.Engine) {
	router := route.Group("/match")
	Controller := controller.MatchController{}
	router.POST("/join", Controller.JoinMatch)
}
