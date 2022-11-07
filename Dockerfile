# syntax=docker/dockerfile:1

## Build
FROM golang:1.19.3-buster AS base

FROM base AS build

ARG ENV

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./
RUN go build -tags prod,railway -o /mygram

## Deploy
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /build/mygram /mygram

EXPOSE 8005

USER nonroot:nonroot

ENTRYPOINT ["/mygram", "app:serve"]
