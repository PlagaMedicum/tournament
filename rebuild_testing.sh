#!/bin/sh

docker-compose -f ./env/testing/docker-compose.yml down
sudo rm -rf databases/postgresql/db-data
docker-compose -f ./env/testing/docker-compose.yml build
docker-compose -f ./env/testing/docker-compose.yml up