package controllers

import (
	"database/sql"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"

	"github.com/ggordan/govc/conf"
	_ "github.com/mattn/go-sqlite3"
)

// Dashboard s
func Dashboard(res http.ResponseWriter, req *http.Request) {

	// Get the directory of the SQLITE database
	databasePath := conf.GetDBPath()
	schemaPath := path.Join(os.Getenv("PWD"), "schema.sql")

	db, err := sql.Open("sqlite3", databasePath)
	if err != nil {
		log.Fatal(err)
	}

	content, err := ioutil.ReadFile(schemaPath)

	_, err = db.Exec(string(content))
	if err != nil {
		log.Printf("%q: %s\n", err, string(content))
		return
	}
}
