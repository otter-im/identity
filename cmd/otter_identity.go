package main

import (
	"flag"
	"github.com/otter-im/identity/internal/app"
	"log"
	"os"
)

func main() {
	flag.Parse()

	err := app.Init()
	if err != nil {
		log.Print(err)
		os.Exit(-1)
	}

	err = app.Run()
	if err != nil {
		log.Print(err)
		os.Exit(-2)
	}

	err = app.Exit()
	if err != nil {
		log.Print(err)
		os.Exit(-3)
	}
}
