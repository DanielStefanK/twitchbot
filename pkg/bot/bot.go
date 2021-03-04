package bot

import (
	"fmt"

	"github.com/DanielStefanK/twitchbot/internal/config"

	"github.com/DanielStefanK/twitchbot/internal/storage"
	"github.com/nicklaw5/helix"

	"github.com/DanielStefanK/twitchbot/internal/logger"
	"github.com/gempir/go-twitch-irc/v2"
	"gorm.io/gorm"
)

var log = logger.NewLogger("bot")

// Bot used for communicating to twitch chats
type Bot struct {
	client    *twitch.Client
	db        *gorm.DB
	channels  []string
	configs   map[string]*storage.ChatConfig
	twitchAPI *helix.Client
	// mapsAPI   *maps.Client
}

// NewBot creates a new chatbot
func NewBot(cfg *config.Config, channels []string, db *gorm.DB) *Bot {

	// Twitch api
	twitchAPI, err := helix.NewClient(&helix.Options{
		ClientID:     cfg.Bot.ClientID,
		ClientSecret: cfg.Bot.ClientSecret,
	})

	if err != nil {
		log.Error("could not creat twitch api client")
		panic(err)
	}

	apptoken, err := twitchAPI.RequestAppAccessToken([]string{})

	if err != nil {
		log.Error("could not get app access token")
		panic(err)
	}

	twitchAPI.SetAppAccessToken(apptoken.Data.AccessToken)

	// twitch bot

	client := twitch.NewClient(cfg.Bot.Username, cfg.Bot.OAuth)

	// maps api
	// c, err := maps.NewClient(maps.WithAPIKey(cfg.Bot.MapsAPIToken))

	// if err != nil {
	// 	log.Error("could not create maps api")
	// 	panic(err)
	// }

	bot := &Bot{client: client, db: db, channels: channels, twitchAPI: twitchAPI}

	bot.reloadConfigs()

	client.OnPrivateMessage(bot.incomingMsg)

	client.OnConnect(func() {
		log.Info("client successfully connected to twitch chat")
	})

	client.Join(channels...)

	go func() { client.Connect() }()

	return bot
}

func (b *Bot) incomingMsg(msg twitch.PrivateMessage) {
	if !isCmd(msg.Message) {
		return
	}
	cmd := getCommand(msg.Message)
	for _, c := range allCommands {
		if cmd == c.Cmd {
			c.Handler(b.db, msg, func(replyMsg string) {
				b.SendMsg(msg.Channel, replyMsg)
			}, b)
			return
		}
	}
}

// SendMsg sends the given message to the given channel
func (b *Bot) SendMsg(channel, msg string) {
	log.Info(fmt.Sprintf("sending message to channel %s", channel))
	b.client.Say(channel, msg)
}

// Join joins a channel
func (b *Bot) Join(channels ...string) {
	for _, c := range channels {
		log.Info(fmt.Sprintf("joining channel %s", c))
		b.client.Join()
		b.RefreshConf(c)
	}
}
