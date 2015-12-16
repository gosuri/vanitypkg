package main

import (
	"flag"

	"github.com/gosuri/vanitypkg"
)

func init() {
	flag.StringVar(&vanitypkg.Listen, "listen", vanitypkg.Listen, "Address to listen on")
	flag.StringVar(&vanitypkg.GitHubUser, "github-user", "", "GitHub User/Org name")
}

func main() {
	flag.Parse()
	vanitypkg.RunServer()
}
