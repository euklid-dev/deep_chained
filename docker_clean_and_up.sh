#!/bin/sh

echo "Clearing all dangling Docker build caches..."
docker builder prune -f

echo "Clearing all Docker build caches, including used ones..."
docker builder prune -a -f

echo "Removing all stopped containers..."
docker container prune -f

echo "Removing all unused images..."
docker image prune -a -f

echo "Removing all unused networks..."
docker network prune -f

echo "Removing all unused volumes..."
docker volume prune -f

echo "Removing all Docker data (containers, images, volumes, and networks)..."
docker system prune -a --volumes -f

echo "Starting Docker Compose without using cache..."
docker-compose up --build --force-recreate

echo "Docker Compose started without caching."
