#!/bin/sh

docker-compose down
sudo rm -rf databases/postgresql/db-data
docker-compose build
docker-compose up