package configs

import (
	"os"
)

type Environment struct {
	Port string
}

func Env() Environment {
	port := os.Getenv("PORT")

	return Environment{
		Port:        port,
	}
}
