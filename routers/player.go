package routers

import (
	"github.com/antfley/go-api/controller"
	"github.com/gin-gonic/gin"
)

type Player struct {}
func (p *Player) Route(route *gin.Engine) {
	router := route.Group("/player")
	Controller := controller.PlayerController{}
	router.POST("/", Controller.CreatePlayer)
}
