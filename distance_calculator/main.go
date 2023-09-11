package main

import (
	"log"
)

const kafkaTopic = "obudata"

// type DistanceCalculator struct {
// 	consumer DataConsumer
// }

// transport (HTTP, gRPC, Kafka) -> attach bussines logic to this transport

func main() {
	svc := NewCalculatorService()
	svc = NewLogMiddleware(svc)
	KafkaConsumer, err := NewKafkaConsumer(kafkaTopic, svc)
	if err != nil {
		log.Fatal(err)
	}

	KafkaConsumer.Start()
}
