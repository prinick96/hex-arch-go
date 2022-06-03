package env

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/prinick96/elog"
)

// Default PORT to set if .env PORT var is empty or can't load
const DEFAULT_PORT_IF_EMPTY = "80"

// Env config struct
type EnvApp struct {
	// Server Envs
	SERVER_PORT string

	// Database Envs
	DB_ENGINE   string
	DB_HOST     string
	DB_PORT     string
	DB_DATABASE string
	DB_USERNAME string
	DB_PASSWORD string
	DB_OPTIONS  string
}

// Get the env configuration
func GetEnv(env_file string) EnvApp {
	err := godotenv.Load(env_file)
	elog.New(elog.PANIC, "Error loading "+env_file+" file", err)

	// Heroku smell
	port := os.Getenv("PORT")
	if port == "" {
		port = DEFAULT_PORT_IF_EMPTY
	}

	return EnvApp{
		SERVER_PORT: port,
		DB_ENGINE:   os.Getenv("DB_ENGINE"),
		DB_HOST:     os.Getenv("DB_HOST"),
		DB_PORT:     os.Getenv("DB_PORT"),
		DB_DATABASE: os.Getenv("DB_DATABASE"),
		DB_USERNAME: os.Getenv("DB_USERNAME"),
		DB_PASSWORD: os.Getenv("DB_PASSWORD"),
		DB_OPTIONS:  os.Getenv("DB_OPTIONS"),
	}
}
