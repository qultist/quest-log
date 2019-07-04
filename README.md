# Quest Log
Simple *distributed* app built and packaged with Docker to be deployed on Kubernetes as playground/demo for my
bachelor's thesis. The project is intended to test tools like Skaffold and Telepresence in interaction with different
communication mechanisms between the services.

## Requirements
| Development          | Deployment           | Additional services |
|----------------------|----------------------|---------------------|
| Go                   | Docker               | Kafka               |
| protoc               | Docker-Compose       | Redis               |
| protoc plugin for Go | Kubernetes           |                     |

In order to test Skaffold and others, they have to be installed as well.

## Build, Package and Run
You can simply run the services locally or build and package them with Docker. Using multi-stage builds and the new
Go modules the build process is fairly quick. In any case make sure a Redis instance is running.
If you want to use Kafka, make sure it's set up. I've included Redis and Kafka in the docker-compose file.
On Kubernetes I'm using [Strimzi](https://strimzi.io) to run Kafka and a simple Deployment and Service
(see `deployments/k8s/redis`) to run Redis.

By default the services are expected at these locations:

| Service      | Location                     |
|--------------|------------------------------|
| gRPC service | grpc-service:50051           |
| HTTP service | http-service:8080            |
| Kafka        | ql-kafka-bootstrap:9092      |
| Zookeeper    | ql-zookeper-client:2181      |
| Redis        | redis:6379                   |

The values are used in the following files:
* `internal/app/gateway/*-handler.go`
* `cmd/kafka-service/main.go`
* `build/package/kafka-service/docker-entrypoint.sh`
* `internal/pkg/reposistory/repository.go`.
* `cmd/grpc-service/main.go` (port of the grpc-service)

### Run locally
**Not recommended!**
1. Change the addresses of the above mentioned services to match your local setup.
2. Spin up the services using `go run`.

### Build and run with Docker
You can build each image individually:
```bash
# For example the gateway
cd deployments/docker-compose
docker-compose build gateway
# Omit service name to build all
```
… or build and run them all at once:
```bash
cd deployments/docker-compose
docker-compose up -d --build
# --build forces a rebuild if the images already exist.
```
**Note:** You probably want to change the image names in the compose file.

### Run on Kubernetes
1. Build the images (see section above).
2. Push your images.
3. Change the image names in the manifests (`deployment/k8s/*`) to your image names.
4. Create the objects.
5. Profit. :)

## Develop on Kubernetes with Skaffold
The project comes with individual Skaffold configs for each service. Don't forget to change the image names!

### Local build
```bash
# For example the gateway
skaffold dev -f developlemt/skaffold/locally/gateway/skaffold.yaml
```

### In-cluster build
You can use Skaffold in combination with kaniko to run builds in the cluster. You can find an example for the
gateway in `development/skaffold/kaniko/gateway`. The configuration uses the `localDir` build context.

Although not using the GCS build context we must provide a secret. Simply create an empty secret with the following
command:
```bash
kubectl create secret generic kaniko-secret --from-literal=
``` 

Additionally to push to other secured registries you must provide a Docker `config.json` as a secret.
```bash
# config.json
{
    "auths": {
        "https://index.docker.io/v1/": {
            "auth": "dXNlcjpwYXNzd29yZAo="
	    }
    }
}
```

`auth` is your username and password separated by a colon and base64 encoded.
```bash
echo -n "user:password" | base64
```
## Development
 Compile protos
```bash
cd internal/pkg/protos
protoc quest-service.proto --go_out=plugins=grpc:.
```