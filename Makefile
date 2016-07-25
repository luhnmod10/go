all: test bench

build:
	go vet
	go build

test:
	go test -cover

bench:
	go test -run=xxx -bench=. -benchmem
