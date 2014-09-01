package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/ggordan/govc/git"
	"github.com/gorilla/mux"
	git2go "github.com/libgit2/git2go"
)

const defaultCommitsPerPage int = 10

type branchJSON struct {
	Name   string
	Head   bool
	Target string
	Remote bool
	Tag    bool
}

type metaStruct struct {
	Branches []branchJSON
}

// Commits returns all the commits for the specified repository
func Branches(res http.ResponseWriter, req *http.Request) {
	var meta metaStruct
	repo, err := git2go.OpenRepository("/home/ggordan/bootstrap")
	if err != nil {
		panic(err)
	}

	// Iterate through all the branch types
	for _, branchType := range []git2go.BranchType{git2go.BranchRemote, git2go.BranchLocal} {
		// Get all the branches
		branches, _ := repo.NewBranchIterator(branchType)

		branch, _, _ := branches.Next()
		for branch != nil {

			var isRemote bool
			if branchType == git2go.BranchRemote {
				isRemote = true
			}

			// Get the branch target SHA1
			var targetString string
			targetOid := branch.Target()
			if targetOid != nil {
				targetString = targetOid.String()
			}

			branchName, _ := branch.Name()
			isHead, _ := branch.IsHead()
			meta.Branches = append(meta.Branches, branchJSON{
				Name:   branchName,
				Target: targetString,
				Head:   isHead,
				Remote: isRemote,
				Tag:    branch.IsTag(),
			})
			branch, _, err = branches.Next()

		}
	}

	b, _ := json.Marshal(meta)
	res.Write(b)
}

// Commits returns all the commits for the specified repository
func Commits(res http.ResponseWriter, req *http.Request) {
	var commits []git.CommitJSON
	var commitJSON git.CommitJSON
	var i int

	// Ignore this error as we fall back to page 0 when p is empty
	pageNumber, _ := strconv.Atoi(req.FormValue("p"))

	// Get the number of commits to display per page
	commitsPerPage, err := strconv.Atoi(req.FormValue("commits"))
	if err != nil {
		commitsPerPage = defaultCommitsPerPage
	}

	log.Println(pageNumber, commitsPerPage)

	// Get the request variables from the URL
	_ = mux.Vars(req)

	repo, err := git2go.OpenRepository("/home/ggordan/bootstrap")
	if err != nil {
		panic(err)
	}

	odb, err := repo.Odb()
	if err != nil {
		panic(err)
	}

	odb.ForEach(func(id *git2go.Oid) error {
		if i > commitsPerPage || id.IsZero() {
			return nil
		}

		// Retrieve a commit
		commit, err := repo.LookupCommit(id)
		if err != nil {
			return err
		}

		i++
		commitJSON.Generate(commit)
		commits = append(commits, commitJSON)
		return nil
	})

	marshaledCommitJSON, _ := json.Marshal(commits)
	res.Write(marshaledCommitJSON)
}
