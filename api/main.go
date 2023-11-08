package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/matthieurobert/gogn/api/handler/spotify"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/login", spotify.LoginHandler)
	r.HandleFunc("/callback", spotify.CallbackHandler)
	r.HandleFunc("/refresh_token", spotify.RefreshTokenHandler)

	http.ListenAndServe(":8080", r)
}
