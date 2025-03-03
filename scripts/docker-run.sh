#!/bin/bash
set -e
. "$(dirname $(realpath "${BASH_SOURCE[0]}"))/../env.sh"

cd ${REPO_PATH}
docker run -it --rm -v .:/app smack0007/sdl-go /bin/sh