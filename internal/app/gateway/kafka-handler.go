package gateway

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/segmentio/kafka-go"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	kafkaAddress = "ql-kafka-bootstrap:9092"
	kafkaTopic   = "quests"
)

type KafkaHandler interface {
	Create(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	Close()
}

type handler struct {
	w *kafka.Writer
}

func (h handler) Create(w http.ResponseWriter, r *http.Request) {
	questJson, _ := ioutil.ReadAll(r.Body)
	err := h.w.WriteMessages(context.Background(),
		kafka.Message{Key: []byte("create"), Value: questJson},
	)
	if err != nil {
		log.Printf("Kafka handler: %v", err)
	}
}

func (h handler) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	err := h.w.WriteMessages(context.Background(),
		kafka.Message{Key: []byte("delete"), Value: []byte(vars["id"])},
	)
	if err != nil {
		log.Printf("Kafka handler: %v", err)
	}
}

func (h handler) Close() {
	_ = h.w.Close()
}

func NewKafkaHandler() KafkaHandler {
	h := handler{}

	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{kafkaAddress},
		Topic:   kafkaTopic,
	})

	h.w = w

	return h
}
