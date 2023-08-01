package config

import (
	"log"
	"text/template"
)

// AppConfig holds variables that we want to share throughout the application
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger
}
