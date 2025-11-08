FROM debian:trixie-slim
LABEL maintainer="Zachary Snow <zachary.snow+docker@gmail.com>"

ARG GO_VERSION
ARG SDL_VERSION
ARG SDL_IMAGE_VERSION

ENV GO_VERSION=$GO_VERSION
ENV SDL_VERSION=$SDL_VERSION
ENV SDL_IMAGE_VERSION=$SDL_IMAGE_VERSION

ENV GOLANG_VERSION=$GO_VERSION
ENV GOTOOLCHAIN=local
ENV GOPATH=/go
ENV PATH=/usr/local/go/bin:$PATH

WORKDIR /data
COPY ./env.sh ./go.mod ./
COPY ./scripts/install-sdl-libs.sh ./scripts/install-sdl-libs.sh

RUN set -eux; \
	arch="$(dpkg --print-architecture)"; arch="${arch##*-}"; \
	url="https://dl.google.com/go/go${GO_VERSION}.linux-$arch.tar.gz"; \
  apt-get update && apt-get upgrade -qqy && apt-get install --no-install-recommends -y build-essential ca-certificates git make pkg-config wget; \
  wget -O go.tgz "$url" --progress=dot:giga; \
  tar -C /usr/local -xzf go.tgz; \
	rm go.tgz; \
  mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 1777 "$GOPATH"; \
  go version; \
  SUDO_CMD= CLEANUP=1 /data/scripts/install-sdl-libs.sh; \
  apt-get remove -y ca-certificates git wget && apt-get autoremove -y && apt-get autoclean -y; \
  rm -rf /data/tmp && rm -rf /var/lib/apt/lists/*;
    
