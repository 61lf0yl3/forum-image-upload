package server

import "forum-image-upload/internal/database"

// Config ...
type Config struct {
	WebPort  string
	Database *database.Config
}

// NewConfig generates configurations for the Server
func NewConfig() *Config {
	return &Config{
		WebPort:  ":8081",
		Database: database.NewConfig(),
	}
}
