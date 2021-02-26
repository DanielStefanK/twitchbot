package storage

import (
	"gorm.io/gorm"
)

// ChatConfig a chat configuration
type ChatConfig struct {
	gorm.Model
	Channel        string
	Internal       uint
	CustomCommands []*CustomCommand `gorm:"foreignKey:ID"`
}

// CustomCommand a custom command with the custom response
type CustomCommand struct {
	ID       uint `gorm:"primarykey"`
	Cmd      string
	Response string
}
