package main

import (
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/dougecomp/kafka-go-example/application/kafka"
	"github.com/joho/godotenv"
	"path/filepath"
	"runtime"
)

func init() {
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)
	godotenv.Load(basepath + "/../.env")
}

func main() {
	deliveryChan := make(chan ckafka.Event)
	producer := kafka.NewKafkaProducer()

	kafka.Publish(producer, deliveryChan)

	kafka.DeliveryReport(deliveryChan)
}
