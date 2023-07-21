package handlers

import (
	"fmt"
	"net/http"
	"text/template"

	api "github.com/vikingpingvin/hackweek-ssp/internal/web/api"
)

type apiDataGetterFunc func() []api.ApiData

var localDummyApiGetterFunc apiDataGetterFunc = func() []api.ApiData {
	apisData := []api.ApiData{
		{
			ApiID:   1,
			ApiName: "My Api 1",
			Message: "This is the content for API 1",
			ApiConfiguration: api.ApiConfiguration{
				Enabled:   true,
				TargetUrl: "http://localhost:8080",
				EndPoint:  "/api1",
			},
		},
		{
			ApiID:   2,
			ApiName: "Super API",
			Message: "This is the content for API 2",
			ApiConfiguration: api.ApiConfiguration{
				Enabled:   false,
				TargetUrl: "http://localhost:8081",
				EndPoint:  "/api2",
			},
		},
		{
			ApiID:   3,
			ApiName: "Animals API",
			Message: "The Animals API provides interesting scientifc facts on thousands of different animal species.",
			ApiConfiguration: api.ApiConfiguration{
				Enabled:   true,
				TargetUrl: "https://api.api-ninjas.com/v1/animals",
				EndPoint:  "/animals",
			},
		},
	}
	return apisData
}

func getApis(getterFunc apiDataGetterFunc) []api.ApiData {
	return getterFunc()
}

func GetApisWithLocalDummies() []api.ApiData {
	return getApis(localDummyApiGetterFunc)
}

func GetApiById(id int) api.ApiData {
	apisData := getApis(localDummyApiGetterFunc)
	for _, v := range apisData {
		if v.ApiID == id {
			return v
		}
	}
	return api.ApiData{}
}

func HandleApiList(w http.ResponseWriter, r *http.Request) {
	apisData := getApis(localDummyApiGetterFunc)

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
