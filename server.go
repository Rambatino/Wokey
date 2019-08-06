package main

import (
	"log"
	"net/http"
	"os"
	"wokey/routes/callback"
	"wokey/routes/home"
	"wokey/routes/login"
	"wokey/routes/logout"
	"wokey/routes/middlewares"
	"wokey/routes/user"

	"github.com/codegangsta/negroni"
	"github.com/dgraph-io/badger"
	"github.com/gorilla/handlers"
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
	r.PathPrefix("/static/").Handler(http.FileServer(http.Dir("./build")))
	r.HandleFunc("/manifest.json", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./public/manifest.json")
	}))
	r.HandleFunc("/favicon.ico", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./public/favicon.ico")
	}))
	http.Handle("/", handlers.CombinedLoggingHandler(os.Stdout, r))

	opts := badger.DefaultOptions("/tmp/wokey")
	db, err := badger.Open(opts)

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	hub := newHub(db)
	go hub.run()
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(hub, w, r)
	})

	// data things
	// r.HandleFunc("/issues", api.AllIssuesHandler)
	// r.HandleFunc("/pulls", api.AllCurrentPullRequests)

	// set up database manager
	// manager := database.NewManager()
	// manager.AddObserver("hi")
	// database.CheckForState("")
	port := os.Getenv("PORT")
	if port == "" {
		port = "1234"
	}
	log.Print("Server listening on http://0.0.0.0:" + port + "/")
	http.ListenAndServe("0.0.0.0:"+port, nil)
}
