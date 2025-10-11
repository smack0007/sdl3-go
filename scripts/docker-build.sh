#!/bin/bash
set -e
. "$(dirname $(realpath "${BASH_SOURCE[0]}"))/../env.sh"

CONTAINER_NAME="smack0007/sdl3-go"

cd ${REPO_PATH}
${DOCKER_EXE} build \
  --label golang=${GO_VERSION} \
  --label sdl=${SDL_VERSION} \
  --build-arg GO_VERSION=${GO_VERSION} \
  --build-arg SDL_VERSION=${SDL_VERSION} \
  --build-arg SDL_IMAGE_VERSION=${SDL_IMAGE_VERSION} \
  -t ${CONTAINER_NAME}:latest \
  -t ${CONTAINER_NAME}:go-${GO_VERSION}_sdl-${SDL_VERSION}_image-${SDL_IMAGE_VERSION} \
  .
