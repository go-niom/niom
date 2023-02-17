package watcher

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/go-niom/niom/pkg/constants"
	"github.com/go-niom/niom/pkg/logger"
	"github.com/go-niom/niom/pkg/terminal"
	"github.com/gookit/color"
)

var (
	WatcherChannel chan string
)

// init initializes the channel to watch changes in the files
func init() {
	WatcherChannel = make(chan string, 1000)
}

func rebuildCheck() {
	go func() {
		for {
			eventName := <-WatcherChannel
			color.Bluep(constants.AppSign)
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
	err = watcher.Add(path)

	if err != nil {
		log.Fatal(err)
	}
}

// Watch keep tracks of the files and notify when there are any changes in the files
func Watch() {
	logger.Info("watching path(s): *.*")
	logger.Info("watching extensions: *")
	rebuildCheck()
	root := "."
	filepath.Walk(root, func(path string, info os.FileInfo, err error) error {

		if info.IsDir() {
			if len(path) > 1 && strings.HasPrefix(filepath.Base(path), ".") {
				return filepath.SkipDir
			}
			watchFolder(path)
		}

		return err
	})
}
