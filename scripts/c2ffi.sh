#!/bin/bash
set -e

SCRIPT_DIRECTORY="$(dirname "$(realpath "${BASH_SOURCE[-1]}")")"
TMP_DIR="./tmp"
C2FFI_BRANCH="llvm-18.1.0"

cd ${SCRIPT_DIRECTORY}

# [[ -d ${TMP_DIR}/c2ffi ]] && rm -rf ${TMP_DIR}/c2ffi
# git clone --depth 1 --branch ${C2FFI_BRANCH} https://github.com/rpav/c2ffi.git ${TMP_DIR}/c2ffi
# cd ${TMP_DIR}/c2ffi
# docker build -f ./Docker/Test-Build-Ubuntu-20.04.docker -t c2ffi .

cd ${SCRIPT_DIRECTORY}

C2FFI_CMD="docker run --rm -it -v $(pkg-config sdl3 --cflags | cut -b 3-):/usr/local/include c2ffi /c2ffi/build/bin/c2ffi"
${C2FFI_CMD} --help
#${C2FFI_CMD} /usr/local/include/SDL3/SDL.h > ../tools/codegen/sdl.json