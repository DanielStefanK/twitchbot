package bot

import (
	"fmt"

	"github.com/DanielStefanK/twitchbot/internal/logger/storage"

	"github.com/DanielStefanK/twitchbot/internal/logger"
	"github.com/gempir/go-twitch-irc/v2"
	"gorm.io/gorm"
)

var log = logger.NewLogger("bot")

// Bot used for communicating to twitch chats
type Bot struct {
	client   *twitch.Client
	db       *gorm.DB
	channels []string
	configs  map[string]*storage.ChatConfig
}

// NewBot creates a new chatbot
func NewBot(username, oauth string, channels []string, db *gorm.DB) *Bot {
	client := twitch.NewClient(username, oauth)

	bot := &Bot{client: client, db: db, channels: channels}

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
			})
			return
		}
	}
}

// SendMsg sends the given message to the given channel
func (b *Bot) SendMsg(channel, msg string) {
	log.Info(fmt.Sprintf("sending message to channel %s", channel))
	b.client.Say(channel, msg)
}
