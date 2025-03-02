FROM golang:1.24.0-bookworm

LABEL maintainer="Zachary Snow <zachary.snow+docker@gmail.com>" golang=1.24.0 sdl=3.2.6

WORKDIR /app
COPY ./scripts/install-sdl-libs.sh ./scripts/install-sdl-libs.sh

RUN apt-get update && apt-get upgrade -qqy \
    && SUDO_CMD= /app/scripts/install-sdl-libs.sh \
    && rm -rf /var/lib/apt/lists/* \
    && rm -rf /app/tmp