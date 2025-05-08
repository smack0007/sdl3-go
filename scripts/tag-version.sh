#!/bin/bash
set -e
. "$(dirname $(realpath "${BASH_SOURCE[0]}"))/../env.sh"

if [[ $1 == "" ]]; then
  echo "ERROR: Please provide a tag name."
  exit 1
fi

git tag -a $1 -m "$(date +%Y-%m-%d) go ${GO_VERSION} SDL ${SDL_VERSION} SDL_image ${SDL_IMAGE_VERSION}"
git push origin tag $1