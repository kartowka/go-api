package controller

import (
	"fmt"
	"net/http"

	"github.com/antfley/go-api/services"
	"github.com/gin-gonic/gin"
)

type RoomController struct {}

type createRoomInput struct {
	OwnerUsername string `json:"owner_username"`
	Title string `json:"title"`
	Description string `json:"description"`
}

func (controller RoomController) CreateRoom (c *gin.Context) {
	var input createRoomInput
	if err := c.ShouldBindJSON(&input); err != nil {
		var response = ErrorResponse{
			Msg: "Validation Error",
			Err: err,
		}
		c.AbortWithStatusJSON(http.StatusOK, response)
		return
	}
	room,err := services.RoomService.CreateOne(input.OwnerUsername, input.Title, input.Description)
	if err != nil {
		var response = ErrorResponse{
			Msg: "Failed to Create Room",
			Err: err,
		}
		c.AbortWithStatusJSON(http.StatusOK, response)
		return
	}
	var responseString = fmt.Sprintf("Successfully created a new room with owner : %s", room.OwnerUsername)
	var response = SuccessResponse{
		Status: true,
		Msg:   responseString,
	}
	c.AbortWithStatusJSON(http.StatusOK, response)
}
