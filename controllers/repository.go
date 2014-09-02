package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	git2go "github.com/ggordan/git2go"
	"github.com/ggordan/govc/git"
	"github.com/gorilla/mux"
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

// Branches returns all the commits for the specified repository
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

// Status returns all the commits for the specified repository
func Status(res http.ResponseWriter, req *http.Request) {

	type statusJSON struct {
		Kind  string
		Entry interface{}
	}

	var b []statusJSON

	repo, err := git2go.OpenRepository("/home/ggordan/bootstrap")
	if err != nil {
		panic(err)
	}

	var so git2go.StatusOptions

	statusList, _ := repo.StatusList(&so)

	entries, _ := statusList.EntryCount()

	for i := 0; i < entries; i++ {
		var t string
		entry, _ := statusList.ByIndex(i)

		switch entry.Status {
		case git2go.StatusIndexNew:
			t = "index_new"
		case git2go.StatusWtModified:
			fallthrough
		case git2go.StatusIndexModified:
			t = "index_modified"
		case git2go.StatusWtDeleted:
			fallthrough
		case git2go.StatusIndexDeleted:
			t = "index_deleted"
		}

		b = append(b, statusJSON{
			Kind:  t,
			Entry: entry,
		})
	}

	fmt.Println("DATA", statusList, entries)

	bb, _ := json.Marshal(b)

	res.Write(bb)
}
