#!/bin/bash

export GOOGLE_APPLICATION_CREDENTIALS=./google_credentials.json

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

            \r  dev, -d         Reboot for dev environment.
            \r  testing, -t     Reboot for testing environment.
            \r  staging, -s     Reboot for staging environment.

            \r  --help, -h      Get help and exit."
        ;;
      *)
        exit 0
        ;;
    esac
fi

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
    reboot | down | -d)
        echo -e "\e[1;32mINFO\e[0m Deleting deployment..."
        kubectl delete -f $CONFIG_PATH
        [ $? -eq 0 ] && echo -e "\e[1;32mINFO\e[0m Deployment deleted!" || echo -e "\e[1;33mWARN\e[0m Error from server! Maybe deployment already deleted."
        ;;&
    down | -d)
        ;;
    reboot | -r | up | -u | build | -b)
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
    reboot | -r | up | -u | start | -s)
        echo -e "\e[1;32mINFO\e[0m Deploying..."
        kubectl create -f $CONFIG_PATH
        if [ $? -eq 0 ]
        then
            echo -e "\e[1;32mINFO\e[0m Deployed!"
        else
            echo -e "\e[1;31mFATA\e[0m Deployment failed, wrong exit code!"
            exit 1
        fi

        echo -e "\n\e[1mDeployment:\e[0m"
        kubectl get deployment
        echo -e "\e[1mSVC:\e[0m"
        kubectl get svc
        echo -e "\e[1mPVC:\e[0m"
        kubectl get pvc

        echo -e "\n\e[1mPODS:\e[0m"
        kubectl get pods
        echo -e "\n\e[1mServices:\e[0m"
        kubectl get services
        ;;
    --help | -h)
    echo -e "A tool for deploying the application with kubernetes.

        \r  reboot, -r      Takes these steps: deleting deployment; building and docker images; creating deployment.
        \r  down, -d        Deleting deployment only.
        \r  up, -u          Takes these steps: building and docker images; creating deployment.
        \r  build, -b       Building deployment only.
        \r  start, -s       Creating deployment only.

        \r  -f [PATH]       To specify configuration path. \e[1m$DEFAULT_CONFIG_PATH\e[0m by default.

        \r  local, -l       To deploy containers locally.

        \r  --help, -h      Getting help and exit.

        \rUsage: \e[1m./deploy.sh reboot -f [PATH]\e[0m"
    ;;
esac

exit 0
