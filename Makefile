run:
	go run ./cmd/main.go

check:
	staticcheck ./...

format:
	gofumpt -l -w .

imports:
	goimports -w .