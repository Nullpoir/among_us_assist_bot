FROM golang:1.16

COPY ./app /app
WORKDIR /app

RUN go mod download

CMD go run application.go
