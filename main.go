package main

import (
	"net/http"

	"github.com/GeertJohan/go.rice"
	"github.com/ggordan/govc/controllers"
	"github.com/ggordan/govc/watch"
)

func main() {

	go watch.WatchRepoForChanges()

	// Get the user dashboard with all the repositories
	http.HandleFunc("/dashboard", controllers.Dashboard)

	// Repository specific handlers
	http.HandleFunc("/commits", controllers.Commits)
	http.HandleFunc("/status", controllers.Status)
	http.HandleFunc("/branches", controllers.Branches)

	// Bundles all the assets into the Go binary
	http.Handle("/", http.FileServer(rice.MustFindBox("built/dev").HTTPBox()))

	// Start listening to requests
	http.ListenAndServe(":8090", nil)
}
