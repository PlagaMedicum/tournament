#!/bin/sh

case "$1" in
  --dev | -d)
    docker-compose down;
    sudo rm -rf databases/postgresql/data;
    docker-compose build;
    docker-compose up;;
  --testing | -t)
    docker-compose -f ./env/testing/docker-compose.yml down;
    docker-compose -f ./env/testing/docker-compose.yml build;
    docker-compose -f ./env/testing/docker-compose.yml up;;
  --staging | -s)
    docker-compose -f ./env/staging/docker-compose.yml down;
    sudo rm -rf env/staging/databases/postgresql/data;
    docker-compose -f ./env/staging/docker-compose.yml build;
    docker-compose -f ./env/staging/docker-compose.yml up;;
  --help | -h)
    echo "
    This script used to recreate and set up docker-compose files
    with ability to choose environment.

    --dev, -d      Run in dev environment.
    --testing, -t  Run in testing environment.
    --staging, -s  Run in staging environment.

    --help, -h     Get help and exit.
    ";;
esac