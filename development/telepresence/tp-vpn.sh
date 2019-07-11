#!/usr/bin/env bash

SERVICE=kafka-service
#PORT=50051

go install ../../cmd/${SERVICE}
telepresence --new-deployment ${SERVICE} ${PORT+--expose ${PORT}} \
    --run ${GOPATH}/bin/${SERVICE}
