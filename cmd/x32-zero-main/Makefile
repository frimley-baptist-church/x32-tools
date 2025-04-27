
app=$(shell basename $$PWD)

fmt:
	go fmt ./...

build: fmt
	go build -ldflags "-w -s" .

test: zero

zero:
	./$(app) x32
