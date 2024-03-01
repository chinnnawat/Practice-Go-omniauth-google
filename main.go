package main

import (
	"go-omniauth-google/controller"
	"net/http"
)

func main() {
	http.Handle("/google/login", http.HandlerFunc(controller.GoogleLogin))
	http.Handle("/google/callback", http.HandlerFunc(controller.GoogleCallback))
	http.ListenAndServe(":3000", nil)
}
