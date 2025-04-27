
app=$(shell basename $$PWD)

fmt:
	go fmt ./...

build: fmt
	go build -ldflags "-w -s" .

test: store

store:
	./$(app) --store ../snapshot.json x32
	cat ../snapshot.json
