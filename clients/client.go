//go:generate go run github.com/ogen-go/ogen/cmd/ogen --target http/ogen -package ogen --clean ../schema/openapi.yaml
package clients

import (
	"context"

	"github.com/google/uuid"
	"github.com/mrOwner/robot/clients/grpc/proto/api"
)

type Client interface {
	ShareByUID(ctx context.Context, uid uuid.UUID) (*api.Share, error)
	Share(ctx context.Context) ([]*api.Share, error)
}
