package fsnotify

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
)

func CallFSNotify() {

	watcher, err := fsnotify.NewWatcher()

	if err != nil {
		fmt.Println(err)
	}

	defer watcher.Close()

	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				if event.Has(fsnotify.Write) {
					fmt.Println("MOdified File", event.Name)
				}

			case err := <-watcher.Errors:
				if err != nil {
					fmt.Println("Got errors", err)
				}

			}

		}
	}()

	if err := watcher.Add("tmp"); err != nil {
		fmt.Println(err)

	}

	if err = watcher.Add("main.go"); err != nil {
		fmt.Println("Err", err)

	}

	// block this data forever
	<-make(chan struct{})
}
