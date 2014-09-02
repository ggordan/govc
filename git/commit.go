package git

import (
	"crypto/md5"
	"encoding/hex"
	"time"

	git2go "github.com/ggordan/git2go"
)

type signature struct {
	Name      string
	Email     string
	MD5       string
	Timestamp time.Time
}

type CommitJSON struct {
	SHA1      string
	Author    signature
	Committer signature
	Message   string
	Timestamp time.Time
	Diff      DiffJSON
}

func (cj *CommitJSON) Generate(commit *git2go.Commit) {

	cj.SHA1 = commit.Id().String()

	// Get the email MD5 for Gravatar support
	authorHash := md5.New()
	authorHash.Write([]byte(commit.Author().Email))

	cj.Author = signature{
		Name:      commit.Author().Name,
		Email:     commit.Author().Email,
		MD5:       hex.EncodeToString(authorHash.Sum(nil)),
		Timestamp: commit.Author().When,
	}

	commiterHash := md5.New()
	commiterHash.Write([]byte(commit.Committer().Email))

	cj.Committer = signature{
		Name:      commit.Committer().Name,
		Email:     commit.Committer().Email,
		MD5:       hex.EncodeToString(commiterHash.Sum(nil)),
		Timestamp: commit.Committer().When,
	}

	cj.Message = commit.Message()
	cj.Timestamp = commit.Author().When

	// Retrieve the Diff
	var diffJSON DiffJSON
	diffJSON.Generate(commit)
	cj.Diff = diffJSON
}
