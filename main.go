package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	cmd "github.com/vikingpingvin/hackweek-ssp/cmd"
	"github.com/vikingpingvin/hackweek-ssp/internal/auth"
)

func main() {
	// init .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file: %s", err.Error())
	}
	auth.OauthConfig.ClientSecret = os.Getenv("GOOGLE_OAUTH2_CLIENT_SECRET")

	cmd.RunServer()
}
