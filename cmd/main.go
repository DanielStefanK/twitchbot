package main

import (
	"github.com/DanielStefanK/twitchbot/internal/config"
	"github.com/DanielStefanK/twitchbot/internal/logger"
	"github.com/DanielStefanK/twitchbot/internal/storage"

	"github.com/DanielStefanK/twitchbot/pkg/bot"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var log = logger.NewLogger("main")

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	cfg := config.LoadConfig()

	if err != nil {
		log.Error("failed to connect to db")
		panic("shutting down")
	}

	log.Info("Migrating DB")
	storage.Migrate(db)

	log.Info("creating bot")
	bot.NewBot(cfg, []string{"im_qt"}, db)

	for {

	}
}
