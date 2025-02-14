package configs

import (
	log "github.com/sirupsen/logrus"
	"os"
)

func CheckEnv(envs []string) {
	for _, env := range envs {
		if os.Getenv(env) == "" {
			log.Fatalf("Env %v is not defined", env)
		}
	}
}
