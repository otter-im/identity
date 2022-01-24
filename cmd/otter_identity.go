package main

import (
	"context"
	"flag"
	"github.com/otter-im/identity/internal/app"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	flag.Parse()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		app.Init()

		err := app.Run()
		if err != nil {
			cancel()
		}

		err = app.Exit(ctx)
		if err != nil {
			cancel()
		}
	}()

	_ = <-sig
	cancel()
}
