FROM golang:alpine AS builder
ENV GO111MODULE=on
WORKDIR /go/src/github.com/qultist/quest-log
RUN apk add --no-cache git
COPY go.mod go.sum ./
RUN go mod download
COPY cmd/gateway ./cmd/gateway
COPY internal ./internal
RUN go build -o /go/bin/gateway cmd/gateway/main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /go/bin/gateway /gateway
ENTRYPOINT ./gateway
EXPOSE 8080
