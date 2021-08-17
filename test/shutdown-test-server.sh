#!/usr/bin/sh

IMAGE="test-server-image"
CONTAINER="test-server-cantainer"

# Shutdown and clean docker image
docker kill "${CONTAINER}"
docker rmi "${IMAGE}"
