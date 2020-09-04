#!/bin/bash

## execute docker-compose up
cd ../build
docker-compose up -d --build

## Run api server
cid=`docker container ls | grep api | cut -b 1-12`
docker container exec -it $cid /bin/sh -c "cd /go/src/github.com/nhannvt/resume && go run cmd/main.go --debug"