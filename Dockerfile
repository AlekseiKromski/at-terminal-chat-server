# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

WORKDIR /app

COPY config.json ./

COPY go.mod go.sum ./
RUN go mod download

COPY ./ ./

RUN go build .

CMD ./at-terminal-chat-server
