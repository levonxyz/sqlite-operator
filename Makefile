run:
	go run ./cmd/main.go

build:
	CGO_ENABLED=0 go build -o bin/sqlite-operator ./cmd

check:
	staticcheck ./...

format:
	gofumpt -l -w .

imports:
	goimports -w .