package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

var (
	config    *Config
	templates *template.Template
	server    *Server
)

const (
	DiscordAPIBase        = "https://discord.com/api"
	DiscordOAuth2URL      = DiscordAPIBase + "/oauth2/authorize"
	DiscordTokenURL       = DiscordAPIBase + "/oauth2/token"
	DiscordConnectionsURL = DiscordAPIBase + "/v10/users/@me/connections"
)

type Config struct {
	ClientID     string
	ClientSecret string
	RedirectURI  string
	Port         string
}

type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    string `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
}

type Connection struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Type     string `json:"type"`
	Verified bool   `json:"verified"`
}

type Server struct {
	config    *Config
	templates *template.Template
}

func loadConfig() *Config {
	return &Config{
		ClientID:     getEnv("DISCORD_CLIENT_ID", ""),
		ClientSecret: getEnv("DISCORD_CLIENT_SECRET", ""),
		RedirectURI:  getEnv("DISCORD_REDIRECT_URI", ""),
		Port:         getEnv("PORT", "8080"),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}

	return defaultValue
}

func loadTemplates() *template.Template {
	return template.Must(template.ParseGlob("templates/*.html"))
}

func init() {
	config = loadConfig()
	templates = loadTemplates()
	server = &Server{
		config:    config,
		templates: templates,
	}
}

func main() {
	fmt.Println("Hello Discord Authentication!")
	r := mux.NewRouter()

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))))

	log.Printf("鯖起動... http://localhost:%s", config.Port)
	log.Fatal(http.ListenAndServe(":"+config.Port, r))
}
