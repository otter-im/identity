package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/go-pg/migrations/v8"
	"github.com/golang/glog"
	"github.com/otter-im/identity-service/internal/app"
	"os"
)

const usageText = `This program runs command on the db. Supported commands are:
  - init - creates version info table in the database
  - up - runs all available migrations.
  - up [target] - runs available migrations up to the target one.
  - down - reverts last migration.
  - reset - reverts all migrations.
  - version - prints current db version.
  - set_version [version] - sets db version without running migrations.
Usage:
  go run *.go <command> [args]
`

func main() {
	flag.Usage = usage
	flag.Parse()

	app.Init()
	defer func() {
		if err := app.Exit(context.Background()); err != nil {
			glog.Error(err)
		}
	}()

	oldVersion, newVersion, err := migrations.Run(app.Postgres(), flag.Args()...)
	if err != nil {
		glog.Exitf("migration %d -> %d failed: %s\n", oldVersion, newVersion, err)
	}

	if newVersion != oldVersion {
		fmt.Printf("migrated from %d to %d\n", oldVersion, newVersion)
	} else {
		fmt.Printf("version is %d\n", newVersion)
	}
}

func usage() {
	fmt.Print(usageText)
	flag.PrintDefaults()
	os.Exit(2)
}
