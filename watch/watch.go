package watch

import (
	"log"

	"github.com/googollee/go-socket.io"
	"gopkg.in/fsnotify.v1"
)

func WatchRepoForChanges(server *socketio.Server) {

	done := make(chan bool)
	changed := make(chan bool)

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	go func(c chan bool) {
		server.On("connection", func(so socketio.Socket) {
			log.Println("socket side")
			for {
				t := <-c
				if t {
					log.Println("change socket fired")
					so.Emit("news")
				}
			}
			so.On("disconnection", func() {
				watcher.Remove("/Users/ggordan/bootstrap")
			})
		})
		server.On("error", func(so socketio.Socket, err error) {
			log.Println("an error occurred")
			watcher.Remove("/Users/ggordan/bootstrap")
		})
	}(changed)

	go func(c chan bool) {
		log.Println("watcher side")
		for {
			select {
			case event := <-watcher.Events:
				c <- true
				if event.Op&fsnotify.Write == fsnotify.Write {
					log.Println("modified file:", event.Name)
				}
			case err := <-watcher.Errors:
				log.Println("error:", err)
			}
		}
	}(changed)

	err = watcher.Add("/Users/ggordan/bootstrap")
	if err != nil {
		log.Println("Watcher error occurred")
	}
	<-done

}
