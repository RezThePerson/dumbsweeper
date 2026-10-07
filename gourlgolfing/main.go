package main

import (
	"fmt"
	"os"

	"github.com/reztheperson/dumbsweeper/gourlgolfing/packages/dev"
)

func usage() {
	fmt.Fprintln(os.Stderr, "Usage: gourlgolfing <dev|build> folder")
	os.Exit(1)
}

func main() {
	if len(os.Args) < 3 {
		usage()
	}

	command := os.Args[1]
	folder := os.Args[2]

	switch command {
	case "dev":
		dev.Serve(folder)
	case "build":
		// runBuild(folder)
	default:
		usage()
	}
}
