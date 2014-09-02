package controllers

import (
	"encoding/json"
	"net/http"

	git2go "github.com/ggordan/git2go"
	"github.com/ggordan/govc/models"
)

type dashboardJSON struct {
	Repo   *models.Repository
	Exists bool
}

// Dashboard re all the repositories that the user has created
func Dashboard(res http.ResponseWriter, req *http.Request) {
	var repositoriesJSON []dashboardJSON

	// Get all the repositories that the user has created
	repositories, err := models.GetAllRepositories()
	if err != nil {
		panic(err)
	}

	// go through all the repositories and verify that they actually exist in the
	// database
	for _, repo := range repositories {
		var dashboardRepo dashboardJSON
		dashboardRepo.Repo = repo
		r, _ := git2go.OpenRepository(repo.Location)
		if r != nil {
			dashboardRepo.Exists = true
		}
		repositoriesJSON = append(repositoriesJSON, dashboardRepo)
	}

	b, err := json.Marshal(repositoriesJSON)
	if err != nil {
		panic(err)
	}

	res.Write(b)
}
