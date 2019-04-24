# --------------------
# Build
# --------------------
FROM golang:1.12-stretch as build
WORKDIR $GOPATH/src/github.com/cachecashproject/go-cachecash
COPY . .
RUN make PREFIX=/artifacts all

FROM docker.elastic.co/beats/filebeat:6.6.2 as filebeat

# --------------------
# Package
# --------------------

FROM debian:stretch

RUN apt-get update \
	&& apt-get install -y --no-install-recommends logrotate cron runit \
	&& apt-get clean && rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*

COPY deploy/omnibus-cache/our_init /sbin/
CMD ["/sbin/our_init"]

RUN mkdir /etc/service/filebeat
COPY deploy/omnibus-cache/filebeat.sh /etc/service/filebeat/run
COPY deploy/omnibus-cache/filebeat.yml.tpl /etc/filebeat.yml.tpl

RUN mkdir -p /etc/service/cron
COPY deploy/omnibus-cache/cron.sh /etc/service/cron/run

RUN mkdir -p /etc/service/cache/log
COPY deploy/omnibus-cache/cache.sh /etc/service/cache/run
COPY deploy/omnibus-cache/cache.log.sh /etc/service/cache/log/run
COPY deploy/omnibus-cache/cache.log.config /var/log/cachecash/cache/config
# # See README.md for a note about this: we require that config be bind-mounted.
# COPY deploy/omnibus-cache/cache.config.json /etc/cache.config.json

COPY --from=filebeat /usr/share/filebeat/filebeat /usr/bin/
COPY --from=build /artifacts/bin/* /usr/local/bin/

# TODO: Still need containernet packages required on Ubuntu images.

# --------------------
