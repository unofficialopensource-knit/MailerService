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
	go fmt -n ./cmd/**.go
	go fmt -n ./pkg/config/**.go
	go fmt -n ./pkg/handler/**.go
	go fmt -n ./pkg/schema/**.go
	go fmt -n ./pkg/service/**.go

tests:
	go test -v --covermode atomic ./pkg/**
