package main

import (
	"os"
)

func main() {

	args := os.Args[1:]

	if args[0] == "server" {
		server(args[1], args[2])
	} else {
		client(args[1])
	}
}
