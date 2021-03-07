package weather

import (
	"fmt"
	"strings"
	"time"

	"github.com/DanielStefanK/twitchbot/internal/logger"
	owm "github.com/briandowns/openweathermap"
)

var log = logger.NewLogger("weather api")

var maxAge = time.Minute * 30

type cacheEntry struct {
	value    owm.CurrentWeatherData
	expiries time.Time
}

// Client api for making request for weather and timezone
type Client struct {
	apiKey string
	cache  map[string]*cacheEntry
}

// NewWeatherAPIClient creates a new weather api
func NewWeatherAPIClient(key string) *Client {
	return &Client{apiKey: key, cache: make(map[string]*cacheEntry)}
}

// GetWeatherText get the weather text for the given city
func (c *Client) GetWeatherText(city string) string {
	weather := c.getWeather(city)

	if weather == nil {
		return fmt.Sprintf("could not find city %s", city)
	}

	weathers := make([]string, len(weather.Weather), len(weather.Weather))

	for idx, v := range weather.Weather {
		weathers[idx] = v.Description
	}

	strings.Join(weathers, ", ")

	return fmt.Sprintf("Current weather in %s: Temperature: %.2f°C %s", weather.Name, weather.Main.Temp, strings.Join(weathers, ", "))
}

// GetTimeText gets the current time in the given city
func (c *Client) GetTimeText(city string) string {
	w := c.getWeather(city)

	if w == nil {
		return fmt.Sprintf("could not find city %s", city)
	}

	destinationUnixSeconds := time.Now().In(time.UTC).Unix() + (int64(w.Timezone))

	currentDate := time.Unix(destinationUnixSeconds, 0)

	return fmt.Sprintf("current time in %s: %s", w.Name, currentDate.In(time.UTC).Format(time.RFC1123))
}

func (c *Client) getFromCache(city string) *owm.CurrentWeatherData {
	data := c.cache[strings.ToLower(city)]
	if data == nil {
		log.Info(fmt.Sprintf("cache miss for %s", city))
		return nil
	}

	if time.Now().After(data.expiries) {
		log.Info(fmt.Sprintf("cache expired for %s", city))
		c.cache[strings.ToLower(city)] = nil
		return nil
	}

	log.Info(fmt.Sprintf("cache hit for %s", city))

	return &data.value
}

func (c *Client) writeToCache(city string, value *owm.CurrentWeatherData) {
	log.Info(fmt.Sprintf("write %s to cache", city))
	c.cache[strings.ToLower(city)] = &cacheEntry{value: *value, expiries: time.Now().Add(maxAge)}
}

func (c *Client) getWeather(city string) *owm.CurrentWeatherData {
	weather := c.getFromCache(city)
	if weather == nil {
		log.Info(fmt.Sprintf("target %s not found in cache trying to fetch it", city))
		var err error
		weather, err = owm.NewCurrent("C", "en", c.apiKey) // fahrenheit (imperial) with Russian output

		if err != nil {
			return nil
		}

		errw := weather.CurrentByName(city)

		if errw != nil {
			return nil
		}

		if weather.ID == 0 {
			return nil
		}
		c.writeToCache(city, weather)
	} else {
		log.Info(fmt.Sprintf("target %s  found in cache", city))
	}
	return weather
}
