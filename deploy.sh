#!/bin/sh
set -e

IMAGE="tgf:latest"

docker stop tgf || true
docker rm tgf || true

docker build -t $IMAGE .
