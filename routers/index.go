package routers

import (
	"github.com/gin-gonic/gin"
)

type RouterInterface interface{
	Route(*gin.Engine)
}
type RouteLoader struct {}

func (r *RouteLoader) LoadRoutes() []RouterInterface {
	player := new (Player)
	match := new (Match)
	room := new (Room)
	return []RouterInterface{player, match, room}
}
