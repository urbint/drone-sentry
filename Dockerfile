FROM gliderlabs/alpine:3.1
RUN apk-install ca-certificates
ADD drone-sentry /bin/
ENTRYPOINT ["/bin/drone-sentry"]
