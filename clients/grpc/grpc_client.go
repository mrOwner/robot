package grpc

import (
	"context"
	"net/http"

	ic "github.com/mrOwner/robot/clients/grpc/gen/investapiconnect"
)

type Client struct {
	Instruments ic.InstrumentsServiceClient
}

func New(url string) *Client {
	return &Client{
		Instruments: ic.NewInstrumentsServiceClient(http.DefaultClient, url),
	}
}

func (c *Client) ShareBy(ctx context.Context) {
	// c.Instruments.ShareBy(ctx, connect.NewRequest(&ic.Instr))
}
