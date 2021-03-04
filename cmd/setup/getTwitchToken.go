package main

import (
	"fmt"

	"github.com/DanielStefanK/twitchbot/internal/config"

	"github.com/nicklaw5/helix"
)

func main() {
	cfg := config.LoadConfig()

	client, err := helix.NewClient(&helix.Options{
		ClientID:    cfg.Bot.ClientID,
		RedirectURI: "http://localhost:8080/redirect",
	})
	if err != nil {
		// handle error
	}

	url := client.GetAuthorizationURL(&helix.AuthorizationURLParams{
		ResponseType: "code", // or "token"
		Scopes:       []string{"user:read:email"},
		State:        "some-state",
		ForceVerify:  false,
	})

	fmt.Printf("%s\n", url)
}
