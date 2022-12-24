package consumer

import (
	"fmt"
	"strings"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)




func Consume(topics []string, server string, msgChan chan *kafka.Message){
	kafkaConsumer, err := kafka.NewConsumer (&kafka.ConfigMap{
		"bootstrap.servers": servers,
        "group.id":          "gostats",
        "auto.offset.reset": "earliest",
	})
	if err != nil{
		panic(err)
	}
	kafkaConsumer.SubscribeTopics(topics, nil)
	for msg := range msgChan {
		fmt.Println( "Received message", string(msg.Value),"on topic", *msg.TopipcPartition.Topic)
		strategy := factory
		
	}
}  

