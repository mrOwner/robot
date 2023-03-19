package grpc

import (
	"context"
	"net/http"

	"github.com/bufbuild/connect-go"
	"github.com/google/uuid"
	"github.com/mrOwner/robot/clients/grpc/proto/api"
	ac "github.com/mrOwner/robot/clients/grpc/proto/api/apiconnect"
	"github.com/mrOwner/robot/util"
)

type Client struct {
	token       string
	Instruments ac.InstrumentsServiceClient
}

func New(url, token string) *Client {
	return &Client{
		token:       token,
		Instruments: ac.NewInstrumentsServiceClient(http.DefaultClient, url, connect.WithGRPC()),
	}
}

func (c *Client) ShareByUID(ctx context.Context, uid uuid.UUID) (*api.Share, error) {
	req := connect.NewRequest(&api.InstrumentRequest{
		IdType: api.InstrumentIdType_INSTRUMENT_ID_TYPE_UID,
		Id:     uid.String(),
	})

	req.Header().Set(util.BearerToken(c.token))

	resp, err := c.Instruments.ShareBy(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp.Msg.Instrument, nil
}

func (c *Client) Share(ctx context.Context) ([]*api.Share, error) {
	req := connect.NewRequest(&api.InstrumentsRequest{
		InstrumentStatus: api.InstrumentStatus_INSTRUMENT_STATUS_ALL,
	})

	req.Header().Set(util.BearerToken(c.token))

	resp, err := c.Instruments.Shares(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp.Msg.GetInstruments(), nil
}
