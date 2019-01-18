FROM golang:1.10.2-alpine as builder

RUN apk add --no-cache alpine-sdk

ARG DIR=${GOPATH}/src/github.com/deliveroo/pgbouncer-healthcheck
WORKDIR $DIR

RUN apk add --update pgbouncer bash shadow && \
    mkdir -p /var/run/postgresql && \
    chown pgbouncer /var/run/postgresql

ADD vendor.sh $DIR/
ADD build.sh $DIR/
ADD Gopkg.toml $DIR/
ADD Gopkg.lock $DIR/
ADD *.go $DIR/
ADD VERSION $DIR/

ADD tests/pgbouncer.ini /etc/pgbouncer/pgbouncer.ini
ADD tests/userlist.txt /etc/pgbouncer/userlist.txt
ADD tests/scripts /tests
RUN chmod 755 /tests/*

RUN $DIR/vendor.sh
RUN $DIR/build.sh

