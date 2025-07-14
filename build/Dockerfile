FROM golang:1.24.5-alpine as builder

WORKDIR /code

COPY go.mod .
COPY go.sum .

RUN go mod tidy

COPY . .

RUN go build -tags jsoniter -o bin/mailer cmd/main.go

FROM alpine:3.22.0

WORKDIR /code

COPY --from=builder /code/bin/mailer .

ENTRYPOINT ["/code/mailer"]
