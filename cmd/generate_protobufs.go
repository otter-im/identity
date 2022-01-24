package main

import (
	"flag"
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"syscall"
)

func main() {
	dirname := flag.String("dir", ".", "Directory to traverse")
	flag.Parse()

	if err := filepath.Walk(*dirname, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if filepath.Ext(path) == ".proto" {
			if err = compileProtobuf(path); err != nil {
				fmt.Printf("compilation failure for \"%s\":\n%v\n", path, err)
			}
		}
		return nil
	}); err != nil {
		fmt.Print(err)
	}
}

func compileProtobuf(path string) error {
	cmd := fmt.Sprintf("protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative %s", path)
	fmt.Printf(cmd)

	cmdSplit := strings.Split(cmd, " ")

	binary, err := exec.LookPath(cmdSplit[0])
	if err != nil {
		return err
	}

	err = syscall.Exec(binary, cmdSplit, os.Environ())
	if err != nil {
		return err
	}
	return nil
}
