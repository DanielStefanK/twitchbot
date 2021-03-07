package api

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/DanielStefanK/twitchbot/internal/config"
	"github.com/DanielStefanK/twitchbot/pkg/weather"

	"github.com/DanielStefanK/twitchbot/pkg/horoscope"
	"github.com/gorilla/mux"
)

// Serve the api endpoints
func Serve(config *config.Config) {
	r := mux.NewRouter()
	horoscope.StartIntervalScraping()
	weatherAPI := weather.NewWeatherAPIClient(config.Bot.OpenWeatherAPI)

	r.HandleFunc("/horoscope", func(w http.ResponseWriter, r *http.Request) {
		vars := r.URL.Query()
		h := horoscope.GetHoroscope(vars["type"][0])
		if h == "" {
			fmt.Fprint(w, "could not load horoscope")
		}
		fmt.Fprint(w, h)
	})

	r.HandleFunc("/weather", func(w http.ResponseWriter, r *http.Request) {
		vars := r.URL.Query()

		if len(vars["city"]) == 0 {
			fmt.Fprint(w, "Specify a city")
			return
		}

		city := strings.Join(vars["city"], " ")
		city = strings.ReplaceAll(city, "+", " ")

		text := weatherAPI.GetWeatherText(city)

		fmt.Fprint(w, text)
	})

	r.HandleFunc("/time", func(w http.ResponseWriter, r *http.Request) {
		vars := r.URL.Query()

		if len(vars["city"]) == 0 {
			fmt.Fprint(w, "Specify a city")
			return
		}

		city := strings.Join(vars["city"], " ")
		city = strings.ReplaceAll(city, "+", " ")

		text := weatherAPI.GetTimeText(city)

		fmt.Fprint(w, text)
	})

	http.ListenAndServe(":3000", r)
}
