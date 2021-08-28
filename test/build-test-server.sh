#!/usr/bin/sh

IMAGE="test-server-image"
CONTAINER="test-server-cantainer"

# Build and run docker image
docker build -t "${IMAGE}" ./test
docker run -dit --name "${CONTAINER}" --rm  -p 8080:80 "${IMAGE}"

# Copy this project to container
docker cp . "${CONTAINER}":/home/working 

# Run integration test
docker exec "${CONTAINER}" /usr/local/go/bin/go test -tags=integration -v
