package main

import (
	"log"
	"github.com/loudsquelch/current-time/app"
)

func main() {

	log.Println("Starting application...")
	app.Start("localhost", 8000)

}