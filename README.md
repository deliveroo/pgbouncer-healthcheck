# PGBouncer healthcheck

## Introduction

This implements a http `/health` check for PGBouncer hosts.
This healthcheck provides a simple yes/no health response based
on attempting to connect to PGBouncer and execute a request, as
well as checking the Datadog agent.

There are also several endpoints used for diagnostics.

## Building and running

This can be built on a local system with docker using the provided `Makefile`.

    $ make build

You can then run the server binary

    $ ./pgbouncer-healthcheck


## Endpoints

#### /

This always returns a `200 OK`

#### /health
The healthcheck fails with a `500 Internal Server Error` response.

TODO: Add details of other endpoints
