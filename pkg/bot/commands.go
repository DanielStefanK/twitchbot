package bot

import (
	"fmt"
	"strings"
	"time"

	"github.com/nicklaw5/helix"

	owm "github.com/briandowns/openweathermap"
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

func timeCmd(db *gorm.DB, msg twitch.PrivateMessage, reply replyFunc, bot *Bot) {
	params := getParams(msg.Message)

	if len(params) == 0 {
		reply("Please specify a city")
		return
	}

	w, err := owm.NewCurrent("C", "en", bot.weatherAPI) // fahrenheit (imperial) with Russian output

	if err != nil {
		log.Error("could not create weather api")
		panic(err)
	}

	city := strings.Join(params, " ")

	errw := w.CurrentByName(city)

	if errw != nil {
		reply(fmt.Sprintf("could not find city %s", city))
		return
	}

	if w.ID == 0 {
		reply(fmt.Sprintf("could not find city %s", city))
		return
	}

	destinationUnixSeconds := time.Now().In(time.UTC).Unix() + (int64(w.Timezone))

	log.Info(fmt.Sprintf("%d %s", destinationUnixSeconds, time.Unix(destinationUnixSeconds, 0).Location().String()))

	currentDate := time.Unix(destinationUnixSeconds, 0)

	//TODO: something does not work here
	reply(fmt.Sprintf("current time in %s: %s", w.Name, currentDate.In(time.UTC).Format(time.RFC1123)))
}

func weather(db *gorm.DB, msg twitch.PrivateMessage, reply replyFunc, bot *Bot) {
	params := getParams(msg.Message)

	if len(params) == 0 {
		reply("Please specify a city")
		return
	}

	w, err := owm.NewCurrent("C", "en", bot.weatherAPI) // fahrenheit (imperial) with Russian output

	if err != nil {
		log.Error("could not create weather api")
		panic(err)
	}

	errw := w.CurrentByName(params[0])

	if errw != nil {
		reply(fmt.Sprintf("could not find city %s", params[0]))
		return
	}

	if w.ID == 0 {
		reply(fmt.Sprintf("could not find city %s", params[0]))
		return
	}

	weathers := make([]string, len(w.Weather), len(w.Weather))

	for idx, v := range w.Weather {
		weathers[idx] = v.Description
	}

	strings.Join(weathers, ", ")

	reply(fmt.Sprintf("current weather in %s: %s", w.Name, strings.Join(weathers, ", ")))
}
