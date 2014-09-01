package controllers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Branches retrieves the Remote/Local branches and tags for the specified
// repository.
func Branches(res http.ResponseWriter, req *http.Request) {

	// Get the request variables from the URL
	params := mux.Vars(req)

	fmt.Println(params)
}
