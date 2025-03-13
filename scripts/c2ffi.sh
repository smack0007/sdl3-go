#!/bin/bash
set -e
. "$(dirname $(realpath "${BASH_SOURCE[0]}"))/../env.sh"

TMP_PATH="${REPO_PATH}/tmp"
C2FFI_BRANCH="llvm-18.1.0"

cd ${REPO_PATH}

if [[ "${1}" == "--rebuild" ]]; then
  [[ -d ${TMP_PATH}/c2ffi ]] && rm -rf ${TMP_PATH}/c2ffi
  mkdir -p "${TMP_PATH}"
  git clone --depth 1 --branch ${C2FFI_BRANCH} https://github.com/rpav/c2ffi.git ${TMP_PATH}/c2ffi
  cd "${TMP_PATH}/c2ffi"
  docker build -f ./Docker/Test-Build-Ubuntu-20.04.docker -t c2ffi .
  cd "${REPO_PATH}"
fi

SDL_INCLUDE_PATH="$(pkg-config sdl3 --cflags | cut -b 3-)"
mkdir -p ${TMP_PATH}/include/SDL3
cp -r ${SDL_INCLUDE_PATH}/SDL3/* ${TMP_PATH}/include/SDL3
mkdir -p ${TMP_PATH}/include/SDL3_image
cp -r ${SDL_INCLUDE_PATH}/SDL3_image/* ${TMP_PATH}/include/SDL3_image

cat >${TMP_PATH}/SDL.c <<EOF
#include <SDL3/SDL.h>
#include <SDL3_image/SDL_image.h>
EOF

C2FFI_CMD="docker run --rm -it"
C2FFI_CMD+=" -v ${TMP_PATH}/include/SDL3:/usr/local/include/SDL3:ro" # Map SDL
C2FFI_CMD+=" -v ${TMP_PATH}/include/SDL3_image:/usr/local/include/SDL3_image:ro" # Map SDL_image
C2FFI_CMD+=" -v ${TMP_PATH}/SDL.c:/usr/local/SDL.c:ro" # Map SDL_image
C2FFI_CMD+=" c2ffi" # Container name
C2FFI_CMD+=" /c2ffi/build/bin/c2ffi /usr/local/SDL.c"
# C2FFI_CMD+=" /bin/bash"
${C2FFI_CMD} \
  | jq -c '[.[] | select(.location | startswith("/usr/include/") or startswith("/usr/lib") | not)]' \
  | sed 's/\/usr\/local\/include\/SDL3/\/SDL3/g' \
  > ${REPO_PATH}/tools/codegen/sdl.json