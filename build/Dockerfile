FROM golang:1.25.3-alpine as builder

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
