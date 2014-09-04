package controllers

import (
	"net/http"
	"time"

	git2go "github.com/ggordan/git2go"
	"github.com/ggordan/govc/models"
	"github.com/gorilla/mux"
)

// sig := &git2go.Signature{
// 	Name:  "Gasdas",
// 	Email: "asdsadsa@asdasd",
// 	When:  time.Now(),
// }

// repo, err := git2go.OpenRepository("/home/ggordan/bootstrap")
// if err != nil {
// 	panic(err)
// }

// reference, _ := repo.DwimReference("master")
// repo.SetHead(reference.Name(), sig, "new head")

// // oid, err := git2go.NewOid("0740332964f059e2f2f1eba9f52712786c919b6e")
// // if err != nil {
// // 	panic("opid")
// // }
// // reference, _ := repo.Head()
// // ref, err := reference.SetTarget(oid, sig, "asda")
// // if err != nil {
// // 	panic(err)
// // }

// var opts git2go.CheckoutOpts
// opts.Strategy = git2go.CheckoutForce
// repo.CheckoutHead(&opts)

// CheckoutBranch returns all the commits for the specified repository
func CheckoutBranch(res http.ResponseWriter, req *http.Request) {

	branchName := req.FormValue("b")

	// Get the repository details
	params := mux.Vars(req)
	repoDB, err := models.GetRepository(params["pid"])
	if err != nil {
		panic("missing")
	}

	repo, err := git2go.OpenRepository(repoDB.Location)
	if err != nil {
		panic(err)
	}

	config, _ := repo.Config()
	userName, _ := config.LookupString("user.name")
	userEmail, _ := config.LookupString("user.email")

	// Create the signature
	signature := &git2go.Signature{
		Name:  userName,
		Email: userEmail,
		When:  time.Now(),
	}

	reference, _ := repo.DwimReference(branchName)

	// t := reference.Target()
	// commit, _ := repo.LookupCommit(t)
	// tree, _ := commit.Tree()

	// Set the head to the new branch
	repo.SetHead(reference.Name(), signature, "branch changed")

	// Checkout the repo
	var opts git2go.CheckoutOpts
	opts.Strategy = git2go.CheckoutForce
	// repo.CheckoutTree(tree, &opts)
	repo.CheckoutHead(&opts)

}
