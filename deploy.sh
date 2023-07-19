#!/bin/sh
set -e

IMAGE="tgf:latest"
CONTAINER_ID=`docker ps -aqf "name=tgf"`

docker build -t IMAGE .

docker stop $CONTAINER_ID
docker run --restart unless-stopped -d -p 80:8080 --name tgf -e TWITCH_CLIENT_ID=${TWITCH_CLIENT_ID} -e TWITCH_CLIENT_SECRET=${TWITCH_CLIENT_SECRET} ${IMAGE}
docker system prune -a -f
