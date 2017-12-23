package app

import (
	"time"

	"github.com/joho/godotenv"
)

const (
	// envFileName file name with enviroment settings
	envFileName = ".env"
	// envTimeOut how rare will be env refreshment (in seconds)
	envTimeOut = 30
)

// Enviroment content from parsed .env file
type Enviroment map[string]string

// initEnv initialize enviroment
func initEnv() {
	fileEnv, error := godotenv.Read(projectRoot + srcFolder + envFileName)
	if error == nil {
		env = fileEnv
	}
}

// GetEnv returns enviroment value by key
func GetEnv(key string) string {
	return env[key]
}

// liveEnv refreshes enviroment from file every envTimeOut seconds
func liveEnv() {
	for {
		time.Sleep(envTimeOut * time.Second)
		initEnv()
	}
}
