package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/gorilla/mux"
	"github.com/mott93/discord-authentication/internal/config"
)

var (
	conf   = &config.Config{}
	tmpl   = &template.Template{}
	server *Server
)

const (
	DiscordAPIBase        = "https://discord.com/api"
	DiscordOAuth2URL      = DiscordAPIBase + "/oauth2/authorize"
	DiscordTokenURL       = DiscordAPIBase + "/oauth2/token"
	DiscordConnectionsURL = DiscordAPIBase + "/v10/users/@me/connections"
)

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
	config    *config.Config
	templates *template.Template
}

func loadTemplates() *template.Template {
	return template.Must(template.ParseGlob("templates/*.html"))
}

func init() {
	conf.LoadConfig()
	templates = loadTemplates()
	server = &Server{
		config:    conf,
		templates: templates,
	}
}

func (s *Server) RenderTemplate(w http.ResponseWriter, tmpl string, data interface{}) error {
	return s.templates.ExecuteTemplate(w, tmpl, data)
}

func (s *Server) getAccessToken(code string) (*TokenResponse, error) {
	data := url.Values{}
	data.Set("client_id", s.config.ClientID)
	data.Set("client_secret", s.config.ClientSecret)
	data.Set("grant_type", "authorization_code")
	data.Set("code", code)
	data.Set("redirect_uri", s.config.RedirectURI)

	res, err := http.Post(DiscordTokenURL, "application/x-www-form-urlencoded", strings.NewReader(data.Encode()))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var tokenRes TokenResponse
	if err := json.NewDecoder(res.Body).Decode(&tokenRes); err != nil {
		return nil, err
	}

	return &tokenRes, nil
}

func (s *Server) GetConnections(token string) ([]Connection, error) {
	req, err := http.NewRequest("GET", DiscordConnectionsURL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	c := &http.Client{}
	res, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var connections []Connection
	if err := json.NewDecoder(res.Body).Decode(&connections); err != nil {
		return nil, err
	}

	return connections, nil

}

func main() {
	fmt.Println("Hello Discord Authentication!")
	r := mux.NewRouter()

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))))

	log.Printf("鯖起動... http://localhost:%s", conf.Port)
	log.Fatal(http.ListenAndServe(":"+conf.Port, r))
}
