FROM golang:1.22.6-alpine as builder

WORKDIR /code

COPY go.mod .
COPY go.sum .

RUN go mod tidy

COPY . .

RUN go build -tags jsoniter -o bin/mailer cmd/main.go

FROM golang:1.22.6-alpine

WORKDIR /code

COPY --from=builder /code/bin/mailer .

ENTRYPOINT ["/code/mailer"]
