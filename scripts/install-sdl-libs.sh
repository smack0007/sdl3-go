#!/bin/bash
set -e

SCRIPT_DIRECTORY="$(dirname "$(realpath "${BASH_SOURCE[-1]}")")"
TMP_DIR="./tmp"
SDL_TAG="release-3.2.4"

cd ${SCRIPT_DIRECTORY}

[[ -d ${TMP_DIR}/SDL ]] && rm -rf ${TMP_DIR}/SDL
git clone --depth 1 --branch ${SDL_TAG} https://github.com/libsdl-org/SDL.git ${TMP_DIR}/SDL
mkdir ${TMP_DIR}/SDL/build
cd ${TMP_DIR}/SDL/build
cmake -DCMAKE_BUILD_TYPE=Release ..
cmake --build . --config Release --parallel
sudo cmake --install . --config Release