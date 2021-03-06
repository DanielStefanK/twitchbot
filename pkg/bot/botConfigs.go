package bot

import (
	"errors"
	"fmt"

	"github.com/DanielStefanK/twitchbot/internal/storage"
	"gorm.io/gorm"
)

func (b *Bot) reloadConfigs() {
	log.Info("refreshing chat configs")

	configs := make(map[string]*storage.ChatConfig)

	for _, c := range b.channels {
		b.RefreshConf(c)
	}

	b.configs = configs
}

func (b *Bot) createDefaultConfig(c string) {
	config := &storage.ChatConfig{}
	config.Channel = c
	config.Internal = allEnabled
	b.db.Save(config)
	b.configs[c] = config
}

// RefreshConf refreshes the config for the given channel
func (b *Bot) RefreshConf(c string) {
	log.Info(fmt.Sprintf("refreshing chat config for %s", c))

	config := &storage.ChatConfig{}
	err := b.db.First(&config, "channel = ?", c).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		b.configs[c] = config
	} else {
		log.Info(fmt.Sprintf("cold not find a config for channel %s", c))
		log.Info(fmt.Sprintf("creating a default one"))
	}
}
