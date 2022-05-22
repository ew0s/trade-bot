package mapper

import (
	"testing"

	"github.com/ew0s/trade-bot/internal/api/request"
	"github.com/ew0s/trade-bot/internal/api/response"
	"github.com/ew0s/trade-bot/internal/domain/entities"
	"github.com/stretchr/testify/require"
)

func TestAuth_MakeUser(t *testing.T) {
	tests := []struct {
		name         string
		req          request.SignUp
		passwordHash string
		expected     entities.User
	}{
		{
			name: "successfully make user",
			req: request.SignUp{
				Name:     "user",
				Username: "username",
				Password: "password",
			},
			passwordHash: "$2a$04$ERbu2c6cXPXS0Higkq9s6uomHah2nAXj2fX0WhoQjsjYS.YrrGmw6",
			expected: entities.User{
				UID:          "",
				Name:         "user",
				Username:     "username",
				PasswordHash: "$2a$04$ERbu2c6cXPXS0Higkq9s6uomHah2nAXj2fX0WhoQjsjYS.YrrGmw6",
			},
		},
	}

	for _, tc := range tests {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			m := Auth{}

			actual := m.MakeUser(tc.req, tc.passwordHash)

			require.Equal(t, tc.expected, actual)
		})
	}
}

func TestAuth_MakeSignUpResponse(t *testing.T) {
	tests := []struct {
		name     string
		uid      string
		expected response.SignUp
	}{
		{
			name: "successfully make sign up response",
			uid:  "123e4567-e89b-12d3-a456-426655440000",
			expected: response.SignUp{
				UID: "123e4567-e89b-12d3-a456-426655440000",
			},
		},
	}

	for _, tc := range tests {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			m := Auth{}

			actual := m.MakeSignUpResponse(tc.uid)

			require.Equal(t, tc.expected, actual)
		})
	}
}

func TestAuth_MakeSignInResponse(t *testing.T) {
	tests := []struct {
		name     string
		token    string
		expected response.SignIn
	}{
		{
			name:  "successfully make sign in response",
			token: "123e4567-e89b-12d3-a456-426655440000",
			expected: response.SignIn{
				AccessToken: "123e4567-e89b-12d3-a456-426655440000",
			},
		},
	}

	for _, tc := range tests {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			m := Auth{}

			actual := m.MakeSignInResponse(tc.token)

			require.Equal(t, tc.expected, actual)
		})
	}
}
