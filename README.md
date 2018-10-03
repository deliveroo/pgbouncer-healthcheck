# EC2 docker healthcheck

## Introduction

This implements a http `/health` check for ECS hosts.

Similar to web containers that use ALB healthchecks to determine liveliness,
we can healthcheck the EC2 hosts if the `dockerd` is behaving correctly.

## Building and running

This can be built on a local system with docker using the provided `Makefile`.

    $ make build

You can then run the server binary

    $ ./ec2-docker-healthcheck


## Endpoints

#### /

This always returns a `200 OK`

#### /health

This endpoint checks if the local dockerd unix socket is available, and that
the current storage graph driver is `devicemapper` to return a `200 OK`.

The healthcheck fails with a `500 Internal Server Error` response.
