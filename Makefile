run:
	go run ./cmd/main.go

build:
	go build -o bin/sqlite-operator ./cmd

check:
	staticcheck ./...

format:
	gofumpt -l -w .

imports:
	goimports -w .