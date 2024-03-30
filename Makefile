include .envs/.local
export

install:
	go mod tidy

build-app:
	go build -tags jsoniter -o bin/mailer cmd/main.go

clean:
	rm bin/mailer

run-dev:
	go run cmd/main.go

format:
	go fmt ./cmd/ ./pkg/**

tests:
	go test -v --covermode atomic ./pkg/**
