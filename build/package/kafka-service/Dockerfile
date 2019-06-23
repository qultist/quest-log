FROM golang:alpine AS builder
ENV GO111MODULE=on
WORKDIR /go/src/github.com/qultist/quest-log
RUN apk add --no-cache git
COPY go.mod go.sum ./
RUN go mod download
COPY cmd/kafka-service ./cmd/kafka-service
COPY internal ./internal
RUN go build -o /go/bin/kafka-service cmd/kafka-service/main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates netcat-openbsd
COPY build/package/kafka-service/docker-entrypoint.sh /docker-entrypoint.sh
COPY --from=builder /go/bin/kafka-service /kafka-service
ENTRYPOINT ./kafka-service
