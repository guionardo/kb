---
title: Detecting filesystem events
tags:
    - golang
    - notification
    - "file system"
---

!!! warning

	Needs testing/refactoring! Last time, this failed to work!

```golang
package main

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
)

// main
func main() {

	// creates a new file watcher
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		fmt.Println("ERROR", err)
	}
	defer watcher.Close()

	//
	done := make(chan bool)

	//
	go func() {
		for {
			select {
			// watch for events
			case event := <-watcher.Events:
				fmt.Printf("EVENT! %#v\n", event)

				// watch for errors
			case err := <-watcher.Errors:
				fmt.Println("ERROR", err)
			}
		}
	}()

	// out of the box fsnotify can watch a single file, or a single directory
	if err := watcher.Add("/Users/skdomino/Desktop/test.html"); err != nil {
		fmt.Println("ERROR", err)
	}

	<-done
}
```