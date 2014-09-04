package main

import (
	"net/http"

	"github.com/GeertJohan/go.rice"
	"github.com/ggordan/govc/controllers"
	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()

	// Get the user dashboard with all the repositories
	router.HandleFunc("/dashboard", controllers.Dashboard)

	// Repository specific handlers
	router.HandleFunc("/api/{pid}/commits", controllers.Commits)
	router.HandleFunc("/api/{pid}/status", controllers.Status)
	router.HandleFunc("/api/{pid}/branches", controllers.Branches)

	// Branch specific handlers
	router.HandleFunc("/api/{pid}/branches/checkout", controllers.CheckoutBranch)

	mux := http.NewServeMux()
	mux.Handle("/", http.StripPrefix("/", http.FileServer(rice.MustFindBox("built/dev").HTTPBox())))
	mux.Handle("/api/", router)

	router.Handle("/static", http.FileServer(rice.MustFindBox("built/dev").HTTPBox()))

	http.ListenAndServe(":8090", mux)

	// Bundles all the assets into the Go binary

	// Start listening to requests
	http.ListenAndServe(":8090", nil)
}
