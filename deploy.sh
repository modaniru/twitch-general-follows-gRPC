#!/bin/sh
set -e

IMAGE="tgf:latest"

docker stop tgf
docker rm tgf

docker build -t $IMAGE .
