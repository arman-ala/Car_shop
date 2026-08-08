# Declare targets that aren't actual files on disk
.PHONY: run build exec

run:
	go run ./cmd/main.go

build:
	go build -o ./cmd/cmd.exe ./cmd/main.go

exec: build
	./cmd/cmd.exe
