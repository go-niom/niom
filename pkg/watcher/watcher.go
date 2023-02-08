package watcher

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/go-niom/niom/pkg/terminal"
	"github.com/gookit/color"
)

var (
	WatcherChannel chan string
)

func init() {
	WatcherChannel = make(chan string, 1000)
}

func rebuildCheck() {
	go func() {
		for {
			eventName := <-WatcherChannel
			color.Greenp("Modified File: ")
			println(eventName)
			terminal.IsCodeUpdated = true
			terminal.KillFunc()
		}
	}()
}

func watchFolder(path string) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				if event.Has(fsnotify.Write) {
					WatcherChannel <- event.String()
				}

			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	// watcherLog("Watching %s", path)
	err = watcher.Add(path)

	if err != nil {
		log.Fatal(err)
	}
}

func Watch() {
	rebuildCheck()
	root := "."
	filepath.Walk(root, func(path string, info os.FileInfo, err error) error {

		if info.IsDir() { //&& !isTmpDir(path)
			if len(path) > 1 && strings.HasPrefix(filepath.Base(path), ".") {
				return filepath.SkipDir
			}

			// if isIgnoredFolder(path) {
			// 	watcherLog("Ignoring %s", path)
			// 	return filepath.SkipDir
			// }
			watchFolder(path)
		}

		return err
	})
}
