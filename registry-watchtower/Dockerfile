FROM busybox:1.34.1-musl AS runner
RUN adduser -D runner
USER runner
RUN mkdir -p /home/runner/static
COPY --chown=runner /src/* /home/runner/static/
CMD ["busybox", "httpd", "-f", "-p", "8080", "-h", "/home/runner/static"]
