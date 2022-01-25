package main

import (
	"flag"
	"github.com/otter-im/identity/internal/app"
	"log"
	"os"
)

func main() {
	flag.Parse()

	app.Init()
	err := app.Run()
	if err != nil {
		log.Print(err)
		os.Exit(-1)
	}

	err = app.Exit()
	if err != nil {
		log.Print(err)
		os.Exit(-1)
	}
}
