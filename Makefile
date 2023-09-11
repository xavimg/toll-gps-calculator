obu:
	@go build -o bin/obu obu/main.go
	@./bin/obu

receiver:
	@go build -o bin/receiver ./data_receiver
	@./bin/receiver

calculator:
	@go build -o bin/calculator ./distance_calculator
	@./bin/calculator

aggregator:
	@go build -o bin/aggregator ./aggregator
	@./bin/aggregator

kafka:
	docker run --name kafka -p 9092 -e ALLOW_PLAINTEXT_LISTENER=yes -e KAFKA_CFG_AUTO_CREATE_TOPICS_ENABLE=true bitnami/kafka:latest

.PHONY: obu receiver calculator aggregator kafka