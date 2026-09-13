FROM golang:1.27.1 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download