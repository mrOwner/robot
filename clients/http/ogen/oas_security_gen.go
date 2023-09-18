// Code generated by ogen, DO NOT EDIT.

package ogen

import (
	"context"
	"net/http"
	"strings"

	"github.com/go-faster/errors"

	"github.com/ogen-go/ogen/ogenerrors"
)

// SecurityHandler is handler for security parameters.
type SecurityHandler interface {
	// HandleBearer handles Bearer security.
	HandleBearer(ctx context.Context, operationName string, t Bearer) (context.Context, error)
}

func findAuthorization(h http.Header, prefix string) (string, bool) {
	v, ok := h["Authorization"]
	if !ok {
		return "", false
	}
	for _, vv := range v {
		scheme, value, ok := strings.Cut(vv, " ")
		if !ok || !strings.EqualFold(scheme, prefix) {
			continue
		}
		return value, true
	}
	return "", false
}

func (s *Server) securityBearer(ctx context.Context, operationName string, req *http.Request) (context.Context, bool, error) {
	var t Bearer
	token, ok := findAuthorization(req.Header, "Bearer")
	if !ok {
		return ctx, false, nil
	}
	t.Token = token
	rctx, err := s.sec.HandleBearer(ctx, operationName, t)
	if err != nil {
		return nil, false, err
	}
	return rctx, true, err
}

// SecuritySource is provider of security values (tokens, passwords, etc.).
type SecuritySource interface {
	// Bearer provides Bearer security value.
	Bearer(ctx context.Context, operationName string) (Bearer, error)
}

func (s *Client) securityBearer(ctx context.Context, operationName string, req *http.Request) error {
	t, err := s.sec.Bearer(ctx, operationName)
	if err != nil {
		if errors.Is(err, ogenerrors.ErrSkipClientSecurity) {
			return ogenerrors.ErrSkipClientSecurity
		}
		return errors.Wrap(err, "security source \"Bearer\"")
	}
	req.Header.Set("Authorization", "Bearer "+t.Token)
	return nil
}
