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

	doSet bool

	ttl = time.Second * 3
)

var cli = flag.CommandLine

func init() {
	flag.BoolVar(&getVersion, "version", getVersion, "print current version")

	flag.BoolVar(&doSet, "set", doSet, "set a password or password template (using an empty password deletes the entry)")

	flag.DurationVar(&ttl, "ttl", ttl, "time until the copied password is erased from the clipboard again")

	orgUsage := flag.Usage
	flag.Usage = func() {
		print := func(format string, a ...interface{}) {
			fmt.Fprintf(cli.Output(), format, a...)
		}

		print("Temporarily copy password for easy pasting.\n\n")

		print("%s [-flags] [NAME]\n\n", os.Args[0])

		orgUsage()
		print("\n")

		print("Examples:\n")

		print("  %s\n", os.Args[0])
		print("\tread main password\n")

		print("  %s -set\n", os.Args[0])
		print("\tset password from prompt or stdin\n")

		print("  %s NAME\n", os.Args[0])
		print("\tread password \"NAME\"\n")

		print("  %s -set NAME\n", os.Args[0])
		print("\tset password \"NAME\" from prompt or stdin\n")
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
	if flag.NArg() > 1 {
		return errors.New("expected zero or one arguments")
	}
	name := flag.Arg(0)

	if doSet {
		pw, err := clio.ReadPw(os.Stdin, cli.Output())
		if err != nil {
			return err
		}
		return pasteword.Write(name, pw)
	}

	pw, err := pasteword.Read(name)
	if err != nil {
		return err
	}

	return pasteword.TempCopy(pw, ttl)
}
