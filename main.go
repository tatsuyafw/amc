package main

import "os"

const Version string = "0.0.1"

func main() {
	// cli := &CLI{outStream: os.Stdout, errStream: os.Stderr}
	cli := NewCLI(os.Stdout, os.Stderr)
	os.Exit(cli.Run(os.Args[1:]))
}
