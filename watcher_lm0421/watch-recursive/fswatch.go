package main

import (
	"fmt"
	"log"
	"os"
	"os/user"
	"path/filepath"
	"strings"

	"github.com/fsnotify/fsnotify"
)

func main() {
	cur, err := user.Current()
	dieOnErr(err)
	home := cur.HomeDir

	watcher, err := fsnotify.NewWatcher()
	dieOnErr(err)
	defer watcher.Close()

	watchInit(watcher)

	err = filepath.Walk(filepath.Join(home, "go"), func(path string, info os.FileInfo, err error) error {
		dieOnErr(err)
		if info.IsDir() {
			err := watcher.Add(path)
			dieOnErr(err)
		}
		return nil
	})
	dieOnErr(err)

	done := make(chan interface{})
	<-done
}

func eventAsString(event fsnotify.Event) string {
	info, err := os.Stat(event.Name)
	dieOnErr(err)
	evShort := (strings.ToLower(event.Op.String()))[0:2]

	dirParts := strings.Split(event.Name, "/")
	pathShort := event.Name
	if len(dirParts) > 3 {
		pathShort = filepath.Join(dirParts[len(dirParts)-3 : len(dirParts)]...)
	}
	return fmt.Sprintf("%s %s %d", evShort, pathShort, info.Size())
}

func watchInit(watcher *fsnotify.Watcher) {
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				if event.Op&fsnotify.Rename == fsnotify.Rename || event.Op&fsnotify.Remove == fsnotify.Remove {
					continue
				}
				log.Printf("%s\n", eventAsString(event))
				info, err := os.Stat(event.Name)
				dieOnErr(err)
				if info.IsDir() {
					err := watcher.Add(event.Name)
					dieOnErr(err)
				}
			case err, _ := <-watcher.Errors:
				panic(err)
			}
		}
	}()
}

func dieOnErr(err error) {
	if err != nil {
		panic(err)
	}
}
