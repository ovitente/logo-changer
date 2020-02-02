#!/bin/bash

cd ./bin

cp ../videos.list ../info.list .
echo ${CONFIG_FILE} | base64 -d > config.yml

./logo-changer >> logo-changer.log
