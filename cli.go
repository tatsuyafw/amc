package main

import (
	"bytes"
	"fmt"
	"io"
	"os/exec"
	"strings"

	flags "github.com/jessevdk/go-flags"
	"github.com/tatsuyafw/amc/aws"
)

const (
	exitCodeOK = iota
	exitCodeParserFlagError
	exitCodeArgumentError
)

type cli struct {
	outStream, errStream io.Writer
}

type options struct {
	OptHelp    bool `short:"h" long:"help" description:"Show this help message and exit"`
	OptVersion bool `short:"v" long:"version" description:"Print the version and exit"`
}

func newCli(o io.Writer, e io.Writer) *cli {
	return &cli{outStream: o, errStream: e}
}

func (c *cli) Run(args []string) int {
	opts, parsed, err := c.parseoptions(args)
	if err != nil {
		return exitCodeParserFlagError
	}

	if opts.OptHelp {
		c.outStream.Write(c.help())
		return exitCodeOK
	}

	if opts.OptVersion {
		c.outStream.Write(c.version())
		return exitCodeOK
	}

	if len(parsed) == 0 {
		c.showHelp()
		return exitCodeArgumentError
	}

	var service string
	var query string
	if len(parsed) > 0 {
		service = parsed[0]
	}
	if len(parsed) > 1 {
		query = parsed[1]
	}

	a, err := aws.New(service, query)
	if err != nil {
		c.showHelp() // TODO: show more details about an error.
		return exitCodeArgumentError
	}
	u := a.URL()
	c.open(u)

	return exitCodeOK
}

func (c *cli) open(url string) {
	fmt.Println(url)
	// TODO: handling an error
	exec.Command("open", url).Run()
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

func (c *cli) showHelp() {
	c.outStream.Write(c.help())
}

func (c *cli) help() []byte {
	buf := bytes.Buffer{}

	fmt.Fprintf(&buf, `
Usage: amc [options] AWS_SERVICE

AWS_SERVICE:
`)

	a := aws.Supported()
	s := strings.Join(a, ",")
	fmt.Fprintln(&buf, "  "+s)

	return buf.Bytes()
}
