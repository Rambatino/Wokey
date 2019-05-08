package logout

import (
	"net/http"
	"net/url"
	"os"

	"github.com/k0kubun/pp"
)

func LogoutHandler(w http.ResponseWriter, r *http.Request) {

	domain := os.Getenv("AUTH0_DOMAIN")

	var Url *url.URL
	Url, err := url.Parse("https://" + domain)

	if err != nil {
		panic("boom")
	}

	Url.Path += "/v2/logout"
	parameters := url.Values{}
	parameters.Add("returnTo", "http://localhost:1234")
	parameters.Add("client_id", os.Getenv("AUTH0_CLIENT_ID"))
	Url.RawQuery = parameters.Encode()

	pp.Println(Url.String())

	http.Redirect(w, r, Url.String(), http.StatusTemporaryRedirect)
}
