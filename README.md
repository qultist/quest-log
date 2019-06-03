# Quest Log
Simple *distributed* Go app built and packaged with Docker to be deployed on Kubernetes as playground for my bachelor
thesis. The project is intended to test tools like skaffold and Telepresence in interaction with different communication
mechanisms between the services.

## Requirements
- Go
- protoc
- protoc plugin for Go
- Redis
- Docker
- Kubernetes

## Build, Package and Run
You can simply run the services locally or build and package them with Docker. Using multi-stage builds and the new
Go modules the build process is fairly quick. In each case make sure a Redis instance is running.

### Running locally
1. Change the adresses to localhost in the gateway
2. Spin up the services using `go run`.

### Running with Docker
You can build each image individually:
```bash
# Gateway for example
docker build -t qultist/ql-gateway:0.1 -f build/package/gateway/Dockerfile .
```
… or leave it to docker-compose:
```bash
cd deployments/docker-compose
docker-compose up -d
# Append --build to rebuild if the images already exist.
# Or use 'docker compose build' to build the images only.
```
**Note:** You probably want to change the image names in the compose file. ;)

## Development
 Compile protos
```bash
cd internal/pkg/protos
protoc quest-service.proto --go_out=plugins=grpc:.
```