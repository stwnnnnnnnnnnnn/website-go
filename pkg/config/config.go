package config

import "html/template"

// app config hold the app config
type AppConfig struct {
	Usecache      bool
	TemplateCache map[string]*template.Template
}
