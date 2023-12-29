package model

import (
	"time"

	"gorm.io/gorm"
)

type Player struct {
	Username string `gorm:"primaryKey"`
	PlayerLevel int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Matches []Match `gorm:"foreignKey:PlayerUsername"`
	Room Room `gorm:"foreignKey:OwnerUsername"`
}
