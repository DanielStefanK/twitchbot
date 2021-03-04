package bot

import (
	"fmt"
	"strings"

	"github.com/nicklaw5/helix"

	"github.com/gempir/go-twitch-irc/v2"
	"gorm.io/gorm"
)

func shoutOut(db *gorm.DB, msg twitch.PrivateMessage, reply replyFunc, bot *Bot) {
	params := getParams(msg.Message)

	if len(params) == 0 {
		reply("sorry, but who should i shoutout?")
		return
	}

	if !strings.HasPrefix(params[0], "@") {
		reply("Please tag the user you want to shoutout. Use !so @fflaminghotcheetos")
		return
	}

	soUserName := params[0][1:]

	res, err := bot.twitchAPI.GetStreams(&helix.StreamsParams{UserLogins: []string{soUserName}})

	if err != nil || len(res.Data.Streams) < 1 {
		reply(fmt.Sprintf("Shoutout to @%s. Currently not live, but check it out later.", soUserName))
		return
	}

	soUser := res.Data.Streams[0]

	log.Info(soUser.Title)

	reply(fmt.Sprintf("Shoutout to @%s. Currently live and streaming \"%s\".", soUserName, soUser.Title))
}

func time(db *gorm.DB, msg twitch.PrivateMessage, reply replyFunc, bot *Bot) {
	// params := getParams(msg.Message)

	reply(fmt.Sprintf("Shoutout to @%s. Currently live and streaming \"%s\".", "soUserName", "soUser.Title"))
}
