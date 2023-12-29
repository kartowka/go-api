package controller

import (
	"fmt"
	"net/http"

	"github.com/antfley/go-api/services"
	"github.com/gin-gonic/gin"
)

type MatchController struct {}

type createMatchInput struct {
	Username 	string 	`json:"username" binding:"required"`
	RoomID		string		`json:"room_id" binding:"required"`
}

func (controller MatchController) JoinMatch (c* gin.Context){
	var input createMatchInput
		if err := c.ShouldBindJSON(&input); err != nil {
		var response = ErrorResponse{
			Msg: "Validation Error",
			Err: err,
		}
		c.AbortWithStatusJSON(http.StatusOK, response)
		return
	}
	isPlayerJoined, err := services.MatchService.FindUserInRoom(input.Username, input.RoomID)
	if err != nil {
		var response = ErrorResponse{
			Msg: "Failed to Join Match",
			Err: err,
		}
		c.AbortWithStatusJSON(http.StatusOK, response)
		return
	}
	fmt.Println(isPlayerJoined)
	match,err := services.MatchService.PlayerJoinRoom( input.Username,input.RoomID)
	if err != nil {
		var response = ErrorResponse{
			Msg: "Failed to Join Match",
			Err: err,
		}
		c.AbortWithStatusJSON(http.StatusOK, response)
		return
	}
	var responseString = fmt.Sprintf("Successfully joined room : %s", match.RoomID)
	var response = SuccessResponse{
		Status: true,
		Msg:   responseString,
	}
	c.AbortWithStatusJSON(http.StatusOK, response)
}
