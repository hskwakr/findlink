#!/usr/bin/sh

IMAGE="test-server-image"
CONTAINER="test-server-cantainer"

# Build and run docker image
docker build -t "${IMAGE}" ./test
docker run -dit --name "${CONTAINER}" --rm  -p 8080:80 "${IMAGE}"