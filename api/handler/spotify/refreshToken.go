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

func RefreshTokenHandler(w http.ResponseWriter, r *http.Request) {
	refreshToken := r.URL.Query().Get("refresh_token")
	clientID := config.CLIENT_ID         // Remplacez par votre client ID Spotify
	clientSecret := config.CLIENT_SECRET // Remplacez par votre client secret Spotify

	authOptions := url.Values{}
	authOptions.Set("grant_type", "refresh_token")
	authOptions.Set("refresh_token", refreshToken)

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

	// Si la réponse est 200 OK
	if resp.StatusCode == http.StatusOK {
		// Gérer la réponse de l'API Spotify
		// Le contenu de la réponse se trouve dans resp.Body
		// Ici, vous pouvez traiter la réponse pour récupérer access_token et refresh_token
		// puis les renvoyer dans la réponse HTTP
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Fprint(w, string(body))
	}

	// Si la réponse n'est pas 200 OK, vous pouvez gérer d'autres statuts de réponse ici
	// Par exemple, renvoyer un message d'erreur ou gérer d'autres scénarios
	fmt.Fprint(w, "erro refresh token")
}
