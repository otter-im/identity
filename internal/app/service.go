package app

import (
	"context"
	"errors"
	"github.com/go-pg/pg/v10"
	"github.com/otter-im/identity/pkg/rpc"
	"golang.org/x/crypto/bcrypt"
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

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(request.Password))
	if err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return &rpc.AuthorizationResponse{
				Status: rpc.AuthorizationResponse_FAIL,
			}, nil
		}
		return nil, err
	}

	return &rpc.AuthorizationResponse{
		Status:   rpc.AuthorizationResponse_SUCCESS,
		Id:       user.Id[:],
		Username: &user.Username,
	}, nil
}
