install:
	go mod tidy

build-app:
	go build -tags jsoniter -o bin/mailer cmd/main.go

clean:
	rm bin/mailer

run-dev:
	dotenv --dotenv dev.env go run cmd/main.go

format:
	go fmt ./...

test-unit:
	dotenv --dotenv test.env go test -v --covermode atomic --coverpkg github.com/unofficialopensource-knit/MailerService/internal/app ./internal/app_test/unit/

test-integration:
	dotenv --dotenv test.env go test -v --covermode atomic --coverpkg github.com/unofficialopensource-knit/MailerService/internal/app/ ./internal/app_test/integration/
