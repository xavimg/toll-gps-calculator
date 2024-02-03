package main

import (
	"log"
	"toll-calculator/aggregator/client"
)

const (
	kafkaTopic         = "obudata"
	aggregatorEndpoint = "http://127.0.0.1:3005/aggregate"
)

func main() {
	calculatorSvc := NewCalculatorService()
	calculatorSvc = NewLogMiddleware(calculatorSvc)

	httpClient := client.NewHTTPClient(aggregatorEndpoint)
	// grpc, err := client.NewGRPCClient(aggregatorEndpoint)
	// if err != nil {
	// 	log.Fatal()
	// }

	kafkaConsumer, err := NewKafkaConsumer(kafkaTopic, calculatorSvc, httpClient)
	if err != nil {
		log.Fatal(err)
	}

	kafkaConsumer.Start()
}
