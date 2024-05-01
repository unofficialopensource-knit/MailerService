install:
	go mod tidy

build-app:
	go build -tags jsoniter -o bin/mailer cmd/main.go

clean:
	rm bin/mailer

run-dev:
	dotenv --dotenv dev.env go run cmd/main.go

format:
	go fmt ./cmd/ ./pkg/**

tests:
	dotenv --dotenv test.env go test -v --covermode atomic ./pkg/**
