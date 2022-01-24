package app

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/otter-im/identity-service/pkg/rpc"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLookupService_Authorize(t *testing.T) {
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
			expectedError: errors.New("context cancelled"),
		},
	}

	for _, test := range dt {
		t.Run(test.name, func(t *testing.T) {
			// Given
			t.Parallel()
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

			ctx, cancel := context.WithCancel(context.Background())
			if test.cancelContext {
				cancel()
			}

			// When
			actual, actualErr := service.Authorize(ctx, &rpc.AuthorizationRequest{
				Username: test.inputUsername,
				Password: test.inputPassword,
			})

			// Then
			cancel()

			if test.expectedError == nil {
				assert.Nil(t, actualErr)
			} else {
				assert.NotNil(t, actualErr)
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
