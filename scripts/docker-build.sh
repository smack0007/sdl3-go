#!/bin/bash
set -e
SCRIPT_DIRECTORY="$(dirname $(realpath "${BASH_SOURCE[0]}"))"

cd ${SCRIPT_DIRECTORY}/..
docker build -t smack0007/sdl-go .