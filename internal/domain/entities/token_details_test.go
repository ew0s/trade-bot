package entities

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

const uuidV4Len = 36

func TestNewTokenDetails(t *testing.T) {
	tests := []struct {
		name     string
		userUID  string
		duration time.Duration
		expected TokenDetails
	}{
		{
			name:     "successfully make token details",
			userUID:  "6e00472a-0128-4adf-87a5-b4fc9a9f2bae",
			duration: 1 * time.Minute,
			expected: TokenDetails{
				UserUID:   "6e00472a-0128-4adf-87a5-b4fc9a9f2bae",
				ExpiresAt: time.Now().Add(1 * time.Minute),
			},
		},
	}

	for _, tc := range tests {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			actual, err := NewTokenDetails(tc.userUID, tc.duration)

			require.NoError(t, err)

			require.WithinDuration(t, tc.expected.ExpiresAt, actual.ExpiresAt, time.Second*3)

			require.Equal(t, uuidV4Len, len(actual.AccessUUID))
		})
	}
}
