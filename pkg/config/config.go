package config

import (
	"html/template"
	"log"

	"github.com/alexedwards/scs/v2"
)

type AppConfig struct {
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger
	UseCache      bool
	InProduction  bool
	Session 	 *scs.SessionManager
}