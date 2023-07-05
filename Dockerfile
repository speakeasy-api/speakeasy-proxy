# syntax=docker/dockerfile:1

FROM golang:1.20-alpine AS build-stage

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY *.go ./
COPY internal/ internal/
COPY pkg/ pkg/

RUN CGO_ENABLED=0 GOOS=linux go build -o /speakeasy-proxy

FROM alpine:latest

WORKDIR /

COPY --from=build-stage /speakeasy-proxy /speakeasy-proxy

ENTRYPOINT ["/speakeasy-proxy"]
