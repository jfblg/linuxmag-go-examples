package main

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
)

func main() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		panic(err)
	}
	defer watcher.Close()

	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				fmt.Printf("%+v\n", event)
			}
		}
	}()

	err = watcher.Add("/tmp/test")
	if err != nil {
		panic(err)
	}

	done := make(chan struct{})
	<-done
}
