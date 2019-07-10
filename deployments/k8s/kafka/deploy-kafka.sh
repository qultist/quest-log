#!/usr/bin/env bash

# Download and extract strimzi
curl -Lo strimzi.tar.gz https://github.com/strimzi/strimzi-kafka-operator/releases/download/0.12.1/strimzi-0.12.1.tar.gz
mkdir strimzi
tar -xzf strimzi.tar.gz -C strimzi --strip-components=1
cd strimzi

# Deploy Cluster Operator in default namespace
sed -i '' 's/namespace: .*/namespace: default/' install/cluster-operator/*RoleBinding*.yaml
kubectl apply -f install/cluster-operator -n default

# Deploy ephemeral Kafka cluster without replication
cd ..
kubectl apply -f kafka.yaml

# Clean up
rm -rf strimzi*