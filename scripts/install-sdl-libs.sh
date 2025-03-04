#!/bin/bash
set -e
. "$(dirname $(realpath "${BASH_SOURCE[0]}"))/../env.sh"

SUDO_CMD=${SUDO_CMD:sudo}

if type "apt" > /dev/null; then
  ${SUDO_CMD} apt-get install --no-install-recommends -y build-essential git make \
    pkg-config cmake ninja-build gnome-desktop-testing libasound2-dev libpulse-dev \
    libaudio-dev libjack-dev libsndio-dev libx11-dev libxext-dev \
    libxrandr-dev libxcursor-dev libxfixes-dev libxi-dev libxss-dev \
    libxkbcommon-dev libdrm-dev libgbm-dev libgl1-mesa-dev libgles2-mesa-dev \
    libegl1-mesa-dev libdbus-1-dev libibus-1.0-dev libudev-dev \
    libpipewire-0.3-dev libwayland-dev libdecor-0-dev liburing-dev
fi

cd ${REPO_PATH}

[[ -d ${TMP_DIR}/SDL ]] && rm -rf ${TMP_DIR}/SDL
git clone --depth 1 --branch ${SDL_TAG} https://github.com/libsdl-org/SDL.git ${TMP_DIR}/SDL
mkdir ${TMP_DIR}/SDL/build
cd ${TMP_DIR}/SDL/build
cmake -DCMAKE_BUILD_TYPE=Release ..
cmake --build . --config Release --parallel
${SUDO_CMD} cmake --install . --config Release