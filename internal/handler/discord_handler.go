package handler

import (
	"html/template"

	"github.com/mott93/discord-authentication/internal/config"
)

type DiscordHandler struct {
	config    *config.Config
	templates *template.Template
}

func NewDiscordHandler(config *config.Config, tmpl *template.Template) *DiscordHandler {
	return &DiscordHandler{
		config:    config,
		templates: tmpl,
	}
}
