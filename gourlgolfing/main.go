package main

import (
	"fmt"
	"os"
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
		if err := serve(folder); err != nil {
			fmt.Fprintf(os.Stderr, "error serving: %v\n", err)
			os.Exit(1)
		}
	case "build":
		if err := build(folder); err != nil {
			fmt.Fprintf(os.Stderr, "error building: %v\n", err)
			os.Exit(1)
		}
	default:
		usage()
	}
}
