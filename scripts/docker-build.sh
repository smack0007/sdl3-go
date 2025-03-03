#!/bin/bash
set -e
. "$(dirname $(realpath "${BASH_SOURCE[0]}"))/../env.sh"

cd ${REPO_PATH}
docker build \
  --label golang=${GO_VERSION} \
  --label sdl=${SDL_VERSION} \
  --build-arg GO_DOCKER_VERSION=${GO_DOCKER_VERSION} \
  -t smack0007/sdl-go:${GO_VERSION}_${SDL_VERSION} \
  .