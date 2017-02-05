package main

import (
	"bytes"
	"fmt"
	"io"
	"os/exec"

	flags "github.com/jessevdk/go-flags"
)

const (
	ExitCodeOK = iota
	ExitCodeParserFlagError
	ExitCodeUnsupportedError
)

type CLI struct {
	outStream, errStream io.Writer
	aws                  AWS
}

type Options struct {
	OptHelp    bool `short:"h" long:"help" description:"Show this help message and exit"`
	OptVersion bool `short:"v" long:"version" description:"Print the version and exit"`
}

func NewCLI(o io.Writer, e io.Writer) *CLI {
	return &CLI{outStream: o, errStream: e, aws: AWS{}}
}

func (cli *CLI) Run(args []string) int {
	opts, parsed, err := cli.parseOptions(args)
	if err != nil {
		return ExitCodeParserFlagError
	}

	if opts.OptHelp {
		cli.outStream.Write(cli.help())
		return ExitCodeOK
	}

	if opts.OptVersion {
		cli.outStream.Write(cli.version())
		return ExitCodeOK
	}

	if len(parsed) == 1 {
		s := parsed[0]
		if !cli.aws.Validate(s) {
			// TODO: display error message
			return ExitCodeUnsupportedError
		}
		cli.open(s)
	}

	return ExitCodeOK
}

func (cli *CLI) open(service string) {
	a := AWS{}
	u := a.Url(service)
	fmt.Println(u)
	// TODO: handling an error
	exec.Command("open", u).Run()
}

func (cli *CLI) parseOptions(args []string) (*Options, []string, error) {
	opts := &Options{}
	p := flags.NewParser(opts, flags.PrintErrors)
	args, err := p.ParseArgs(args)
	if err != nil {
		cli.errStream.Write(cli.help())
		return nil, nil, err
	}
	return opts, args, nil
}

func (CLI) version() []byte {
	buf := bytes.Buffer{}
	fmt.Fprintln(&buf, "silver-cornival version "+Version)
	return buf.Bytes()
}

func (CLI) help() []byte {
	buf := bytes.Buffer{}

	fmt.Fprintf(&buf, `
Usage: silver-carnival [options] AWS_SERVICE

AWS_SERVICE: ec2
`)
	return buf.Bytes()
}
