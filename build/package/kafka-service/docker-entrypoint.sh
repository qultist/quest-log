#!/bin/sh

echo dump | nc zookeeper 2181 | grep -q brokers
while [ $? -eq 1 ]; do
    sleep 5
    echo dump | nc zookeeper 2181 | grep -q brokers
done
./kafka-service