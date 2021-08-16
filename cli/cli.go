package cli

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/url"
	"os"

	"github.com/hskwakr/findlink/crawler"
)

const (
	AppName = "findlink"

	ExitCodeOK               = 0
	ExitCodeParseFlagError   = 1
	ExitCodeArgumentsError   = 1
	ExitCodeApplicationError = 1
)

type CLI struct {
	OutStream, ErrStream io.Writer
	args
}

type args struct {
	// arguments
	url string

	// options
	o *string
	d *string
}

func (c *CLI) Run(args []string) int {
	if r := c.parse(args); r != ExitCodeOK {
		return r
	}

	links, err := crawler.GetLinks(c.url, *c.d, c.OutStream)
	if err != nil {
		log.Println(err)

		return ExitCodeApplicationError
	}

	if *c.o != "" {
		writeJSON(links, *c.o)
	} else {
		writeOut(c.OutStream, links)
	}

	return ExitCodeOK
}

func (c *CLI) parse(args []string) int {
	flags := flag.NewFlagSet(AppName, flag.ContinueOnError)
	flags.SetOutput(c.ErrStream)

	flags.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage:\n\t"+AppName+" [option] URL\n\n")
		flags.PrintDefaults()
		os.Exit(0)
	}

	// options
	c.o = flags.String("o", "", "The path to the json file for output.")
	c.d = flags.String("d", "", "Filter the output by domain.")

	if err := flags.Parse(args[1:]); err != nil {
		return ExitCodeParseFlagError
	}

	// Two arguments are required
	if len(flags.Args()) < 1 {
		return ExitCodeArgumentsError
	}

	c.url = flags.Arg(0)
	if !urlValidation(c.url) {
		return ExitCodeArgumentsError
	}

	return ExitCodeOK
}

func urlValidation(raw string) bool {
	u, err := url.ParseRequestURI(raw)
	if err != nil {
		log.Println(err)

		return false
	}

	if u.Scheme != "http" && u.Scheme != "https" {
		log.Println("Wrong scheme: should be http or https")

		return false
	}

	return true
}

func writeOut(o io.Writer, data []crawler.Link) {
	if _, err := o.Write([]byte("\n")); err != nil {
		log.Fatal(err)
	}

	for _, v := range data {
		// fmt.Println(v.URL)
		if _, err := o.Write([]byte(v.URL + "\n")); err != nil {
			log.Fatal(err)
		}
	}
}

func writeJSON(data []crawler.Link, path string) {
	var b bytes.Buffer

	enc := json.NewEncoder(&b)
	enc.SetEscapeHTML(false)
	enc.SetIndent("", "  ")

	if err := enc.Encode(data); err != nil {
		log.Fatal(err)
	}

	const mode = 0644
	_ = ioutil.WriteFile(path, b.Bytes(), os.FileMode(mode))
}
