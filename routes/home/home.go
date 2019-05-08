package home

import (
	"net/http"
)

// func HomeHandler(w http.ResponseWriter, r *http.Request) {
// 	data := struct {
// 		Auth0ClientId     string
// 		Auth0ClientSecret string
// 		Auth0Domain       string
// 		Auth0CallbackURL  template.URL
// 	}{
// 		os.Getenv("AUTH0_CLIENT_ID"),
// 		os.Getenv("AUTH0_CLIENT_SECRET"),
// 		os.Getenv("AUTH0_DOMAIN"),
// 		template.URL(os.Getenv("AUTH0_CALLBACK_URL")),
// 	}
// 	templates.RenderTemplate(w, "home", data)
// }

func AppHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "build/index.html")
}
