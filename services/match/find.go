package match

import (
	"github.com/antfley/go-api/model"
)
func (s MatchService) FindAll () ([]model.Match, error) {
	var matches []model.Match
	result := s.DB.Find(&matches)
		return matches, result.Error
}
func (s MatchService) FindUserInRoom (username string,roomID string) (model.Match,error){
	var match model.Match
	result := s.DB.Where(&model.Match{PlayerUsername: username,RoomID: roomID})
	return match, result.Error
}
