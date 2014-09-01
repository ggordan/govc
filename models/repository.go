package models

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/ggordan/govc/conf"
	_ "github.com/mattn/go-sqlite3"
)

type Repository struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Location string `json:"location"`
}

// Create creates a new repository
func (r *Repository) Create(name, location string) {

	r.Name = name
	r.Location = location

}

func GetFromID(id string) *Repository {

	db, err := sql.Open("sqlite3", conf.GetDBPath())
	if err != nil {
		log.Fatal(err)
	}

	var repo Repository

	rows, err := db.Query("SELECT * FROM repos WHERE id = ?", id)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(rows.Columns())
	rows.Scan(&repo)

	return &repo
}
