# Luhn (Mod 10)

[![Go docs](https://img.shields.io/badge/godoc-reference-blue.svg)](https://godoc.org/github.com/luhnmod10/go)

A fast and simple in-place implementation of the [luhn check algorithm](https://en.wikipedia.org/wiki/Luhn_algorithm) in Go. Implementations in other languages can be found at [github.com/luhnmod10](https://github.com/luhnmod10).

## Usage

```
go get github.com/luhnmod10/go
```

```go
_ = luhnmod10.Valid("4242424242424242")
```

## Contributing

Contributions are welcome! If you can improve the execution time of this implementation without increasing its complexity, please open a pull request. To test your change, run `make` in the repository to run the tests and the benchmarks.
