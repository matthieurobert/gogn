package spotify

import (
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/matthieurobert/gogn/api/config"
)

func CallbackHandler(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	state := r.URL.Query().Get("state")

	if state == "" {
		http.Redirect(w, r, "/#"+url.Values{"error": {"state_mismatch"}}.Encode(), http.StatusSeeOther)
	} else {
		clientID := config.CLIENT_ID         // Remplacez par votre client ID Spotify
		clientSecret := config.CLIENT_SECRET // Remplacez par votre client secret Spotify
		redirectURI := config.REDIRECT_URI

		authOptions := url.Values{}
		authOptions.Set("code", code)
		authOptions.Set("redirect_uri", redirectURI)
		authOptions.Set("grant_type", "authorization_code")

		formData := authOptions.Encode()
		req, _ := http.NewRequest("POST", "https://accounts.spotify.com/api/token", strings.NewReader(formData))
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		authHeader := "Basic " + base64.StdEncoding.EncodeToString([]byte(clientID+":"+clientSecret))
		req.Header.Set("Authorization", authHeader)

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			// Gérer l'erreur
			return
		}
		defer resp.Body.Close()

		// Gérer la réponse de l'API Spotify
		// Le contenu de la réponse se trouve dans resp.Body

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Fprint(w, string(body))
	}
}
