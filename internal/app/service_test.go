package app

import (
	"context"
	"encoding/hex"
	"errors"
	"github.com/go-pg/pg/v10"
	"github.com/google/uuid"
	"github.com/otter-im/identity/pkg/rpc"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

func TestLookupService_Authorize_ValidLogin(t *testing.T) {
	t.Parallel()

	dt := []struct {
		name string

		shouldCancel bool

		inputUsername string
		inputPassword string

		expectedResult *rpc.AuthorizationResponse
		expectedError  error
	}{
		{
			name: "valid login",

			inputUsername: "tinyfluffs",
			inputPassword: "changeme",

			expectedResult: &rpc.AuthorizationResponse{
				Id:       uuidBytes("d6ef6dc7-ce36-449c-8265-07f60ca3b2ff"),
				Username: "tinyfluffs",
			},
			expectedError: nil,
		},
		{
			name: "wrong password",

			inputUsername: "tinyfluffs",
			inputPassword: "here be dragons",

			expectedResult: nil,
			expectedError:  errors.New("password mismatch"),
		},
		{
			name: "unknown user",

			inputUsername: "alice",
			inputPassword: "hunter2",

			expectedResult: nil,
			expectedError:  errors.New("unknown user"),
		},
		{
			name: "canceled context",

			shouldCancel: true,

			inputUsername: "tinyfluffs",

			expectedResult: nil,
			expectedError:  errors.New("context canceled"),
		},
	}

	for _, test := range dt {
		t.Run(test.name, func(t *testing.T) {
			// Given
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			hash, _ := hex.DecodeString("70b7a5a7dab303fbf3880c5f75943b53bb1e2818ba9d2330555a78d30e69afd1")
			salt, _ := hex.DecodeString("2780cb19d7f864d49179ffb725284fa0")

			userService := new(MockUserService)
			userService.On("SelectUserByUsername", ctx, "alice").Return(nil, pg.ErrNoRows)
			userService.On("SelectUserByUsername", ctx, "tinyfluffs").Return(&User{
				Id:           uuid.MustParse("d6ef6dc7-ce36-449c-8265-07f60ca3b2ff"),
				Username:     "tinyfluffs",
				Hash:         hash,
				Salt:         salt,
				CreationDate: time.Now(),
			}, nil)

			service := &LookupService{
				Users: userService,
			}

			if test.shouldCancel {
				cancel()
			}

			// When
			actual, actualErr := service.Authorize(ctx, &rpc.AuthorizationRequest{
				Username: test.inputUsername,
				Password: test.inputPassword,
			})

			// Then
			if test.expectedError == nil {
				assert.Nil(t, actualErr)
			} else {
				assert.NotNil(t, actualErr)
				assert.Equal(t, test.expectedError, actualErr)
			}

			if test.expectedResult != nil {
				assert.NotNil(t, actual)
				assert.Equal(t, test.expectedResult, actual)
			} else {
				assert.Nil(t, actual)
			}
		})
	}
}

type MockUserService struct {
	mock.Mock
}

func (m *MockUserService) SelectUser(ctx context.Context, id uuid.UUID) (*User, error) {
	args := m.Called(ctx, id)
	result := args.Get(0)
	if result == nil {
		return nil, args.Error(1)
	}
	return result.(*User), args.Error(1)
}

func (m *MockUserService) SelectUserByUsername(ctx context.Context, username string) (*User, error) {
	args := m.Called(ctx, username)
	result := args.Get(0)
	if result == nil {
		return nil, args.Error(1)
	}
	return result.(*User), args.Error(1)
}

func uuidBytes(id string) []byte {
	result := uuid.MustParse(id)
	return result[:]
}
