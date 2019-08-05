#!/bin/sh

case "$1" in
  --dev | -d)
    docker-compose down;
    sudo rm -rf databases/postgresql/db-data;
    docker-compose build;
    docker-compose up;;
  --testing | -t)
    docker-compose -f ./env/testing/docker-compose.yml down;
    sudo rm -rf databases/postgresql/db-data;
    docker-compose -f ./env/testing/docker-compose.yml build;
    docker-compose -f ./env/testing/docker-compose.yml up;;
  --staging | -s)
    docker-compose -f ./env/staging/docker-compose.yml down;
    sudo rm -rf databases/postgresql/db-data;
    docker-compose -f ./env/staging/docker-compose.yml build;
    docker-compose -f ./env/staging/docker-compose.yml up;;
  --help | -h)
    echo "
    This script used to recreate and set up docker-compose files
    with ability to choose environement.

    --dev, -d      Run in dev environement.
    --testing, -t  Run in testing environement.
    --staging, -s  Run in staging environement.

    --help, -h     Get help and exit.
    ";;
esac