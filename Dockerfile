ARG GO_DOCKER_VERSION=latest

FROM golang:${GO_DOCKER_VERSION}

LABEL maintainer="Zachary Snow <zachary.snow+docker@gmail.com>"

WORKDIR /app
COPY ./env.sh ./env.sh
COPY ./go.mod ./go.mod
COPY ./scripts/install-sdl-libs.sh ./scripts/install-sdl-libs.sh

RUN apt-get update && apt-get upgrade -qqy \
    && SUDO_CMD= /app/scripts/install-sdl-libs.sh \
    && rm -rf /var/lib/apt/lists/* \
    && rm -rf /app/tmp