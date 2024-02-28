package main

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/kuroshibaz/config"
	"log"
)

var topics = []string{"SUMMARY:REPORT"}

func main() {
	cfg := loadConfig()

	consumer, kafErr := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": cfg.Kafka.Hostname,
		"group.id":          "",
		"auto.offset.reset": "smallest",
	})
	if kafErr != nil {
		log.Fatalf("kafka connection failed: %v", kafErr)
	}
	defer consumer.Close()

	_ = consumer.SubscribeTopics(topics, nil)

	//var running = true
	//go func() {
	//}()
	for {
		e := consumer.Poll(100)
		log.Printf("event: %v", e)
		switch ev := e.(type) {
		case *kafka.Message:
			{
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Failed to deliver message %v\n", ev.TopicPartition)
				}
			}
		default:
			log.Println("Ignore message")
		}
	}

	consumer.Close()
}

/* loadConfig read/map to environment config */
func loadConfig() *config.Env {
	cfg, err := config.ReadConfig("config")
	if err != nil {
		log.Fatalf("error config: %v", err)
	}

	return cfg
}
