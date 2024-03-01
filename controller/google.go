package controller

import (
	"context"
	"fmt"
	"go-omniauth-google/config"
	"io"
	"net/http"
)

func GoogleLogin(res http.ResponseWriter, req *http.Request) {
	googleConfig := config.SetupConfig()
	url := googleConfig.AuthCodeURL("randomstate")

	// redirect to google login page
	http.Redirect(res, req, url, http.StatusSeeOther)
}

func GoogleCallback(res http.ResponseWriter, req *http.Request) {
	// state
	state := req.URL.Query()["state"][0]
	if state != "randomstate" {
		fmt.Println(res, "state not match")
		return
	}

	// code
	code := req.URL.Query()["code"][0]

	// configuration
	googleConfig := config.SetupConfig()

	// exchange code for token
	token, err := googleConfig.Exchange(context.Background(), code)
	if err != nil {
		fmt.Println("Code exchang fail")
	}

	// use google api to getuser info
	resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		fmt.Println(res, "Data fetch failed")
	}

	// parse respone
	userData, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(res, "Json parsing failed")
	}
	fmt.Fprintln(res, string(userData))
}
