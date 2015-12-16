package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/gosuri/vanitypkg"
)

var gbrepos string

func init() {
	flag.StringVar(&vanitypkg.Listen, "listen", vanitypkg.Listen, "Address to listen on")
	flag.StringVar(&vanitypkg.GitHubUser, "github-user", "", "GitHub User/Org name")
	flag.StringVar(&gbrepos, "gb-repos", "", "Comma seperated list of github repos that use gb")
}

func main() {
	flag.Parse()

	if len(gbrepos) > 0 {
		for _, rep := range strings.Split(gbrepos, ",") {
			vanitypkg.GBRepos = append(vanitypkg.GBRepos, strings.TrimSpace(rep))
		}
	}

	if err := vanitypkg.RunServer(); err != nil {
		fmt.Fprintln(os.Stderr, "FATAL:", err)
		os.Exit(1)
	}
	os.Exit(0)
}
