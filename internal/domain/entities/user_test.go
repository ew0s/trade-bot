package entities

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUser_GeneratePasswordHash_Success(t *testing.T) {
	tests := []struct {
		name            string
		password        string
		expectedHashLen int
	}{
		{
			name:            "successfully generate password hash",
			password:        "abc",
			expectedHashLen: 60,
		},
	}

	for _, tc := range tests {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			user := User{}

			actual, err := user.GeneratePasswordHash(tc.password)

			require.NoError(t, err)
			require.Equal(t, tc.expectedHashLen, len(actual))
		})
	}
}

func TestUser_ValidPassword(t *testing.T) {
	tests := []struct {
		name         string
		passwordHash string
		password     string
		expected     bool
	}{
		{
			name:         "successfully validate password",
			passwordHash: "$2a$04$ZUv1MyDmv3NCXykjuMruAuos8t6sKI/1/bzPIH4QC9jow6yU9LEZm",
			password:     "abc",
			expected:     true,
		},
		{
			name:         "invalid password passed",
			passwordHash: "$2a$04$ZUv1MyDmv3NCXykjuMruAuos8t6sKI/1/bzPIH4QC9jow6yU9LEZm",
			password:     "cab",
			expected:     false,
		},
	}

	for _, tc := range tests {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			user := User{
				PasswordHash: tc.passwordHash,
			}

			actual := user.ValidPassword(tc.password)

			require.Equal(t, tc.expected, actual)
		})
	}
}
