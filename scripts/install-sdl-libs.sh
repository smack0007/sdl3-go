#!/bin/bash
set -e
. "$(dirname $(realpath "${BASH_SOURCE[0]}"))/../env.sh"

# Allow Dockerfile to set SUDO_CMD= as the sudo command isn't needed in container.
SUDO_CMD=${SUDO_CMD-sudo}

BASE_SDL_URL="https://github.com/libsdl-org/"

LIBS="SDL SDL_image"

declare -A GIT_TAGS=(
  ["SDL"]="${SDL_TAG}"
  ["SDL_image"]="${SDL_IMAGE_TAG}"
)

if type "apt" > /dev/null; then
  # "build-essential", "make" and "pkg-config" are assumed to be installed 
  ${SUDO_CMD} apt-get install --no-install-recommends -y \
    cmake ninja-build gnome-desktop-testing libasound2-dev libpulse-dev \
    libaudio-dev libjack-dev libsndio-dev libx11-dev libxext-dev \
    libxrandr-dev libxcursor-dev libxfixes-dev libxi-dev libxss-dev \
    libxkbcommon-dev libdrm-dev libgbm-dev libgl1-mesa-dev libgles2-mesa-dev \
    libegl1-mesa-dev libdbus-1-dev libibus-1.0-dev libudev-dev \
    libpipewire-0.3-dev libwayland-dev libdecor-0-dev liburing-dev
fi

cd ${REPO_PATH}

[[ -d ${TMP_DIR} ]] && rm -rf ${TMP_DIR}
mkdir ${TMP_DIR}
for lib in $LIBS; do
  git clone --depth 1 --branch ${GIT_TAGS[$lib]} ${BASE_SDL_URL}${lib}.git ${TMP_DIR}/$lib
  mkdir ${TMP_DIR}/${lib}/build
  cd ${TMP_DIR}/${lib}/build
  cmake -DCMAKE_BUILD_TYPE=Release ..
  cmake --build . --config Release --parallel
  ${SUDO_CMD} cmake --install . --config Release
done

if [[ "${CLEANUP}" = "1" ]]; then
  if type "apt" > /dev/null; then
    ${SUDO_CMD} apt-get remove -y \
      cmake ninja-build gnome-desktop-testing libasound2-dev libpulse-dev \
      libaudio-dev libjack-dev libsndio-dev libx11-dev libxext-dev \
      libxrandr-dev libxcursor-dev libxfixes-dev libxi-dev libxss-dev \
      libxkbcommon-dev libdrm-dev libgbm-dev libgl1-mesa-dev libgles2-mesa-dev \
      libegl1-mesa-dev libdbus-1-dev libibus-1.0-dev libudev-dev \
      libpipewire-0.3-dev libwayland-dev libdecor-0-dev liburing-dev
  fi
fi