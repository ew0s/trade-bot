package security

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGeneratePasswordHash(t *testing.T) {
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
			actual, err := GeneratePasswordHash(tc.password)

			require.NoError(t, err)
			require.Equal(t, tc.expectedHashLen, len(actual))
		})
	}
}
