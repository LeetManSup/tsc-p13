FROM golang:1.24-alpine

WORKDIR /app
COPY . .

RUN go mod init publisher && go get github.com/nats-io/nats.go && go build -o app

CMD ["./app"]
