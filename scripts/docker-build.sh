#!/bin/bash
set -e
. "$(dirname $(realpath "${BASH_SOURCE[0]}"))/../env.sh"

CONTAINER_NAME="smack0007/sdl-go"

cd ${REPO_PATH}
docker build \
  --label golang=${GO_VERSION} \
  --label sdl=${SDL_VERSION} \
  --build-arg GO_VERSION=${GO_VERSION} \
  --build-arg SDL_VERSION=${SDL_VERSION} \
  -t ${CONTAINER_NAME}:latest \
  .

CONTAINER_ID="$(docker images | grep $CONTAINER_NAME:latest | head -n 1 | awk '{print $3}')"

docker tag ${CONTAINER_ID} ${CONTAINER_NAME}:${GO_VERSION}_${SDL_VERSION}