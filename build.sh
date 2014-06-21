#!/usr/bin/env bash

PLATFORMS="darwin/386 darwin/amd64 freebsd/386 freebsd/amd64 linux/386
linux/amd64 linux/arm"

cd `dirname $0`

for PLATFORM in $PLATFORMS; do
  export GOOS=${PLATFORM%/*}
  export GOARCH=${PLATFORM#*/}
	RELEASE_PATH=release/${GOOS}-${GOARCH}
	mkdir -p $RELEASE_PATH
  CMD="go build -o ${RELEASE_PATH}/docker-exec ."
  echo "GOOS=${GOOS} GOARCH=${GOARCH} $CMD"
  $CMD
	PACKAGE="tar -czvf release/${GOOS}-${GOARCH}.tar.gz ${RELEASE_PATH}"
	$PACKAGE
	CLEANUP="rm -rf ${RELEASE_PATH}"
	$CLEANUP
done
