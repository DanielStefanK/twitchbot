package bot

import (
	"fmt"
	"math"

	"github.com/gempir/go-twitch-irc/v2"
	"gorm.io/gorm"
)

type replyFunc func(msg string)

type commandHandler func(db *gorm.DB, msg twitch.PrivateMessage, reply replyFunc)

// Command for the chat
type Command struct {
	Cmd     string
	Help    string
	DBNum   uint
	Handler commandHandler
}

var allEnabled uint = 65407

var devHandler = func(db *gorm.DB, msg twitch.PrivateMessage, reply replyFunc) {
	fmt.Println(msg.Message)
	reply("ok")
}

var allCommands = [15]*Command{
	&Command{Cmd: "so", DBNum: calNumber(0), Help: "Shoutout other streams", Handler: devHandler},
	&Command{Cmd: "gamble", DBNum: calNumber(1), Help: "Gamble with you channel points", Handler: devHandler},
	&Command{Cmd: "weather", DBNum: calNumber(2), Help: "Get the weather in the given city", Handler: devHandler},
	&Command{Cmd: "time", DBNum: calNumber(3), Help: "Get the time in the given country/city", Handler: devHandler},
	&Command{Cmd: "horoscope", DBNum: calNumber(4), Help: "Get the current horoscope", Handler: devHandler},
	&Command{Cmd: "pat", DBNum: calNumber(5), Help: "Pat a user", Handler: devHandler},
	&Command{Cmd: "soulmate", DBNum: calNumber(6), Help: "Find a soulmate in the chat", Handler: devHandler},
	&Command{Cmd: "robbery", DBNum: calNumber(1), Help: "Try to rob the bot", Handler: devHandler},
	&Command{Cmd: "robberychances", DBNum: calNumber(7), Help: "find out you chances to rob the bot", Handler: devHandler},
	&Command{Cmd: "robberygiveup", DBNum: calNumber(8), Help: "give up the fight", Handler: devHandler},
	&Command{Cmd: "robberyfight", DBNum: calNumber(9), Help: "fight", Handler: devHandler},
	&Command{Cmd: "timer", DBNum: calNumber(10), Help: "Set a timer for you in chat", Handler: devHandler},
	&Command{Cmd: "watchtime", DBNum: calNumber(11), Help: "Set a timer for you in chat", Handler: devHandler},
	&Command{Cmd: "push", DBNum: calNumber(12), Help: "push someone off a cliff", Handler: devHandler},
	&Command{Cmd: "hug", DBNum: calNumber(13), Help: "hug someone in chat", Handler: devHandler},
}

func calNumber(i uint) uint {
	return uint(math.Pow(2, float64(i)))
}
