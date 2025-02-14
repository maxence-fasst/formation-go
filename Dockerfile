# Builder image
FROM golang:1.23-alpine AS builder

ARG GITHUB_USER=${GITHUB_USER}
ARG GITHUB_PASS=${GITHUB_PASS}

## Install git in order to download private dependencies from github.com/fasst and build-base (gcc) for swag init
RUN apk add git build-base

RUN go install github.com/swaggo/swag/cmd/swag@v1.8.12

WORKDIR /app

## Install dependencies
COPY go.mod ./
COPY go.sum ./


RUN echo "machine github.com login ${GITHUB_USER} password ${GITHUB_PASS}" >> ~/.netrc && \
    go env -w GOPRIVATE=github.com/fasst/* && \
    go mod download && \
    rm ~/.netrc

## Build
COPY . ./

RUN swag init
RUN go build -o /go-project

# Runner image
FROM alpine:3.18

WORKDIR /

COPY --from=builder /go-project ./go-project

CMD ["/go-project"]