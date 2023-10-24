#!/bin/sh
cd docker
. ../.env
echo "\e[34m >>> Creating external network \e[97m"

docker network create $PROJECT_NAME"_network"

if [ $? -eq 0 ]; then
    echo "\e[32m >>> Network created \e[97m"
else
    echo "\e[31m >>> Error occurred \e[97m"
    exit 1
fi
