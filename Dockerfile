# syntax=docker/dockerfile:1

FROM golang:1.17-alpine as builder
RUN apk add bash && apk add postgresql-client


WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . /app
RUN chmod +x wait-for-postgres.sh

RUN go build -o auth-service

EXPOSE 8080
CMD ["./auth-service"]