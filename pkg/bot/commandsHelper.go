package bot

import (
	"fmt"
	"math"
	"math/rand"

	"github.com/gempir/go-twitch-irc/v2"
	"gorm.io/gorm"
)

type replyFunc func(msg string)

type commandHandler func(db *gorm.DB, msg twitch.PrivateMessage, reply replyFunc, client *twitch.Client)

// Command for the chat
type Command struct {
	Cmd     string
	Help    string
	DBNum   uint
	Handler commandHandler
}

var allEnabled uint = 65407

var devHandler = func(db *gorm.DB, msg twitch.PrivateMessage, reply replyFunc, client *twitch.Client) {
	reply(fmt.Sprintf("not implemented yet"))
}

var allCommands = [15]*Command{
	&Command{Cmd: "so", DBNum: calNumber(0), Help: "Shoutout other streams", Handler: func(db *gorm.DB, msg twitch.PrivateMessage, reply replyFunc, client *twitch.Client) {
		params := getParams(msg.Message)

		if len(params) == 0 {
			reply("sorry, but who should i shoutout?")
			return
		}
		//TODO: insert callback for user immediate answer?

		reply(fmt.Sprintf("Shoutout to @%s", params[0]))
	}},
	&Command{Cmd: "time", DBNum: calNumber(1), Help: "Get the time in the given country/city", Handler: func(db *gorm.DB, msg twitch.PrivateMessage, reply replyFunc, client *twitch.Client) {
		//TODO: implement
		reply(fmt.Sprintf("not implemented yet"))
	}},
	&Command{Cmd: "weather", DBNum: calNumber(2), Help: "Get the weather in the given city", Handler: devHandler},
	&Command{Cmd: "horoscope", DBNum: calNumber(3), Help: "Get the current horoscope", Handler: devHandler},
	&Command{Cmd: "gamble", DBNum: calNumber(4), Help: "Gamble with you channel points", Handler: devHandler},
	&Command{Cmd: "pat", DBNum: calNumber(5), Help: "Pat a user", Handler: func(db *gorm.DB, msg twitch.PrivateMessage, reply replyFunc, client *twitch.Client) {
		params := getParams(msg.Message)
		users, err := client.Userlist(msg.Channel)

		if err != nil {
			log.Error("could not load the current user in the chat")
			reply("sorry there seems to be a problem")
			return
		}

		if len(params) == 0 || len(params) > 1 {
			reply("you can pet one other chatter at a time")
			return
		}

		if !contains(users, params[0]) {
			//FIXME: does this work=
			reply(fmt.Sprintf("the user %s does not appear to be in the chat right now.", params[0]))
			return
		}

		reply(fmt.Sprintf("You are doing a great job @%s!!", params[0]))
	}},
	&Command{Cmd: "soulmate", DBNum: calNumber(6), Help: "Find a soulmate in the chat", Handler: func(db *gorm.DB, msg twitch.PrivateMessage, reply replyFunc, client *twitch.Client) {
		users, err := client.Userlist(msg.Channel)

		if err != nil {
			log.Error("could not load the current user in the chat")
			reply("sorry there seems to be a problem")
			return
		}

		if len(users)-1 <= 0 {
			reply("sorry, but there are not enough chatter in the chat")
			return
		}

		reply(fmt.Sprintf("your soul mate is... @%s!!", users[rand.Intn(len(users)-1)]))
	}},
	&Command{Cmd: "robbery", DBNum: calNumber(1), Help: "Try to rob the bot", Handler: devHandler},
	&Command{Cmd: "robberychances", DBNum: calNumber(7), Help: "find out you chances to rob the bot", Handler: devHandler},
	&Command{Cmd: "robberygiveup", DBNum: calNumber(8), Help: "give up the fight", Handler: devHandler},
	&Command{Cmd: "robberyfight", DBNum: calNumber(9), Help: "fight", Handler: devHandler},
	&Command{Cmd: "timer", DBNum: calNumber(10), Help: "Set a timer for you in chat", Handler: devHandler},
	&Command{Cmd: "watchtime", DBNum: calNumber(11), Help: "Set a timer for you in chat", Handler: devHandler},
	&Command{Cmd: "push", DBNum: calNumber(12), Help: "push someone off a cliff", Handler: devHandler},
	&Command{Cmd: "hug", DBNum: calNumber(13), Help: "hug someone in chat", Handler: func(db *gorm.DB, msg twitch.PrivateMessage, reply replyFunc, client *twitch.Client) {
		params := getParams(msg.Message)
		users, err := client.Userlist(msg.Channel)

		if err != nil {
			log.Error("could not load the current user in the chat")
			reply("sorry there seems to be a problem")
			return
		}

		if len(params) == 0 || len(params) > 1 {
			reply("you can hug one other chatter at a time")
			return
		}

		if !contains(users, params[0]) {
			//FIXME: does this work
			reply(fmt.Sprintf("the user %s does not appear to be in the chat right now. Maybe hug him later", params[0]))
			return
		}

		reply(fmt.Sprintf("%s is give %s a hug!", msg.User.DisplayName, params[0]))
	}},
}

func calNumber(i uint) uint {
	return uint(math.Pow(2, float64(i)))
}
