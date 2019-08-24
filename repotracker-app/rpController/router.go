package rpController

import (
	"github.com/gorilla/mux"
	"github.com/thaniri/repotracker-app/rpController/login"
	"github.com/thaniri/repotracker-app/rpController/static"
	"net/http"
)

// Constructs are router for the website.
func New() http.Handler {
	router := mux.NewRouter()

	// Login Controllers
	router.HandleFunc("/{loginPost:loginPost\\/?}", login.LoginPostHandler).Methods("POST")
	router.HandleFunc("/{logout:logout\\/?}", login.LogoutPostHandler).Methods("POST")
	router.HandleFunc("/{register:register\\/?}", login.RegisterHandler).Methods("POST")
	router.HandleFunc("/{internal:internal\\/?}", login.InternalPageHandler)

	// TODO: model related Controllers

	// Static Content
	router.HandleFunc("/{login:login\\/?}", static.LoginHandler)
	router.HandleFunc("/", static.IndexHandler)

	return router
}
