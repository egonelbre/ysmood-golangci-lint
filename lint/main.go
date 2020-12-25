package main

import (
	"flag"
	"io/ioutil"
	"os"

	lint "github.com/ysmood/golangci-lint"
)

var ver = flag.String("golangci-lint-version", lint.DefaultVer, "version of the golangci-lint to use")
var quiet = flag.Bool("q", false, "disable the log not generated by golangci-lint")

func main() {
	args := getLintArgs()

	flag.Parse()

	ltr := lint.New()
	ltr.Version = *ver

	if *quiet {
		ltr.Logger.SetOutput(ioutil.Discard)
	}

	ltr.Lint(args...)
}

func getLintArgs() []string {
	args := []string{}
	lintArgs := []string{}
	sep := false
	for _, v := range os.Args {
		if v == "--" {
			sep = true
			continue
		}
		if sep {
			lintArgs = append(lintArgs, v)
		} else {
			args = append(args, v)
		}
	}

	os.Args = args

	return lintArgs
}
