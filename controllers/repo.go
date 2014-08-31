package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/ggordan/govc/git"
	git2go "github.com/libgit2/git2go"
)

type branchJSON struct {
	Name   string
	Head   bool
	Target string
	Tag    bool
}

type metaStruct struct {
	Branches []branchJSON
}

func MetaHandler(res http.ResponseWriter, req *http.Request) {
	var meta metaStruct

	repo, _ := git2go.OpenRepository("/Users/ggordan/bootstrap")

	// iterate remote Branches
	branches, err := repo.NewBranchIterator(git2go.BranchLocal)
	if err == nil {
		for {
			branch, _, err := branches.Next()
			if err != nil {
				break
			}
			branchName, _ := branch.Name()
			isHead, _ := branch.IsHead()
			meta.Branches = append(meta.Branches, branchJSON{
				Name:   branchName,
				Target: branch.Target().String(),
				Head:   isHead,
				Tag:    branch.IsTag(),
			})
		}
	}

	// iterate remote Branches
	branches, err = repo.NewBranchIterator(git2go.BranchRemote)
	if err == nil {
		for {
			branch, _, err := branches.Next()
			if err != nil {
				break
			}
			branchName, _ := branch.Name()
			isHead, _ := branch.IsHead()
			meta.Branches = append(meta.Branches, branchJSON{
				Name:   branchName,
				Target: branch.Target().String(),
				Head:   isHead,
				Tag:    branch.IsTag(),
			})
		}
	}

	b, _ := json.Marshal(meta)
	res.Write(b)
}

// CommitsHandler returns all the commits for the specified repository
func CommitsHandler(res http.ResponseWriter, req *http.Request) {
	var commits []git.CommitJSON

	repo, err := git2go.OpenRepository("/home/ggordan/GutterColor")
	if err != nil {
		panic(err)
	}

	odb, err := repo.Odb()
	if err != nil {
		panic(err)
	}

	var i int
	odb.ForEach(func(id *git2go.Oid) error {
		var commitJSON git.CommitJSON
		// Testing, only one commit
		i++
		if i > 50 {
			return nil
		}

		// Retrieve a commit
		commit, err := repo.LookupCommit(id)
		if err != nil {
			return err
		}

		commitJSON.Generate(commit)
		commits = append(commits, commitJSON)
		return nil
	})

	b, _ := json.Marshal(commits)
	res.Write(b)
}
