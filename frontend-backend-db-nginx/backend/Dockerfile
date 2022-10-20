# backend

FROM golang:1.19.2

ENV PORT=$PORT

COPY src /docker/backend

WORKDIR /docker/backend

RUN go build .
