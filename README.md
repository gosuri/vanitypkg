# vanitypkg [![GoDoc](https://godoc.org/github.com/gosuri/vanitypkg?status.svg)](https://godoc.org/github.com/gosuri/vanitypkg) [![Build Status](https://travis-ci.org/gosuri/vanitypkg.svg?branch=master)](https://travis-ci.org/gosuri/vanitypkg)

vanitypkg provides a server that hosts vanity package names for Go package hosted on github

## Usage

```sh
$ vanitypkg -github-user=<user> -port=<port>
```

### Example

```sh
# start on port 80 using sudo
$ sudo vanitypkg -github-user=gosuri -port=80

# make go.gregosuri.com resolvable, this should be dns entry
$ sudo echo "127.0.0.1 go.gregosuri.com" > /etc/hosts 

$ go get -insecure go.gregosuri.com/uiprogress
...
$ ls $GOPATH/src/go.gregosuri.com
uiprogress
```

## Installation

```
$ go get github.com/gosuri/vanitypkg/cmd
```
