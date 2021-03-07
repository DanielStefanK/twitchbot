package horoscope

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/DanielStefanK/twitchbot/internal/logger"
	"github.com/PuerkitoBio/goquery"
)

var log = logger.NewLogger("horoscope-scraper")

var url = "https://www.horoscope.com/us/horoscopes/general/horoscope-general-daily-today.aspx?sign=%d"

var signToInt = map[string]int{
	"Aries":       1,
	"Taurus":      2,
	"Gemini":      3,
	"Cancer":      4,
	"Leo":         5,
	"Virgo":       6,
	"Libra":       7,
	"Scorpius":    8,
	"Sagittarius": 9,
	"Capricornus": 10,
	"Aquarius":    11,
	"Pisces":      12,
}

var signToHoroscope = map[int]string{
	1:  "Aries",
	2:  "Taurus",
	3:  "Gemini",
	4:  "Cancer",
	5:  "Leo",
	6:  "Virgo",
	7:  "Libra",
	8:  "Scorpius",
	9:  "Sagittarius",
	10: "Capricornus",
	11: "Aquarius",
	12: "Pisces",
}

var signToAlt = map[int][]string{
	1:  []string{"Aries", "Ram"},
	2:  []string{"Taurus", "Bull"},
	3:  []string{"Gemini", "Twins"},
	4:  []string{"Cancer", "Crab"},
	5:  []string{"Leo", "Lion"},
	6:  []string{"Virgo", "Virgin"},
	7:  []string{"Libra", "Balance"},
	8:  []string{"Scorpius", "Scorpion"},
	9:  []string{"Sagittarius", "Archer"},
	10: []string{"Capricornus", "Goat"},
	11: []string{"Aquarius", "Water Bearer"},
	12: []string{"Pisces", "Fish"},
}

// GetHoroscope get the horoscope that is currently stored
func GetHoroscope(name string) string {
	signNo := getInt(name)
	if signNo < 0 {
		return ""
	}
	return signToHoroscope[signNo]
}

func getInt(name string) int {
	for i := 1; i <= 12; i++ {
		for _, alt := range (signToAlt)[i] {
			if strings.ToLower(alt) == strings.ToLower(name) {
				return i
			}
		}
	}
	return -1
}

// StartIntervalScraping starts the scraper and runs every hour
func StartIntervalScraping() {
	go func() {
		for {
			scrape()
			time.Sleep(time.Hour)
		}
	}()
}
func scrape() {
	for i := 1; i <= 12; i++ {
		// Get the HTML
		resp, err := http.Get(fmt.Sprintf(url, i))

		if err != nil {
			log.Error("could not scrape horoscopes")
			log.Error(err.Error())
			return
		}

		// Convert HTML into goquery document
		doc, err := goquery.NewDocumentFromReader(resp.Body)

		if err != nil {
			log.Error("could not parse horoscopes page")
			log.Error(err.Error())
			return
		}

		main := doc.Find(".main-horoscope").First()
		horoscope := main.Find("p").First()

		signToHoroscope[i] = horoscope.Text()
	}
}
