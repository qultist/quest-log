package repository

import (
	"encoding/json"
	"fmt"
	"github.com/gomodule/redigo/redis"
	pb "quest-log/internal/pkg/protos"
)

const (
	redisAddress = "redis:6379"
)

type Repository interface {
	StoreQuest(quest *pb.Quest)
	FetchQuests() []*pb.Quest
	DeleteQuest(id int64) bool
	CloseConnection()
}

type repository struct {
	conn redis.Conn
}

func (repo repository) StoreQuest(quest *pb.Quest) {
	id, err := redis.Int64(repo.conn.Do("INCR", "id:quests"))
	if err != nil {
		panic(err)
	}

	quest.Id = id

	b, err := json.Marshal(quest)
	if err != nil {
		panic(err)
	}

	_, err = repo.conn.Do("SET", fmt.Sprintf("quest:%d", id), string(b))
	if err != nil {
		panic(err)
	}

	_, err = repo.conn.Do("SADD", "quests", id)
	if err != nil {
		panic(err)
	}
}

func (repo repository) FetchQuests() []*pb.Quest {
	ids, err := redis.Int64s(repo.conn.Do("SMEMBERS", "quests"))
	if err != nil {
		panic(err)
	}

	var quests []*pb.Quest

	for _, id := range ids {
		questJson, err := redis.String(repo.conn.Do("GET", fmt.Sprintf("quest:%d", id)))
		if err != nil {
			panic(err)
		}

		var quest pb.Quest
		err = json.Unmarshal([]byte(questJson), &quest)
		if err != nil {
			panic(err)
		}

		quests = append(quests, &quest)
	}

	return quests
}

func (repo repository) DeleteQuest(id int64) bool {
	deleted, err := redis.Int64(repo.conn.Do("DEL", fmt.Sprintf("quest:%d", id)))
	if err != nil {
		panic(err)
	}

	_, err = repo.conn.Do("SREM", "quests",id)
	if err != nil {
		panic(err)
	}

	return deleted == 1
}

func (repo repository) CloseConnection() {
	repo.conn.Close()
}

func NewRepository() Repository {
	repo := repository{}

	c, err := redis.Dial("tcp", redisAddress)
	if err != nil {
		panic(err)
	}

	repo.conn = c

	return repo
}
