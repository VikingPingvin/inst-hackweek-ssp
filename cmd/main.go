package main

import (
	"log"
	"net/http"

	handlers "github.com/vikingpingvin/hackweek-ssp/internal/web/handlers"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./src/views/mainPage.html")
	})

	http.HandleFunc("/content", handlers.HandleContentView)

	http.HandleFunc("/apiList", handlers.HandleApiList)
	http.HandleFunc("/api/", handlers.HandleApiContent)

	log.Println("Server starting on port 8080...")
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
