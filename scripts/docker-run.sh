#!/bin/bash
set -e
. "$(dirname $(realpath "${BASH_SOURCE[0]}"))/../env.sh"

DOCKER_FLAGS="--rm -v .:/data/sdl-go ${DOCKER_FLAGS}"

if [[ ! "${DOCKER_FLAGS}" == *" -w "* ]]; then
  DOCKER_FLAGS="${DOCKER_FLAGS} -w /data/sdl-go"
fi

if [[ ! "${CI}" = "1" ]]; then 
  DOCKER_FLAGS="${DOCKER_FLAGS} -it"
fi

CONTAINER_NAME="smack0007/sdl-go:${GO_VERSION}_${SDL_VERSION}"

DOCKER_CMD=${1:-/bin/bash}

cd ${REPO_PATH}
docker run ${DOCKER_FLAGS} ${CONTAINER_NAME} ${DOCKER_CMD}
