#!/usr/bin/env bash

SERVICE=grpc-service
PORT=50051

telepresence --new-deployment ${SERVICE} ${PORT+--expose ${PORT}} --env-json ${SERVICE}-env.json