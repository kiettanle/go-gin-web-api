package config

import (
	"os"

	"github.com/subosito/gotenv"
)

func init() {
	gotenv.Load()
}

type Config struct {
	Uri  string
	Port string
}

func GetConfig() Config {
	env := Config{
		Uri:  os.Getenv("URI"),
		Port: os.Getenv("PORT"),
	}

	return env
}
