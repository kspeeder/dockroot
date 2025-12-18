#!/bin/bash

set -euo pipefail

builds="amd64 arm64"
#builds="amd64"
#builds="arm64"

GV=$(git log --pretty=format:'%H' -n 1)

mkdir -p bin

for ARCH in $builds; do
  echo "Building for $ARCH"
  CGO_ENABLED=0 GOOS=linux GOARCH=${ARCH} GO111MODULE=on \
    go build -ldflags "-X main.gitCommit=${GV} -s -w" \
    -tags "containers_image_openpgp exclude_graphdriver_btrfs" \
    -o "bin/DockRoot.${ARCH}" ./cmd/dockroot

  if command -v upx >/dev/null 2>&1; then
    upx "bin/DockRoot.${ARCH}"
  else
    echo "upx not found; skipping compression for bin/DockRoot.${ARCH}"
  fi
done
