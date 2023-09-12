package main

import (
	"context"
	"log"
	"time"
	"toll-calculator/aggregator/client"
	"toll-calculator/types"
)

func main() {
	c, err := client.NewGRPCClient(":3006")
	if err != nil {
		log.Fatal(err)
	}
	if err := c.Aggregate(context.Background(), &types.AggregateRequest{
		ObuID: 1,
		Value: 58.55,
		Unix:  time.Now().Unix(),
	}); err != nil {
		log.Fatal(err)
	}
}
