#!/bin/sh

echo dump | nc ql-zookeper-client 2181  | grep -q brokers
while [ $? -eq 1 ]; do
    sleep 5
    echo dump | nc ql-zookeper-client 2181 | grep -q brokers
done
./kafka-service