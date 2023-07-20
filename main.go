package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"strconv"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./src/views/mainPage.html")
	})

	http.HandleFunc("/content", handleContentView)

	http.HandleFunc("/apiList", handleApiList)

	log.Println("Server starting on port 8080...")
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

type ApiData struct {
	ApiID   int
	ApiName string
	Message string
}

func handleContentView(w http.ResponseWriter, r *http.Request) {
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

	data := getApiById(apiId)

	tmpl, err := template.ParseFiles("./src/views/content.html")
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

func getAPIs() []ApiData {
	apisData := []ApiData{
		{
			ApiID:   1,
			ApiName: "My Api 1",
			Message: "This is the content for API 1",
		},
		{
			ApiID:   2,
			ApiName: "Super API",
			Message: "This is the content for API 2",
		},
		{
			ApiID:   3,
			ApiName: "Useful API",
			Message: "This is the content for API 3",
		},
	}
	return apisData
}

func getApiById(id int) ApiData {
	apisData := getAPIs()
	for _, v := range apisData {
		if v.ApiID == id {
			return v
		}
	}
	return ApiData{}
}

func handleApiList(w http.ResponseWriter, r *http.Request) {
	apisData := getAPIs()

	tmpl, err := template.ParseFiles("./src/views/sidebar/apiTopicButton.html")
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		http.Error(w, "Failed to load template", http.StatusInternalServerError)
		return
	}

	for _, v := range apisData {
		err = tmpl.Execute(w, v)
		if err != nil {
			http.Error(w, "Failed to render template", http.StatusInternalServerError)
			fmt.Printf("Error: %s\n", err.Error())
			return
		}
	}
}
