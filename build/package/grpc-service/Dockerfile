FROM golang:alpine AS builder
ENV GO111MODULE=on
WORKDIR /go/src/github.com/qultist/quest-log
RUN apk add --no-cache git
COPY go.mod go.sum ./
RUN go mod download
COPY cmd/grpc-service ./cmd/grpc-service
COPY internal ./internal
RUN go build -o /go/bin/grpc-service cmd/grpc-service/main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /go/bin/grpc-service /grpc-service
ENTRYPOINT ./grpc-service
EXPOSE 50051
