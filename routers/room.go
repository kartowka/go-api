package routers

import (
	"github.com/antfley/go-api/controller"
	"github.com/gin-gonic/gin"
)

type Room struct {}

func (r *Room) Route(route *gin.Engine) {
	router := route.Group("/room")
	Controller := controller.RoomController{}
	router.POST("/", Controller.CreateRoom)
}
