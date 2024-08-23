FROM golang:1.23.0-alpine as builder

WORKDIR /code

COPY go.mod .
COPY go.sum .

RUN go mod tidy

COPY . .

RUN go build -tags jsoniter -o bin/mailer cmd/main.go

FROM alpine:3.20.2

WORKDIR /code

COPY --from=builder /code/bin/mailer .

ENTRYPOINT ["/code/mailer"]
