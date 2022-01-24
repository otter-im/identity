package app

import (
	"context"
	"errors"
	"github.com/otter-im/identity-backend/internal/rpc"
	"golang.org/x/crypto/bcrypt"
)

type LookupService struct {
	rpc.UnimplementedLookupServiceServer
	mainCtx context.Context
}

func (s *LookupService) Authorize(ctx context.Context, request *rpc.AuthorizationRequest) (*rpc.AuthorizationResponse, error) {
	select {
	case <-s.mainCtx.Done():
		return nil, errors.New("server stopping")
	case <-ctx.Done():
		return nil, errors.New("request cancelled")
	default:
	}

	user, err := SelectUserByUsername(ctx, request.GetUsername())
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(request.Password))
	if err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return &rpc.AuthorizationResponse{
				Status: rpc.AuthorizationResponse_FAIL_MISMATCHED_PASSWORD,
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
