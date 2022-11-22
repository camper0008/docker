# build files with seperate container, prevents having source code and other useless files in final container
# alpine is a slimmer image than debian, and since we won't need anything else it's fine
FROM golang:1.19.2-alpine AS builder

COPY src /src
WORKDIR /src

# disable cgo, which means it's able to statically link required libraries, which allows us to run it on scratch
ENV CGO_ENABLED=0

RUN ["go", "build", "."]

FROM scratch AS runner

# chown to make runner the owner of the built binary
COPY --from=builder /src/example_backend /

# pass environment variables to current container
# this should be done last, since it allows docker to cache more efficiently
ENV DB_USER=$DB_USER
ENV DB_PASSWORD=$DB_PASSWORD
ENV DB_DATABASE=$DB_DATABASE
ENV DB_HOST=$DB_HOST

CMD ["/example_backend"]
