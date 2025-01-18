package main

import (
	"file-map-server/app"
	"file-map-server/log"
)

func main() {
	go log.LogRotation()
	log.LogInit()
	app.Run()
}
