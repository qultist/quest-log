package main

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"log"
	"net"
	pb "quest-log/internal/pkg/protos"
	"quest-log/internal/pkg/repository"
)

const (
	port = ":50051"
)

var repo repository.Repository

type server struct{}

func (s *server) ListQuests(context.Context, *empty.Empty) (*pb.QuestList, error) {
	quests := repo.FetchQuests()
	return &pb.QuestList{Quests: quests}, nil
}

func (s *server) CreateQuest(ctx context.Context, quest *pb.Quest) (*pb.Status, error) {
	repo.StoreQuest(quest)
	return &pb.Status{Message: fmt.Sprintf("Created: \"%s\", id: %d", quest.Title, quest.Id)}, nil
}

func (s *server) DeleteQuest(ctx context.Context, quest *pb.Quest) (*pb.Status, error) {
	deleted := repo.DeleteQuest(quest.Id)

	if deleted {
		return &pb.Status{Message: fmt.Sprintf("Successfully deleted quest with id %d", quest.Id)}, nil
	} else {
		return &pb.Status{Message: fmt.Sprintf("There is no quest with id %d", quest.Id)}, nil
	}
}

func main() {
	repo = repository.NewRepository()
	defer repo.Close()
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterQuestServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
