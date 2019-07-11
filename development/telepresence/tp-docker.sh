#!/usr/bin/env bash

SERVICE=grpc-service
PORT=50051
LOGDIR=`pwd`

cd ../../deployments/docker-compose
docker-compose build ${SERVICE}
telepresence --new-deployment ${SERVICE} ${PORT+--expose ${PORT}} \
    --logfile ${LOGDIR}/telepresence.log \
    --docker-run --rm ${PORT+-p ${PORT}:$PORT} qultist/ql-${SERVICE}:0.1