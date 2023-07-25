package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	auth "github.com/vikingpingvin/hackweek-ssp/internal/auth"
	handlers "github.com/vikingpingvin/hackweek-ssp/internal/web/handlers"
)

func main() {
	// init .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	auth.OauthConfig.ClientSecret = os.Getenv("GOOGLE_OAUTH2_CLIENT_SECRET")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./src/views/mainPage.html")
	})

	http.HandleFunc("/content", handlers.HandleContentView)

	http.HandleFunc("/apiList", handlers.HandleApiList)
	http.HandleFunc("/api/", handlers.HandleApiContent)

	// Authentication
	http.HandleFunc("/auth/login", auth.HandleLogin)
	http.HandleFunc("/auth/callback", auth.HandleCallback)

	log.Println("Server starting on port 8080...")
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
