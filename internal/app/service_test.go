package app

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/otter-im/identity/pkg/rpc"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLookupService_Authorize(t *testing.T) {
	t.Parallel()

	dt := []struct {
		name string

		cancelContext bool

		inputUsername string
		inputPassword string

		expectResult     bool
		expectedStatus   rpc.AuthorizationResponse_Status
		expectedId       string
		expectedUsername string
		expectedError    error
	}{
		{
			name: "valid login",

			cancelContext: false,

			inputUsername: "tinyfluffs",
			inputPassword: "changeme",

			expectResult:     true,
			expectedStatus:   rpc.AuthorizationResponse_SUCCESS,
			expectedId:       "d6ef6dc7-ce36-449c-8265-07f60ca3b2ff",
			expectedUsername: "tinyfluffs",
			expectedError:    nil,
		},
		{
			name: "wrong password",

			cancelContext: false,

			inputUsername: "tinyfluffs",
			inputPassword: "here be dragons",

			expectResult:     true,
			expectedStatus:   rpc.AuthorizationResponse_FAIL,
			expectedUsername: "",
			expectedId:       "",
			expectedError:    nil,
		},
		{
			name: "unknown user",

			cancelContext: false,

			inputUsername: "alice",
			inputPassword: "hunter2",

			expectResult:     true,
			expectedStatus:   rpc.AuthorizationResponse_FAIL,
			expectedUsername: "",
			expectedId:       "",
			expectedError:    nil,
		},
		{
			name: "cancelled context",

			cancelContext: true,
			expectedError: errors.New("context canceled"),
		},
	}

	for _, test := range dt {
		t.Run(test.name, func(t *testing.T) {
			// Given
			service := &LookupService{}

			var expectedId []byte
			var expectedUsername *string

			if test.expectedId == "" {
				expectedId = nil
			} else {
				id := uuid.MustParse(test.expectedId)
				expectedId = id[:]
			}

			if test.expectedUsername == "" {
				expectedUsername = nil
			} else {
				expectedUsername = &test.expectedUsername
			}

			expected := &rpc.AuthorizationResponse{
				Status:   test.expectedStatus,
				Id:       expectedId[:],
				Username: expectedUsername,
			}

			ctx := context.Background()
			if test.cancelContext {
				ctx2, cancel := context.WithCancel(ctx)
				ctx = ctx2
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

			if test.expectResult {
				assert.NotNil(t, actual)
				assert.Equal(t, expected, actual)
			} else {
				assert.Nil(t, actual)
			}
		})
	}
}
