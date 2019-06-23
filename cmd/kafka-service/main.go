package main

import (
	"context"
	"encoding/json"
	"github.com/segmentio/kafka-go"
	"log"
	pb "quest-log/internal/pkg/protos"
	"quest-log/internal/pkg/repository"
	"strconv"
)

const (
	kafkaAddress = "kafka:9092"
	kafkaTopic   = "quests"
)

func main() {
	repo := repository.NewRepository()
	defer repo.Close()

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{kafkaAddress},
		Topic:   kafkaTopic,
	})

	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			log.Printf("Kafka service: %v", err)
			break
		}

		switch string(m.Key) {
		case "delete":
			s := string(m.Value)
			if id, err := strconv.ParseInt(s, 10, 64); err == nil {
				repo.DeleteQuest(id)
			}
		case "create":
			var quest pb.Quest
			if err = json.Unmarshal(m.Value, &quest); err != nil {
				log.Printf("Kafka Service: %v", err)
				break
			}
			quest.AdditionalInfo = "Created from Kafka service"
			repo.StoreQuest(&quest)
		default:
			break
		}
	}

	_ = r.Close()
}
