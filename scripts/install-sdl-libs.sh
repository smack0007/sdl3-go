#!/bin/bash
set -e

SCRIPT_DIRECTORY="$(dirname "$(realpath "${BASH_SOURCE[-1]}")")"
TMP_DIR="./tmp"
SDL_TAG="release-2.30.8"

cd ${SCRIPT_DIRECTORY}

[[ -d ${TMP_DIR}/SDL ]] && rm -rf ${TMP_DIR}/SDL
git clone --depth 1 --branch ${SDL_TAG} https://github.com/libsdl-org/SDL.git ${TMP_DIR}/SDL
mkdir ${TMP_DIR}/SDL/build
cd ${TMP_DIR}/SDL/build
../configure
make -j8
sudo make install