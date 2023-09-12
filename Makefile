obu:
	@go build -o bin/obu obu/main.go
	@./bin/obu

consumer:
	@go build -o bin/consumer ./distance_producer
	@./bin/consumer

producer:
	@go build -o bin/producer ./distance_consumer
	@./bin/producer

aggregator:
	@go build -o bin/aggregator ./aggregator
	@./bin/aggregator

kafka:
	docker run --name kafka -p 9092 -e ALLOW_PLAINTEXT_LISTENER=yes -e KAFKA_CFG_AUTO_CREATE_TOPICS_ENABLE=true bitnami/kafka:latest

proto:
	protoc --go_out=. --go_opt=paths=source_relative  --go-grpc_out=. --go-grpc_opt=paths=source_relative types/ptypes.proto

.PHONY: obu receiver calculator aggregator kafka