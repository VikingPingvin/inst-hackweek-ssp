package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"

	"github.com/vikingpingvin/hackweek-ssp/internal/auth"
	"golang.org/x/oauth2"
)

func HandleContentView(w http.ResponseWriter, r *http.Request) {
	authCookie, err := auth.GetAuthCookie(r)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		http.Redirect(w, r, "/", http.StatusUnauthorized)
		return
	}
	err = auth.VerifyUserTokenOnline(&oauth2.Token{
		AccessToken: authCookie.Value,
	})
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		http.Error(w, "Failed to verify user token", http.StatusInternalServerError)
		return
	}
	fmt.Printf("AccessToken Online verification successful\n")

	topicStr := r.URL.Query().Get("topic")
	if topicStr == "" {
		http.Error(w, "Missing topic parameter", http.StatusBadRequest)
		return
	}

	apiId, err := strconv.Atoi(topicStr)
	if err != nil {
		http.Error(w, "Invalid topic parameter", http.StatusBadRequest)
		return
	}

	data := GetApiById(apiId)

	tmpl, err := template.ParseFiles("./src/views/contentView.html",
		"./src/views/api-view/apiSettingComponent.html",
		"./src/views/api-view/settingValueTextComponent.html",
		"./src/views/api-view/openApiSpecComponent.html",
		"./src/views/api-view/sliderButton.html")

	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		http.Error(w, "Failed to load template", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Failed to render template", http.StatusInternalServerError)
		fmt.Printf("Error: %s\n", err.Error())
		return
	}
}
