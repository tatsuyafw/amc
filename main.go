package main

import "os"

const version string = "0.0.1"

func main() {
	cli := newCli(os.Stdout, os.Stderr)
	os.Exit(cli.Run(os.Args[1:]))
}
