package handlers

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/vikingpingvin/hackweek-ssp/internal/auth"
	api "github.com/vikingpingvin/hackweek-ssp/internal/web/api"
	"golang.org/x/oauth2"
)

type apiDataGetterFunc func() []api.ApiData

var localDummyApiGetterFunc apiDataGetterFunc = func() []api.ApiData {
	apisData := []api.ApiData{
		{
			ApiID:   1,
			ApiName: "My Api 1",
			Message: "Please imagine this is 100s of lines of openapi spec schema \n",
			ApiConfiguration: api.ApiConfiguration{
				Enabled:   true,
				TargetUrl: "http://localhost:8080",
				EndPoint:  "/api1",
			},
		},
		{
			ApiID:   2,
			ApiName: "Super API",
			Message: "Totally an OpenAPI spec for API 2",
			ApiConfiguration: api.ApiConfiguration{
				Enabled:   false,
				TargetUrl: "http://localhost:8081",
				EndPoint:  "/api2",
			},
		},
		{
			ApiID:   3,
			ApiName: "Animals API",
			Message: "I'm not going to paste 100s of lines into dummy hardcoded strings....\nThe Animals API provides interesting scientifc facts on thousands of different animal species.",
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
	var apisData []api.ApiData
	if cookie, err := auth.GetAuthCookie(r); err == nil {
		auth.VerifyUserTokenOnline(&oauth2.Token{
			AccessToken: cookie.Value,
		})
		apisData = getApis(localDummyApiGetterFunc)
	} else {
		apisData = []api.ApiData{}
	}

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

func HandleApiContent(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost || r.Method == http.MethodPut {

		apiId := r.URL.Path[len("/api/"):]
		fmt.Printf("Handling API setting change for api-id: %s\n", apiId)

		r.ParseForm()

		for key, value := range r.Form {
			fmt.Printf("   Key: %s, Value: %s\n", key, value)
		}

		w.WriteHeader(http.StatusOK)
	} else if r.Method == http.MethodGet {
		fmt.Printf(("Handling GET for api: %s\n"), r.URL.Path[len("/api/"):])
	}
}
