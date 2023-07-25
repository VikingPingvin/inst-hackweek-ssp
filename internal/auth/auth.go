package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

const (
	clientID     = "456160238646-pmp69tleqtm74bthvgo7p5p05cf9gqv2.apps.googleusercontent.com"
	clientSecret = "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"
)

var OauthConfig = oauth2.Config{
	ClientID:     clientID,
	ClientSecret: clientSecret,
	RedirectURL:  "http://localhost:8080/auth/callback",
	Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
	Endpoint:     google.Endpoint,
}

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Handling login: \n")

	url := OauthConfig.AuthCodeURL("state", oauth2.AccessTypeOffline)
	http.Redirect(w, r, url, http.StatusFound)
}

func HandleCallback(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Handling login callback: \n")

	code := r.URL.Query().Get("code")
	if code == "" {
		http.Error(w, "Missing code parameter", http.StatusBadRequest)
		return
	}

	token, err := OauthConfig.Exchange(context.Background(), code)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to exchange token: %s", err), http.StatusInternalServerError)
		return
	}

	GetUserDetails(token)

	loginCookie := http.Cookie{
		Name:     "gateway-ssp-login",
		Value:    token.AccessToken,
		Path:     "/",
		MaxAge:   3600,
		HttpOnly: true,
		SameSite: http.SameSiteDefaultMode,
	}
	http.SetCookie(w, &loginCookie)
	// w.Write([]byte("Login successful"))
	fmt.Printf("Cookie header: %s", w.Header().Get("Set-Cookie"))

	http.Redirect(w, r, "/", http.StatusFound)
}

func GetUserDetails(accessToken *oauth2.Token) error {
	client := OauthConfig.Client(context.Background(), accessToken)

	resp, err := client.Get("https://www.googleapis.com/oauth2/v3/userinfo")
	if err != nil {
		return fmt.Errorf("Failed to get user details: %s", err)
	}
	defer resp.Body.Close()

	var userInfo struct {
		Email string `json:"email"`
		Name  string `json:"name"`
	}

	err = json.NewDecoder(resp.Body).Decode(&userInfo)
	if err != nil {
		return err
	}

	fmt.Printf("User login details: %v\n", userInfo)
	fmt.Println(resp.Status)
	return nil
}

func VerifyUserTokenOnline(token *oauth2.Token) error {

	r, err := OauthConfig.Client(context.Background(), token).Get("https://www.googleapis.com/oauth2/v3/tokeninfo")
	if err != nil {
		return fmt.Errorf("Failed to get user details: %s", err)
	}
	defer r.Body.Close()

	var tokenInfo map[string]interface{}
	err = json.NewDecoder(r.Body).Decode(&tokenInfo)
	if err != nil {
		return fmt.Errorf("Failed to decode token info: %s", err)
	}

	if r.StatusCode != http.StatusOK {
		return fmt.Errorf("access token not valid, response statuscode: %s", r.StatusCode)
	}

	return nil
}

func VerifyUserTokenOffline(token *oauth2.Token) error {

	tokenParsed, err := jwt.Parse(token.AccessToken, nil)
	if err != nil {
		return fmt.Errorf("Failed to parse token: %s", err)
	}

	if !tokenParsed.Valid {
		return fmt.Errorf("access token not valid")
	}

	if tokenParsed.Claims.(jwt.MapClaims)["aud"] != clientID {
		return fmt.Errorf("audience not valid")
	}

	return nil
}

func GetAuthCookie(r *http.Request) (*http.Cookie, error) {
	cookie, err := r.Cookie("gateway-ssp-login")
	if err != nil {
		return nil, fmt.Errorf("Failed to get cookie: %s", err)
	}
	return cookie, nil
}
