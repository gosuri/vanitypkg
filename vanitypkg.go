// Package vanitypkg provides a server that enables vanity package names for golang libraries on github
package vanitypkg

import (
	"log"
	"net/http"
	"net/url"
	"path"
	"strings"
	"text/template"
)

// Listen is the default listen address
var Listen = ":8080"

// GitHubUser is the user or org id to redirect requests to
var GitHubUser string

// RunServer starts an http server that server packages
func RunServer() {
	http.HandleFunc("/", handler)
	log.Println("vanitypkg: starting server on", Listen)
	if err := http.ListenAndServe(Listen, nil); err != nil {
		log.Fatal(err)
	}
}

// handler is the http handler that renders response
func handler(w http.ResponseWriter, req *http.Request) {
	p := &Package{
		Host:       req.Host,
		Path:       req.URL.Path,
		GitHubUser: GitHubUser,
	}
	logReq(req, p)
	tmpl := template.Must(template.New("html").Parse(htmlT))
	tmpl.Execute(w, p)
}

// Package represents a golang package
type Package struct {
	Host       string
	Path       string
	GitHubUser string
}

// RootPath returns the root path for a package
func (p *Package) RootPath() string {
	return path.Join(p.Host, p.rootpkg())
}

// GitHubRepo returns the github repo for a package
func (p *Package) GitHubRepo() string {
	return "https://" + path.Join("github.com", p.GitHubUser, p.rootpkg())
}

// PackagePath returns the the full path to the package
func (p *Package) PackagePath() string {
	return path.Join(p.Host, p.Path)
}

func (p *Package) rootpkg() string {
	// get the root package, clean any leading slashes
	root := strings.TrimPrefix(p.Path, "/")
	root = strings.Split(root, "/")[0]
	u, err := url.Parse(root)
	if err != nil {
		return ""
	}
	return u.Path
}

func logReq(r *http.Request, p *Package) {
	log.Printf("path=%q host=%q pkg=%q github=%s", r.URL.Path, r.Host, p.PackagePath(), p.GitHubRepo())
}

const htmlT = `<!DOCTYPE html>
<html>
<head>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8"/>
<meta name="go-import" content="{{ .RootPath }} git {{ .GitHubRepo }}">
<meta name="go-source" content="{{ .RootPath }} {{ .GitHubRepo }} {{ .GitHubRepo }}/tree/master{/dir} {{ .GitHubRepo }}/blob/master{/dir}/{file}#L{line}">
<meta http-equiv="refresh" content="0; url=https://godoc.org/{{ .PackagePath }}">
</head>
<body>
Nothing to see here; <a href="https://godoc.org/{{ .PackagePath }}">move along</a>.
</body>
</html>`
