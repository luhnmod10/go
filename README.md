# Luhn (Mod 10)

[![Build Status](https://img.shields.io/travis/luhnmod10/go.svg?label=linux%20%26%20osx)](https://travis-ci.org/luhnmod10/go)
[![Codecov](https://img.shields.io/codecov/c/github/luhnmod10/go.svg)](https://codecov.io/gh/luhnmod10/go)
[![Go Report Card](https://goreportcard.com/badge/github.com/luhnmod10/go)](https://goreportcard.com/report/github.com/luhnmod10/go)
[![Go docs](https://img.shields.io/badge/godoc-reference-blue.svg)](https://godoc.org/github.com/luhnmod10/go)

A fast and simple in-place implementation of the [luhn check algorithm](https://en.wikipedia.org/wiki/Luhn_algorithm) in Go. Implementations in other languages can be found at [github.com/luhn-algorithm](https://github.com/luhnmod10).

## Usage

```
go get github.com/luhnmod10/go
```

```go
_ = luhn.Valid("4242424242424242")
```

## Contributing

Contributions are welcome! If you can improve the execution time of this implementation without increasing its complexity, please open a pull request. To test your change, run `make` in the repository to run the tests and the benchmarks.
