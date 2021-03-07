package main

import (
	"github.com/DanielStefanK/twitchbot/internal/config"
	"github.com/DanielStefanK/twitchbot/pkg/api"
)

func main() {
	cfg := config.LoadConfig()
	api.Serve(cfg)
}
