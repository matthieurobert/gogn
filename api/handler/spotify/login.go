package spotify

import (
	"net/http"
	"net/url"

	"github.com/matthieurobert/gogn/api/config"
	"github.com/matthieurobert/gogn/api/utils"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	state := utils.GenerateRandomString(16)
	scope := "user-read-private user-read-email"

	values := url.Values{}
	values.Set("response_type", "code")
	values.Set("client_id", config.CLIENT_ID) // Remplacez "CLIENT_ID" par votre propre client_id
	values.Set("scope", scope)
	values.Set("redirect_uri", config.CLIENT_SECRET)
	values.Set("state", state)

	redirectURL := "https://accounts.spotify.com/authorize?" + values.Encode()
	http.Redirect(w, r, redirectURL, http.StatusSeeOther)
}
