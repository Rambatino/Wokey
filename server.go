package main

import (
	"log"
	"net/http"
	"wokey/routes/api"
	"wokey/routes/callback"
	"wokey/routes/home"
	"wokey/routes/login"
	"wokey/routes/logout"
	"wokey/routes/middlewares"
	"wokey/routes/user"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func StartServer() {
	r := mux.NewRouter()

	// go to / and if not logged in, redirect to home page, if logged in, then show
	// the app itself

	r.Handle("/", negroni.New(
		negroni.HandlerFunc(middlewares.IsAuthenticated),
		negroni.Wrap(http.HandlerFunc(home.AppHandler)),
	))
	r.HandleFunc("/home", home.HomeHandler)
	r.HandleFunc("/login", login.LoginHandler)
	r.HandleFunc("/logout", logout.LogoutHandler)
	r.HandleFunc("/callback", callback.CallbackHandler)
	r.Handle("/user", negroni.New(
		negroni.HandlerFunc(middlewares.IsAuthenticated),
		negroni.Wrap(http.HandlerFunc(user.UserHandler)),
	))
	r.PathPrefix("/auth0Login/").Handler(http.StripPrefix("/auth0Login/", http.FileServer(http.Dir("auth0Login/"))))
	r.PathPrefix("/static/").Handler(http.FileServer(http.Dir("build/")))
	r.PathPrefix("/manifest.json").HandlerFunc(func(w http.ResponseWriter, h *http.Request) {
		http.ServeFile(w, h, "build/manifest.json")
	})
	http.Handle("/", r)
	r.HandleFunc("/all-issues", api.AllIssuesHandler)
	log.Print("Server listening on :3000")
	http.ListenAndServe("0.0.0.0:3000", nil)
}
