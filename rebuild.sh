#!/bin/sh

docker-compose down
sudo rm -rf db-data
docker-compose build
docker-compose up