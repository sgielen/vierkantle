#!/bin/bash

set -e

PUSH=0
DEPLOY=0
if [ "$1" = "--push" ]; then
	PUSH=1
elif [ "$1" = "--deploy" ]; then
	PUSH=1
	DEPLOY=1
elif [ "$1" != "" ]; then
	echo "Usage: $0 [ --push | --deploy ]"
	exit 1
fi

IMAGE_BASENAME="rg.nl-ams.scw.cloud/sc-ams-registry/vierkantle/"

GIT_REF_NAME="$(git symbolic-ref -q --short HEAD || git describe --tags --exact-match)"

export DOCKER_BUILDKIT=1

PLATFORM="linux/amd64"

pushd cmd/ui
docker build . \
	-t "${IMAGE_BASENAME}ui:latest-${GIT_REF_NAME}" \
	--platform "$PLATFORM"
if [ "$PUSH" = "1" ]; then
	docker push "${IMAGE_BASENAME}ui:latest-${GIT_REF_NAME}"
fi
popd

for gobinary in backend; do
	docker build . \
		--target "$gobinary" \
		--platform "$PLATFORM" \
		-t "${IMAGE_BASENAME}${gobinary}:latest-${GIT_REF_NAME}"
	if [ "$PUSH" = "1" ]; then
		docker push "${IMAGE_BASENAME}${gobinary}:latest-${GIT_REF_NAME}"
	fi
done

if [ "$DEPLOY" = "1" ]; then
	kubectl rollout restart deployment -n vierkantle ui
	kubectl rollout restart deployment -n vierkantle backend
fi
