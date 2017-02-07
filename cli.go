package main

import (
	"bytes"
	"fmt"
	"io"
	"os/exec"
	"strings"

	flags "github.com/jessevdk/go-flags"
)

const (
	exitCodeOK = iota
	exitCodeParserFlagError
	exitCodeUnsupportedError
)

type cli struct {
	outStream, errStream io.Writer
	aws                  AWS
}

type options struct {
	optHelp    bool `short:"h" long:"help" description:"Show this help message and exit"`
	optVersion bool `short:"v" long:"version" description:"Print the version and exit"`
}

func newCli(o io.Writer, e io.Writer) *cli {
	return &cli{outStream: o, errStream: e, aws: AWS{}}
}

func (c *cli) Run(args []string) int {
	opts, parsed, err := c.parseoptions(args)
	if err != nil {
		return exitCodeParserFlagError
	}

	if opts.optHelp {
		c.outStream.Write(c.help())
		return exitCodeOK
	}

	if opts.optVersion {
		c.outStream.Write(c.version())
		return exitCodeOK
	}

	if len(parsed) == 1 {
		s := parsed[0]
		if !c.aws.Validate(s) {
			// TODO: display error message
			return exitCodeUnsupportedError
		}
		c.open(s)
	}

	return exitCodeOK
}

func (c *cli) open(service string) {
	a := AWS{}
	u := a.URL(service)
	fmt.Println(u)
	// TODO: handling an error
	exec.Command("open", u).Run()
}

func (c *cli) parseoptions(args []string) (*options, []string, error) {
	opts := &options{}
	p := flags.NewParser(opts, flags.PrintErrors)
	args, err := p.ParseArgs(args)
	if err != nil {
		c.errStream.Write(c.help())
		return nil, nil, err
	}
	return opts, args, nil
}

func (cli) version() []byte {
	buf := bytes.Buffer{}
	fmt.Fprintln(&buf, "amc version "+version)
	return buf.Bytes()
}

func (c *cli) help() []byte {
	buf := bytes.Buffer{}

	fmt.Fprintf(&buf, `
Usage: amc [options] AWS_SERVICE

AWS_SERVICE:
`)

	s := strings.Join(c.aws.supported(), ",")
	fmt.Fprintln(&buf, "  "+s)

	return buf.Bytes()
}
