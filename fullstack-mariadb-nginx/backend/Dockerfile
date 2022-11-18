# build files with seperate container, prevents having source code and other useless files in final container
# alpine is a slimmer image than debian, and since we won't need anything else it's fine
FROM golang:1.19.2-alpine AS builder

COPY src /src
WORKDIR /src

RUN ["go", "build", "."]

# for whatever reason, it's unable to find the executable if it uses busybox or scratch as the base image, despite it appearing with ls
FROM alpine:3.14.8 AS runner

# create seperate user for running process as non-root
RUN adduser -D runner
USER runner
WORKDIR /home/runner

# copy from 'builder', chown to make runner the owner of the built binary
COPY --chown=runner --from=builder /src/example_backend /home/runner

# pass environment variables to current container
ENV DB_USER=$DB_USER
ENV DB_PASSWORD=$DB_PASSWORD
ENV DB_DATABASE=$DB_DATABASE
ENV DB_HOST=$DB_HOST

CMD ["./example_backend"]
