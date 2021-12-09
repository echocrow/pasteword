package cmd

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/echocrow/pasteword/pkg/clio"
	"github.com/echocrow/pasteword/pkg/pasteword"
)

var (
	getVersion bool

	pwName      = ""
	placeholder = "{}"
	ttl         = time.Second * 3

	doSet bool
)

var cli = flag.CommandLine

func init() {
	flag.BoolVar(&getVersion, "version", getVersion, "print current version")

	flag.StringVar(&pwName, "k", pwName, "shorthand for -key")
	flag.StringVar(&pwName, "key", pwName, "name of the key")

	flag.StringVar(&placeholder, "placeholder", placeholder, "password template placeholder for dynamic variables")

	flag.DurationVar(&ttl, "ttl", ttl, "time until the copied password is erased")

	flag.BoolVar(&doSet, "set", doSet, "set a password or password template (using an empty password deletes the entry)")

	flag.Usage = func() {
		fmt.Fprintf(cli.Output(), "Paste basic or template-based keys and passwords .\n\n")

		fmt.Fprintf(cli.Output(), "Usage of %s:\n\n", os.Args[0])

		fmt.Fprintf(cli.Output(), "%s\n", os.Args[0])
		fmt.Fprintf(cli.Output(), "\tRead password\n\n")

		fmt.Fprintf(cli.Output(), "%s [VAR, ...]\n", os.Args[0])
		fmt.Fprintf(cli.Output(), "\tRead password and replace placeholders with VARs\n\n")

		fmt.Fprintf(cli.Output(), "%s -set\n", os.Args[0])
		fmt.Fprintf(cli.Output(), "\tSet password from prompt or stdin\n\n")

		flag.PrintDefaults()
	}
}

// Execute executes the root command.
func Execute(
	version string,
) {
	flag.Parse()

	if getVersion {
		fmt.Fprintf(cli.Output(), "%s %s", os.Args[0], version)
		os.Exit(0)
	}

	if err := run(); err != nil {
		fmt.Fprintf(cli.Output(), "%s\n", err)
		os.Exit(1)
	}
}

func run() error {
	if doSet {
		if flag.NArg() != 0 {
			return errors.New("expected no arguments when setting template")
		}

		tpl, err := clio.ReadPw(os.Stdin, cli.Output())
		if err != nil {
			return err
		}

		return pasteword.Write(pwName, tpl)
	}

	pw, err := pasteword.Read(pwName, placeholder, flag.Args()...)
	if err != nil {
		return err
	}

	return pasteword.TempCopy(pw, ttl)
}
