package vanitypkg

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGithubRepo(t *testing.T) {
	pkg := &Package{GitHubUser: "gosuri", Path: "/uilive/foo/...?go-get=1"}
	got := pkg.GitHubRepo()
	want := "github.com/gosuri/uilive"
	if got != want {
		t.Fatal("want", want, "got", got)
	}
}

func TestSourcePrefix(t *testing.T) {
	GBRepos = []string{"github.com/user/repo"}
	pkg := &Package{
		Host:       "example.com",
		GitHubUser: "user",
		Path:       "/repo/inner",
	}
	got := pkg.SourcePrefix()
	want := "master/src"
	if got != want {
		t.Fatal("want", want, "got", got)
	}
}

func TestHttp(t *testing.T) {
	req, err := http.NewRequest("GET", "http://example.com/foo/bar/...?go-get=1", nil)
	if err != nil {
		t.Fatal(err)
	}
	w := httptest.NewRecorder()
	GitHubUser = "user"
	handler(w, req)

	got := strings.TrimSpace(w.Body.String())
	want := strings.TrimSpace(testResp)

	if got != want {
		t.Fatal("\nwant --\n", want, "\ngot --\n", got)
	}
}

const testResp = `<!DOCTYPE html>
<html>
<head>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8"/>
<meta name="go-import" content="example.com/foo git https://github.com/user/foo">
<meta name="go-source" content="example.com/foo https://github.com/user/foo https://github.com/user/foo/tree/master{/dir} https://github.com/user/foo/blob/master{/dir}/{file}#L{line}">
<meta http-equiv="refresh" content="0; url=https://godoc.org/example.com/foo/bar/...">
</head>
<body>
Nothing to see here; <a href="https://godoc.org/example.com/foo/bar/...">move along</a>.
</body>
</html>`
