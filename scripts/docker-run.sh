#!/bin/bash
set -e
. "$(dirname $(realpath "${BASH_SOURCE[0]}"))/../env.sh"

DOCKER_FLAGS="--rm -v .:/app"
if [[ ! "${CI}" = "1" ]]; then 
  DOCKER_FLAGS="${DOCKER_FLAGS} -it"
fi

DOCKER_CMD=${1:-/bin/bash}

cd ${REPO_PATH}
docker run ${DOCKER_FLAGS} smack0007/sdl-go ${DOCKER_CMD}