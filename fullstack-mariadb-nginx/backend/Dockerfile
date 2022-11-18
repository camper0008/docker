FROM golang:1.19.2-alpine AS builder

COPY src /src

WORKDIR /src

ENV GOOS=linux

ENV GOARCH=amd64

RUN ["go", "build", "."]

# for whatever reason, it's unable to find the executable if it uses busybox or scratch as the base image, despite it appearing with ls

FROM alpine:3.14.8 AS runner

RUN adduser -D runner

USER runner

WORKDIR /home/runner

COPY --chown=runner --from=builder /src/example_backend /home/runner

ENV DB_USER=$DB_USER
ENV DB_PASSWORD=$DB_PASSWORD
ENV DB_DATABASE=$DB_DATABASE
ENV DB_HOST=$DB_HOST

CMD ["./example_backend"]
