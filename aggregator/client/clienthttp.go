package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"toll-calculator/types"
)

type ClientHTTP struct {
	Endpoint string
}

func NewClient(endpoint string) *ClientHTTP {
	return &ClientHTTP{
		Endpoint: endpoint,
	}
}

func (c *ClientHTTP) AggregateInvoice(dist types.Distance) error {
	b, err := json.Marshal(dist)
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
