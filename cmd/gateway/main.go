package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"quest-log/internal/app/gateway"
)

var grpcHandler gateway.GrpcHandler
var kafkaHandler gateway.KafkaHandler

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!\n"))
}

func main() {
	// gRPC handler
	grpcHandler = gateway.NewGrpcHandler()
	defer grpcHandler.Close()

	// Kafka handler
	kafkaHandler = gateway.NewKafkaHandler()
	defer kafkaHandler.Close()

	// Router
	router := mux.NewRouter()
	router.HandleFunc("/", hello)
	router.HandleFunc("/grpc/quests", grpcHandler.Create).Methods("POST")
	router.HandleFunc("/grpc/quests", grpcHandler.GetAll).Methods("GET")
	router.HandleFunc("/grpc/quests/{id:[0-9]+}", grpcHandler.Delete).Methods("DELETE")
	router.HandleFunc("/http/quests", gateway.HttpCreate).Methods("POST")
	router.HandleFunc("/http/quests", gateway.HttpGetAll).Methods("GET")
	router.HandleFunc("/http/quests/{id:[0-9]+}", gateway.HttpDelete).Methods("DELETE")
	router.HandleFunc("/kafka/quests", kafkaHandler.Create).Methods("POST")
	router.HandleFunc("/kafka/quests/{id:[0-9]+}", kafkaHandler.Delete).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}
