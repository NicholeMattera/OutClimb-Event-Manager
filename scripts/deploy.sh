#!/bin/sh

echo "🧹 Cleaning up temp directory"
rm -rf /tmp/outclimb-event-manager

echo "📦 Downloading the latest"
git clone --depth 1 https://github.com/NicholeMattera/OutClimb-Event-Manager.git /tmp/outclimb-event-manager
cd /tmp/outclimb-event-manager

echo "🩹 Patching configs"
patch -i ~/outclimb-event-manager-secrets.diff

echo "🛠️ Building"
docker compose build --no-cache

echo "🪦 Bringing down and removing old container"
docker container stop outclimb-event-manager-service-1
docker container rm outclimb-event-manager-service-1

echo "💡 Bringing up new container"
docker compose up --detach

echo "🧹 Cleaning up old images and temp directory"
docker image prune --all --force
rm -rf /tmp/outclimb-event-manager