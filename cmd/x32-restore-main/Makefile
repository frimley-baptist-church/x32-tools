
app=$(shell basename $$PWD)

fmt:
	go fmt ./...

build: fmt
	go build -ldflags "-w -s" .

test: restore

restore:
	./$(app) --file ../snapshot.json x32
