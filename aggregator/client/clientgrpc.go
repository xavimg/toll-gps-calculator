package client

import (
	"context"
	"toll-calculator/types"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ClientGRPC struct {
	Endpoint string
	client   types.AggregatorClient
}

func NewGRPCClient(endpoint string) (*ClientGRPC, error) {
	conn, err := grpc.Dial(endpoint, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	c := types.NewAggregatorClient(conn)
	return &ClientGRPC{
		Endpoint: endpoint,
		client:   c,
	}, nil
}

func (c *ClientGRPC) Aggregate(ctx context.Context, req *types.AggregateRequest) error {
	_, err := c.client.Aggregate(ctx, req)
	return err
}
