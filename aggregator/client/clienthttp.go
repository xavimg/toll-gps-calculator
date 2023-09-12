package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"toll-calculator/types"
)

type ClientHTTP struct {
	Endpoint string
}

func NewHTTPClient(endpoint string) *ClientHTTP {
	return &ClientHTTP{
		Endpoint: endpoint,
	}
}

func (c *ClientHTTP) Aggregate(ctx context.Context, r *types.AggregateRequest) error {
	b, err := json.Marshal(r)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", c.Endpoint, bytes.NewReader(b))
	if err != nil {
		return err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("service responded with a non 200 status code %d", resp.StatusCode)
	}

	return nil
}
