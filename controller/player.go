package controller

import (
	"fmt"
	"net/http"

	"github.com/antfley/go-api/services/player"
	"github.com/gin-gonic/gin"
)

type PlayerController struct {}
type createPlayerInput struct {
	Username string `json:"username" binding:"required"`
}
func (controller PlayerController) CreatePlayer (c *gin.Context){
	var input createPlayerInput
	if err := c.ShouldBindJSON(&input); err != nil {
		var response = ErrorResponse{
			Msg: "Validation error",
			Err: err,
		}
		c.AbortWithStatusJSON(http.StatusOK,response)
		return
	}
	player,err := player.PlayerService.CreateOne(input.Username)
	if err != nil {
		var response = ErrorResponse{
			Msg: "Error creating player",
			Err: err,
		}
		c.AbortWithStatusJSON(http.StatusOK,response)
		return
	}
	var responseString = fmt.Sprintf("Player %s created",player.Username)
	var response = SuccessResponse{
		Status: true,
		Msg: responseString,
	}
	c.AbortWithStatusJSON(http.StatusOK,response)
}
