FROM golang:1.24-alpine

WORKDIR /app

RUN apk add --no-cache git bash

COPY go.mod go.sum* ./
RUN go mod download

COPY . .

RUN go build -o /nexo ./cmd/api

EXPOSE 8080

CMD ["/nexo"]