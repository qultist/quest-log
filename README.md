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
Go modules the build process is fairly quick. In any case make sure a Redis instance is running.

### Run locally
1. In `cmd/gateway/main.go` change the address for example to `localhost`.
Do the same in `internal/pkg/repository/repository.go`
2. Spin up the services using `go run`.

### Run with Docker
You can build each image individually:
```bash
# For example the gateway
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

### Run in Kubernetes
1. Build the images (see section above).
2. Push your images.
3. Change the images in the manifests `deployment/k8s/*` to yours.
4. Create the objects
5. Profit. :)

## Develop in Kubernetes with Skaffold
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