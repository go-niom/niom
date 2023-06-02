package watcher

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/go-niom/niom/pkg/constants"
	"github.com/go-niom/niom/pkg/logger"
	"github.com/go-niom/niom/pkg/terminal"
	"github.com/go-niom/niom/pkg/utils"
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

func watchFolder(path string, ext []string) {
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

					if ext[0] != "*" {
						for _, v := range ext {
							if strings.HasSuffix(event.Name, v) {
								WatcherChannel <- event.String()
							}
						}
					} else {
						WatcherChannel <- event.String()
					}

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
func Watch(args []string) {
	ext := []string{"*"}
	if extensions := utils.ReadArgs("-e=", args); extensions != "" {
		ext = strings.Split(extensions, ",")
	}
	logger.Info("watching path(s): *.*")
	logger.Info(fmt.Sprintf("watching extensions: %v", ext))
	rebuildCheck()
	root := "."
	filepath.Walk(root, func(path string, info os.FileInfo, err error) error {

		if info.IsDir() {
			if len(path) > 1 && strings.HasPrefix(filepath.Base(path), ".") {
				println(path)
				return filepath.SkipDir
			}
			watchFolder(path, ext)
		}

		return err
	})
}
