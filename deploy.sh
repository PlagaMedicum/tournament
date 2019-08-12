#!/bin/bash

if [ "$1" == "local" ] || [ "$1" == "-l" ]
then
  case "$2" in
    dev | -d)
      docker-compose down
      docker-compose build
      docker-compose up
    ;;
    testing | -t)
      docker-compose -f ./env/testing/docker-compose.yml down
      docker-compose -f ./env/testing/docker-compose.yml build
      docker-compose -f ./env/testing/docker-compose.yml up
    ;;
    staging | -s)
      docker-compose -f ./env/staging/docker-compose.yml down
      sudo rm -rf env/staging/databases/postgresql/data
      docker-compose -f ./env/staging/docker-compose.yml build
      docker-compose -f ./env/staging/docker-compose.yml up
    ;;
    --help | -h)
      echo -e "Local deploying composed of these steps: getting containers down; building containers; getting them up.

        \r  dev, -d       Reboot for dev environment.
        \r  testing, -t   Reboot for testing environment.
        \r  staging, -s   Reboot for staging environment.

        \r  --help, -h    Get help and exit.

        \rUsage: \e[1m./deploy.sh local dev\e[0m"
    ;;
    *)
      exit 0
    ;;
  esac
fi

export GOOGLE_APPLICATION_CREDENTIALS=./google_credentials.json

SERVER_NAME=server
SERVER_IMAGE=plagamedicum/tournament_server:tag

DEFAULT_CONFIG_PATH=./env/staging/k8s/
if [ "$2" == "-f" ]
then
  CONFIG_PATH=$3
  echo -e "\e[1;32mINFO\e[0m Configuration path set to \e[1m$CONFIG_PATH\e[0m.\n"
else
  CONFIG_PATH=$DEFAULT_CONFIG_PATH
fi

case "$1" in
  reset | -r | down | -d)
    echo -e "\e[1;32mINFO\e[0m Deleting deployment..."
    kubectl delete -f $CONFIG_PATH
    [ $? -eq 0 ] && echo -e "\e[1;32mINFO\e[0m Deployment deleted!" || echo -e "\e[1;33mWARN\e[0m Error from server! Maybe deployment already deleted."
  ;;&
  down | -d)
  ;;
  reset | -r | up | -u | build | -b)
    echo -e "\e[1;32mINFO\e[0m Building container for \e[1m$SERVER_NAME\e[0m service..."
    sudo rm -rf env/staging/databases/postgresql/data
    docker-compose -f ./env/staging/docker-compose.yml down
    docker-compose -f ./env/staging/docker-compose.yml build $SERVER_NAME
    [ $? -eq 0 ] && echo -e "\e[1;32mINFO\e[0m Image building completed!" || echo -e "\e[1;33mWARN\e[0m Cannot build docker image, wrong exit code."

    echo -e "\e[1;32mINFO\e[0m Pushing \e[1m$SERVER_IMAGE\e[0m image to \e[1mdocker.io\e[0m..."
    docker push $SERVER_IMAGE
    [ $? -eq 0 ] && echo -e "\e[1;32mINFO\e[0m Image pushing completed!" || echo -e "\e[1;33mWARN\e[0m Cannot push docker image, wrong exit code."
  ;;&
  build | -b)
  ;;
  reset | -r | up | -u | start | -s)
    echo -e "\e[1;32mINFO\e[0m Deploying..."
    kubectl create -f $CONFIG_PATH
    if [ $? -eq 0 ]
    then
      echo -e "\e[1;32mINFO\e[0m Deployed!"
    else
      echo -e "\e[1;31mFATA\e[0m Deployment failed, wrong exit code!"
      exit 1
    fi
  ;;&
  reset | -r | up | -u | start | -s | info | -i)
    echo -e "\n\e[1;36mDeployment:\e[0m"
    kubectl get deployment
    echo -e "\e[1;36mSVC:\e[0m"
    kubectl get svc
    echo -e "\e[1;36mPVC:\e[0m"
    kubectl get pvc

    echo -e "\n\e[1;36mPODS:\e[0m"
    kubectl get pods
    echo -e "\n\e[1;36mServices:\e[0m"
    kubectl get services
  ;;
  --help | -h)
    echo -e "A tool for deploying the application with kubernetes in gcloud.

      \r  reset, -r     Take these steps: delete deployment; build and pushing docker images; create deployment.
      \r  down, -d      Delete deployment only.
      \r  up, -u        Take these steps: build and docker images; create deployment.
      \r  build, -b     Build deployment only.
      \r  start, -s     Create deployment only.
      \r  info, -i      Get info about deployment(Deployment, SVC, PVC, Pods, Services).

      \r  -f [PATH]     Specify configuration path. \e[1m$DEFAULT_CONFIG_PATH\e[0m by default.

      \r  local, -l     Deploy containers locally.

      \r  --help, -h    Get help and exit.

      \rUsage: \e[1m./deploy.sh reset -f [PATH]\e[0m"
  ;;
esac

exit 0
