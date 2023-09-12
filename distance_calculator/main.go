package main

import (
	"log"
	"toll-calculator/aggregator/client"
)

const (
	kafkaTopic         = "obudata"
	aggregatorEndpoint = "http://127.0.0.1:3005/aggregate"
)

// type DistanceCalculator struct {
// 	consumer DataConsumer
// }

// transport (HTTP, gRPC, Kafka) -> attach bussines logic to this transport

func main() {
	svc := NewCalculatorService()
	svc = NewLogMiddleware(svc)

	KafkaConsumer, err := NewKafkaConsumer(kafkaTopic, svc, client.NewClient(aggregatorEndpoint))
	if err != nil {
		log.Fatal(err)
	}

	KafkaConsumer.Start()
}
