#!/bin/sh
set -e

IMAGE="tgf:latest"
CONTAINER_ID=`docker ps -aqf "name=tgf"`

echo "TESTS"
echo "$CONTAINER_ID"
echo "${TWITCH_CLIENT_ID}"
echo "${TWITCH_CLIENT_SECRET}"

docker stop $CONTAINER_ID


docker build -t $IMAGE .

docker run --restart unless-stopped -d -p 80:8080 --name tgf -e TWITCH_CLIENT_ID=${TWITCH_CLIENT_ID} -e TWITCH_CLIENT_SECRET=${TWITCH_CLIENT_SECRET} $IMAGE
docker system prune -a -f
