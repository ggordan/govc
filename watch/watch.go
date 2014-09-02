package watch

import (
	"log"
	"net/http"

	"github.com/googollee/go-socket.io"
	"gopkg.in/fsnotify.v1"
)

func WatchRepoForChanges() {

	done := make(chan bool)

	server, err := socketio.NewServer(nil)
	if err != nil {
		log.Fatal(err)
	}

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	server.On("connection", func(so socketio.Socket) {
		go func() {
			log.Println("error:")

			for {
				select {
				case event := <-watcher.Events:
					if event.Op&fsnotify.Write == fsnotify.Write {
						so.Emit("Changed")
						log.Println("modified file:", event.Name)
					}
				case err := <-watcher.Errors:
					log.Println("error:", err)
				}
			}
		}()
	})

	err = watcher.Add("/home/ggordan/bootstrap")
	if err != nil {
		log.Fatal(err)
	}
	<-done

	http.Handle("/socket.io/", server)
}
