package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"
)

func HandleContentView(w http.ResponseWriter, r *http.Request) {
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
		return
	}
}

func HandleApiContent(w http.ResponseWriter, r *http.Request) {
	// apiId := r.URL.Query().Get("apiId")
	// env := r.URL.Query().Get("env")
	// region := r.URL.Query().Get("region")
}
