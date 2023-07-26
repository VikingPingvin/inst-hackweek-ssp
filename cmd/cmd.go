package cmd

import (
	"log"
	"net/http"

	auth "github.com/vikingpingvin/hackweek-ssp/internal/auth"
	handlers "github.com/vikingpingvin/hackweek-ssp/internal/web/handlers"
)

func RunServer() {

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
	// log.Fatal(http.ListenAndServeTLS("myapp.local:8080", certs.GetCerts().CertFile, certs.GetCerts().KeyFile, nil))
}
