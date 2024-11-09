package internal

import (
	"flag"
	"os"
)

func env(env, def string) string {
	value := os.Getenv(env)
	if value == "" {
		return def
	}

	return value

}

func NewConfig() Config {
	config := Config{
		ApplicationCommand: env("APPLICATION_COMMAND", ""),
		DatabaseConnection: env("DATABASE_CONNECTION", "file:snippet.sqlite"),
		ServerPort:         env("PORT", "8080"),
		ServerHost:         env("HOST", "127.0.0.1"),
	}

	flag.StringVar(&config.ApplicationCommand, "cmd", config.ApplicationCommand, "command to run")
	flag.StringVar(&config.DatabaseConnection, "db", config.DatabaseConnection, "database connection string")
	flag.StringVar(&config.ServerPort, "port", config.ServerPort, "port to listen on")
	flag.StringVar(&config.ServerHost, "host", config.ServerHost, "host to listen on")
	flag.Parse()

	return config
}

type Config struct {
	ApplicationCommand string

	DatabaseConnection string

	ServerPort string
	ServerHost string
}
