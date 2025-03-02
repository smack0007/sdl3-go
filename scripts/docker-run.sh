#!/bin/bash
set -e
SCRIPT_DIRECTORY="$(dirname $(realpath "${BASH_SOURCE[0]}"))"

cd ${SCRIPT_DIRECTORY}/..
docker run -it --rm -v .:/app smack0007/sdl-go /bin/sh