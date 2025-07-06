package config

import "os"

type Config struct {
	ClientID     string
	ClientSecret string
	RedirectURI  string
	Port         string
}

func (c *Config) LoadConfig() {
	c.ClientID = getEnv("DISCORD_CLIENT_ID", "")
	c.ClientSecret = getEnv("DISCORD_CLIENT_SECRET", "")
	c.RedirectURI = getEnv("DISCORD_REDIRECT_URI", "")
	c.Port = getEnv("PORT", "8080")
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}

	return defaultValue
}
