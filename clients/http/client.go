package http

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/mrOwner/robot/clients/http/ogen"
)

type Client struct {
	*ogen.Client
}

type Bear struct {
	token string
}

func New(url, token string) (*Client, error) {
	var (
		c   Client
		err error
	)

	c.Client, err = ogen.NewClient(url, &Bear{token: token})
	if err != nil {
		return nil, err
	}

	return &c, nil
}

func (b *Bear) Bearer(ctx context.Context, _ string) (ogen.Bearer, error) {
	return ogen.Bearer{Token: b.token}, nil
}

func (c *Client) ShareByUID(ctx context.Context, uid uuid.UUID) error {
	req := &ogen.V1InstrumentRequest{}
	req.IdType.SetTo(ogen.V1InstrumentIdTypeINSTRUMENTIDTYPEPOSITIONUID)
	req.ID.SetTo(uid.String())

	res, err := c.Client.InstrumentsServiceShareBy(ctx, req)
	if err != nil {
		return err
	}

	switch result := res.(type) {
	case *ogen.RpcStatusStatusCode:
		return errors.New(result.Response.Message.Value)
	case *ogen.V1ShareResponse:
	}

	return nil
}
