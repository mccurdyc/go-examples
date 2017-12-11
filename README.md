## An Introduction to Go

## Setup

### Installing Go

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
