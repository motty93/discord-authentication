package main

import "fmt"

const (
	DiscordAPIBase        = "https://discord.com/api"
	DiscordOAuth2URL      = DiscordAPIBase + "/oauth2/authorize"
	DiscordTokenURL       = DiscordAPIBase + "/oauth2/token"
	DiscordConnectionsURL = DiscordAPIBase + "/v10/users/@me/connections"
)

func main() {
	fmt.Println("Hello Discord Authentication!")
}
