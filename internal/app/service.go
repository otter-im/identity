package app

import (
	"bytes"
	"context"
	"errors"
	"github.com/go-pg/pg/v10"
	"github.com/otter-im/identity/pkg/rpc"
	"golang.org/x/crypto/argon2"
	"runtime"
)

const (
	argonTime      = 2
	argonMemory    = 64 * 1024
	argonKeyLength = 32
)

var (
	argonThreads = uint8(runtime.NumCPU())
)

type LookupService struct {
	rpc.UnimplementedLookupServiceServer
}

func (s *LookupService) Authorize(ctx context.Context, request *rpc.AuthorizationRequest) (*rpc.AuthorizationResponse, error) {
	user, err := SelectUserByUsername(ctx, request.GetUsername())
	if err != nil {
		if err == pg.ErrNoRows {
			return &rpc.AuthorizationResponse{
				Status: rpc.AuthorizationResponse_FAIL,
			}, nil
		}
		return nil, err
	}

	select {
	case <-ctx.Done():
		return nil, errors.New("request cancelled")
	default:
	}

	passwordKey := argon2.IDKey([]byte(request.Password), user.Salt[:], argonTime, argonMemory, argonThreads, argonKeyLength)
	if bytes.Compare(passwordKey, user.Hash[:]) != 0 {
		return &rpc.AuthorizationResponse{
			Status: rpc.AuthorizationResponse_FAIL,
		}, nil
	}

	return &rpc.AuthorizationResponse{
		Status:   rpc.AuthorizationResponse_SUCCESS,
		Id:       user.Id[:],
		Username: &user.Username,
	}, nil
}
