package gateway

import (
	"context"
	"encoding/json"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/gorilla/mux"
	"google.golang.org/grpc"
	"log"
	"net/http"
	pb "quest-log/internal/pkg/protos"
	"strconv"
	"time"
)

const grpcSvcAddress = "grpc-service:50051"

type GrpcHandler interface {
	Create(w http.ResponseWriter, r *http.Request)
	GetAll(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	Close()
}

type grpcHandler struct {
	conn   *grpc.ClientConn
	client pb.QuestServiceClient
}

func (h grpcHandler) Create(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var quest pb.Quest

	err := decoder.Decode(&quest)
	if err != nil {
		log.Printf("gRPC handler: %v", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	status, err := h.client.CreateQuest(ctx, &quest)
	if err != nil {
		handleGrpcError(err, w)
		return
	}

	_, err = w.Write([]byte(status.Message))
}

func (h grpcHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	questList, err := h.client.ListQuests(ctx, &empty.Empty{})
	if err != nil {
		handleGrpcError(err, w)
		return
	}

	b, err := json.Marshal(questList.Quests)
	if err != nil {
		handleGrpcError(err, w)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(b)
}

func (h grpcHandler) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		log.Printf("gRPC handler: ID must be an int")
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	quest := pb.Quest{Id: id}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	status, err := h.client.DeleteQuest(ctx, &quest)
	if err != nil {
		handleGrpcError(err, w)
	}

	_, err = w.Write([]byte(status.Message))
}

func (h grpcHandler) Close() {
	_ = h.conn.Close()
}

func handleGrpcError(err error, w http.ResponseWriter) {
	log.Printf("gRPC handler: %v", err)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func NewGrpcHandler() GrpcHandler {
	h := grpcHandler{}

	conn, err := grpc.Dial(grpcSvcAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("gRPC handler: %v", err)
	}

	h.conn = conn
	h.client = pb.NewQuestServiceClient(h.conn)

	return h
}
