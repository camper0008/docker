# statically build files with builder container
# prevents having node modules and other useless files in final running container
FROM node:16.18.1-alpine AS builder

COPY src /src
WORKDIR /src

RUN ["npm", "ci", "--omit=dev"]
RUN ["npm", "run", "build"]

# serve statically built files with busybox' httpd
FROM busybox:1.34.1-musl AS runner

# run as seperate user to prevent root access
RUN adduser -D runner
USER runner
RUN mkdir -p /home/runner/static
# chown to make runner the owner of the static files rather than root
COPY --chown=runner --from=builder /src/public/* /home/runner/static/

CMD ["busybox", "httpd", "-f", "-p", "8080", "-h", "/home/runner/static"]
