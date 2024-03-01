
FROM golang:1.22.0-alpine3.19

WORKDIR /client

COPY . .

RUN go mod init pw-client
RUN go build -o client .

CMD ["./client"]