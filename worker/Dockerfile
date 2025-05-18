FROM golang:1.24-alpine

WORKDIR /app
COPY . .

RUN go mod init worker && go get github.com/nats-io/nats.go && go build -o app

CMD ["./app"]
