# frontend

FROM node:16.18.1-alpine AS builder

COPY src /src

WORKDIR /src

RUN ["npm", "ci", "--omit=dev"]
RUN ["npm", "run", "build"]

FROM busybox:musl AS runner

RUN adduser -D runner

USER runner

WORKDIR /home/runner

RUN mkdir -p /home/runner/static

COPY --from=builder /src/public/* /home/runner/static/

CMD ["busybox", "httpd", "-f", "-p", "8080", "-h", "/home/runner/static"]
