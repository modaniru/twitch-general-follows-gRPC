#!/bin/sh
set -e

IMAGE="tgf:latest"
CONTAINER_ID=`docker ps -aqf "name=tgf"`

docker stop $CONTAINER_ID
docker rm $CONTAINER_ID

docker build -t $IMAGE .
