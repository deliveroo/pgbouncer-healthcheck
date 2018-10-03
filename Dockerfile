FROM golang:1.10.2-alpine as builder

RUN apk add --no-cache alpine-sdk

ARG DIR=${GOPATH}/src/github.com/deliveroo/pgbouncer-healthcheck
WORKDIR $DIR/cmd

ADD vendor.sh $DIR/
ADD Gopkg.toml $DIR/
ADD Gopkg.lock $DIR/
RUN ../vendor.sh

ADD build.sh $DIR/
ADD ./cmd $DIR/cmd
RUN ../build.sh
