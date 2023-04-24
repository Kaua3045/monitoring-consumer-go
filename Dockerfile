FROM golang:1.20-alpine

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server

ENTRYPOINT [ "/app/server" ]