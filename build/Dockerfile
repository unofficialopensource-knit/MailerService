FROM golang:1.22.1-alpine as builder

WORKDIR /code

COPY go.mod .
COPY go.sum .

RUN go mod tidy

COPY . .

RUN mkdir bin && go build -tags jsoniter -o bin/mailer cmd/main.go

FROM golang:1.22.1-alpine

WORKDIR /code

COPY --from=builder /code/bin/mailer .
