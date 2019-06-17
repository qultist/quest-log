package main

import (
	"context"
	"encoding/json"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"log"
	"net/http"
	"quest-log/internal/app/gateway"
	"strconv"
	"time"

	"github.com/gorilla/mux"

	pb "quest-log/internal/pkg/protos"
)

const (
	grpcSvcAddress = "grpc-service:50051"
)

var grpcServiceClient pb.QuestServiceClient

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!\n"))
}

func grpcCreate(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var quest pb.Quest

	err := decoder.Decode(&quest)
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	status, err := grpcServiceClient.CreateQuest(ctx, &quest)
	if err != nil {
		log.Fatalf("could not create: %v", err)
	}
	w.Write([]byte(status.Message))
}

func grpcList(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	questList, err := grpcServiceClient.ListQuests(ctx, &empty.Empty{})
	if err != nil {
		log.Fatalf("could not list: %v", err)
	}

	b, err := json.Marshal(questList.Quests)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}

func grpcDelete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		panic(err)
	}

	quest := pb.Quest{Id: id}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	status, err := grpcServiceClient.DeleteQuest(ctx, &quest)
	if err != nil {
		log.Fatalf("could not delete: %v", err)
	}

	w.Write([]byte(status.Message))
}

func main() {
	// gRPC client
	conn, err := grpc.Dial(grpcSvcAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	grpcServiceClient = pb.NewQuestServiceClient(conn)

	// Router
	router := mux.NewRouter()
	router.HandleFunc("/", hello)
	router.HandleFunc("/grpc/quests", grpcCreate).Methods("POST")
	router.HandleFunc("/grpc/quests", grpcList).Methods("GET")
	router.HandleFunc("/grpc/quests/{id:[0-9]+}", grpcDelete).Methods("DELETE")
	router.HandleFunc("/http/quests", gateway.HttpCreate).Methods("POST")
	router.HandleFunc("/http/quests", gateway.HttpGetAll).Methods("GET")
	router.HandleFunc("/http/quests/{id:[0-9]+}", gateway.HttpDelete).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}
