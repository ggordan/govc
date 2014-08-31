package main

import (
	"log"
	"net/http"

	"github.com/GeertJohan/go.rice"
	"github.com/ggordan/govc/controllers"
	"github.com/googollee/go-socket.io"
	// "github.com/libgit2/git2go"
	// _ "github.com/mattn/go-sqlite3"
)

func main() {

	http.Handle("/", http.FileServer(rice.MustFindBox("built/dev").HTTPBox()))

	server, err := socketio.NewServer(nil)
	if err != nil {
		log.Fatal(err)
	}
	server.On("connection", func(so socketio.Socket) {
		log.Println("on connection")
		so.Join("chat")
		so.On("chat message", func(msg string) {
			log.Println("emit:", so.Emit("chat message", msg))
			so.BroadcastTo("chat", "chat message", msg)
		})
		so.On("disconnection", func() {
			log.Println("on disconnect")
		})
	})
	server.On("error", func(so socketio.Socket, err error) {
		log.Println("error:", err)
	})

	http.HandleFunc("/meta", controllers.MetaHandler)
	http.HandleFunc("/commits", controllers.CommitsHandler)
	http.Handle("/socket.io/", server)
	http.ListenAndServe(":8080", nil)

	// db, err := sql.Open("sqlite3", "~/.gogi.db")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// repo, err := git.OpenRepository("/home/ggordan/bootstrap")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("---", repo.Workdir())
	// odb, err := repo.Odb()
	// if err != nil {
	// 	panic(err)
	// }
	// // iterate branches
	// branches, err := repo.NewBranchIterator(git.BranchLocal)
	// if err != nil {
	// 	panic(err)
	// }
	// for {
	// 	branch, branchType, err := branches.Next()
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	fmt.Println(branchType)
	// 	fmt.Println(branch.Name())
	// }
	// odb.ForEach(func(id *git.Oid) error {
	// 	// obj, _ := repo.Lookup(id)
	// 	// if obj.Type() == git.ObjectCommit {
	// 	// 	commit, _ := repo.LookupCommit(id)
	// 	// 	commits = append(commits, commit)
	// 	// 	fmt.Println("%v", commit.Author().Name)
	// 	// 	// tree, err := commit.Tree()
	// 	// 	// if err != nil {
	// 	// 	// 	panic(err)
	// 	// 	// }
	// 	// 	// fmt.Println(tree)
	// 	// }
	// 	return nil
	// })
}
