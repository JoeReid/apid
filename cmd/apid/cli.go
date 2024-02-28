package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"strings"

	"github.com/JoeReid/apid"
	"github.com/google/uuid"
)

type CLI struct {
	Stdin io.Reader
	Args  []string
}

func (c *CLI) Run() (string, error) {
	var (
		prefix string
		decode bool
	)

	fs := flag.NewFlagSet("", flag.ContinueOnError)
	fs.StringVar(&prefix, "p", "unknown", "Prefix for the APID")
	fs.BoolVar(&decode, "d", false, "Decode rather than encode")

	if err := fs.Parse(c.Args); err != nil {
		// This is a hack to get the help message to print without the flag
		// package printing it's own error, but preserving the exit code 1
		if errors.Is(err, flag.ErrHelp) {
			return "", errors.New("")
		}

		return "", fmt.Errorf("error parsing flags: %w", err)
	}

	// Get the data either from a positional argument or from stdin
	var data string
	if fs.NArg() > 0 {
		data = fs.Arg(0)
	} else {
		out, err := io.ReadAll(c.Stdin)
		if err != nil {
			return "", fmt.Errorf("error reading data: %w", err)
		}

		data = strings.TrimSpace(string(out))
	}

	if decode {
		return c.decode(data)
	}

	return c.encode(prefix, data)
}

func (c *CLI) decode(data string) (string, error) {
	a, err := apid.Parse(data)
	if err != nil {
		return "", fmt.Errorf("error parsing APID %q: %w", data, err)
	}

	return a.UUID.String(), nil
}

func (c *CLI) encode(prefix, data string) (string, error) {
	u, err := uuid.Parse(data)
	if err != nil {
		return "", fmt.Errorf("error parsing UUID %q: %w", data, err)
	}

	return (&apid.APID{
		Prefix: prefix,
		UUID:   u,
	}).String(), nil
}
