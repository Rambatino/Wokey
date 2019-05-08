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
	r.HandleFunc("/", home.AppHandler)
	r.HandleFunc("/login", login.LoginHandler)
	r.HandleFunc("/logout", logout.LogoutHandler)
	r.HandleFunc("/login/callback", callback.CallbackHandler)
	r.Handle("/user", negroni.New(
		negroni.HandlerFunc(middlewares.IsAuthenticated),
		negroni.Wrap(http.HandlerFunc(user.UserHandler)),
	))
	r.PathPrefix("/static/").Handler(http.FileServer(http.Dir("build/")))
	http.Handle("/", r)

	// data things
	r.HandleFunc("/issues", api.AllIssuesHandler)
	r.HandleFunc("/pulls", api.AllCurrentPullRequests)

	log.Print("Server listening on http://localhost:1234/")
	http.ListenAndServe("0.0.0.0:1234", nil)
}
