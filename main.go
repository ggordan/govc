package main

import (
	"net/http"

	"github.com/GeertJohan/go.rice"
	"github.com/ggordan/govc/controllers"

	// "github.com/libgit2/git2go"
	// _ "github.com/mattn/go-sqlite3"
)

func main() {

	// The dashboard handler

	// Repository specific handlers
	http.HandleFunc("/commits", controllers.Commits)
	http.HandleFunc("/branches", controllers.Branches)

	http.Handle("/", http.FileServer(rice.MustFindBox("built/dev").HTTPBox()))

	// Start listening to requests
	http.ListenAndServe(":8090", nil)
}
