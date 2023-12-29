package services

import (
	"github.com/antfley/go-api/services/match"
	"github.com/antfley/go-api/services/player"
	"github.com/antfley/go-api/services/room"
	"gorm.io/gorm"
)

var (
	PlayerService 	player.PlayerService
	RoomService		room.RoomService
	MatchService	match.MatchService
)

func InjectDBIntoServices (db *gorm.DB) {
	PlayerService.DB 	= db
	RoomService.DB		= db
	MatchService.DB		= db
}
