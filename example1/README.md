# example1

## Obtaining the Dependencies

This example uses [dep](https://github.com/golang/dep), which is a prototype
dependency management tool for Go.

To install dep, follow the instructions in dep's [documention](https://golang.github.io/dep/docs/installation.html).

Installing the correct versions of the depencies used in example1 can be done through
the use of dep's `ensure` command. More documentation on the `ensure` command can
be found [here](https://golang.github.io/dep/docs/daily-dep.html).

For this example, run the following to catch up on installing the dependencies.

```
$ dep ensure -v
```

## Testing example1

To testing example1, you will leverage the toolchain's `test` command, specifically with the following:

```
$ go test -v .
```

This should produce the following output:

```
=== RUN   TestAdd
--- PASS: TestAdd (0.00s)
=== RUN   TestDivide
--- PASS: TestDivide (0.00s)
PASS
ok  	github.com/mccurdyc/go-examples/example1	0.008s
```

## Build Static Binary

To build the single, static, binary for example1, run the following in `example1/`:

```
$ go build -o bin/example1
```

This should create an executable binary in the local `bin` directory.

## Running the Executable Binary

To run the built binary, use the following:

```
$ ./bin/example1
```

This should produce the following output:

```
sum: 5
quotient1: 4
quotient2: error dividing by zero
```
