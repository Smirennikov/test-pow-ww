FROM golang:1.22.0-alpine3.19

WORKDIR /server

COPY . .

RUN go mod init pw-server
RUN go build -o server .

EXPOSE 8080

CMD ["./server"]