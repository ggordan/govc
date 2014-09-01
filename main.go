package main

import (
	"net/http"

	"github.com/GeertJohan/go.rice"
	"github.com/ggordan/govc/controllers"
	"github.com/gorilla/mux"

	// "github.com/libgit2/git2go"
	// _ "github.com/mattn/go-sqlite3"
)

func main() {

	router := mux.NewRouter()

	// The dashboard handler
	router.HandleFunc("/", controllers.Dashboard)

	// Repository specific handlers
	router.HandleFunc("/{RID}/commits", controllers.Commits)
	router.HandleFunc("/{RID}/branches", controllers.Branches)

	router.Handle("/a", http.StripPrefix("/a", http.FileServer(rice.MustFindBox("built/dev").HTTPBox())))

	// Start listening to requests
	http.ListenAndServe(":8090", router)
}
