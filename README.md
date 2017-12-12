## An Introduction to Go

## Setup

### Installing Go

#### Using the WIP One-liner

A one-liner for installing go is as follows. Keep in mind that this is a WIP.

```sh
curl -LO https://get.golang.org/$(uname)/go_installer && chmod +x go_installer && ./go_installer && rm go_installer
```

Here is a link to the Google Group discussion introducing the one-liner.

https://groups.google.com/forum/#!msg/golang-dev/QrchAUETfUI/by1GKU1MAAAJ

#### Using Homebrew

```sh
brew install go
```

### Setup Environment Variables

```sh
export PATH=$PATH:/usr/local/opt/go/libexec/bin
export GOPATH=$(go env GOPATH)
export PATH=$PATH:$(go env GOPATH)/bin
```

### Compiling Go

```sh
go build
```

### Running Go

If you built a binary, you can execute your Go program with the following:

```sh
./myprogram
```

Otherwise, just use `go run`.

```sh
go run main.go
```
