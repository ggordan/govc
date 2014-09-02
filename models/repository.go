package models

import (
	"fmt"

	"github.com/ggordan/govc/database"
	_ "github.com/mattn/go-sqlite3" // Blank import to add support for SQLite
)

// RepositoryTableName is the name of the tables in SQLite
const RepositoryTableName string = "repos"

// Repository corresponds to a single row in the repository table
type Repository struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Location string `json:"location"`
}

// CreateRepository a new repository
func CreateRepository(name, location string) (repository *Repository, err error) {

	// Set the repository data
	repository.Name = name
	repository.Location = location

	createQuery := fmt.Sprintf("INSERT INTO %s(name, location) VALUES(?, ?)", RepositoryTableName)

	// Get the database
	db, err := database.GetDB()
	if err != nil {
		return nil, err
	}

	// Save the new repository to the database
	result, err := db.Exec(createQuery, name, location)
	if err != nil {
		return nil, err
	}

	// Retrieve the ID of the newly inserted repository
	repository.ID, err = result.LastInsertId()
	if err != nil {
		return nil, err
	}

	return repository, nil
}

// GetRepository retries the repository with the id
func GetRepository(searchID int) (repository Repository, err error) {

	getQuery := fmt.Sprintf("SELECT * FROM %s WHERE id = ?", RepositoryTableName)

	// Get the database
	db, err := database.GetDB()
	if err != nil {
		return Repository{}, err
	}

	row := db.QueryRow(getQuery, searchID)

	row.Scan(&repository.ID, &repository.Name, &repository.Location)

	return repository, nil
}

// GetAllRepositories returns all the repositories that the user has created
func GetAllRepositories() (repositories []*Repository, err error) {

	allQuery := fmt.Sprintf("SELECT * FROM %s", RepositoryTableName)

	// Get the database
	db, err := database.GetDB()
	if err != nil {
		return nil, err
	}

	rows, err := db.Query(allQuery)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var repository Repository
		rows.Scan(&repository.ID, &repository.Name, &repository.Location)
		repositories = append(repositories, &repository)
	}

	return repositories, nil
}
