# Telepresence

Three scripts are included to test the different proxying methods of Telepresence.
The `tcp-inject` method is not tested.

Set SERVICE and PORT variables to the desired values:

|               | SERVICE       | PORT         |
|---------------|---------------|--------------|
| gateway       | gateway       | 80           |
| grpc-service  | grpc-service  | 50051        |
| kafka-service | kafka-service | Comment out! |

The scripts must be executed within this folder!

## tp-vpn.sh
Builds and runs the service locally.

## tp-docker.sh
Builds and runs the service using Docker.

## tp-debug.sh
Creates the proxy and outputs a JSON file containing various environment variables.
This file can be loaded using the [EnvFile](https://plugins.jetbrains.com/plugin/7861-envfile) plugin for JetBrains
IDEs. Then you can start/debug your service from the IDE like usual.