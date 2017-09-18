FROM alpine:latest AS cafile
RUN apk add --no-cache ca-certificates


FROM busybox:latest
COPY --from=cafile /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

ADD crontabs /var/spool/cron/crontabs
ADD override_bin/sendmail /bin/sendmail

ENTRYPOINT ["/bin/crond", "-f", "-d", "4", "-L", "/dev/stderr"]
