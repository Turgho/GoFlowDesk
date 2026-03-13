# Dockerfile left for reference only.  The application binary is built
# and run locally during development; only the database is containerized
# via docker-compose.
#
# To build an API image manually you can run:
#
#   docker build -t goflowdesk-api .
#
# but `docker-compose` no longer uses this file.

FROM alpine:3.23 AS base

LABEL maintainer="GoFlowDesk Team"

# no binaries provided by default; this image is a placeholder.
ENTRYPOINT []
