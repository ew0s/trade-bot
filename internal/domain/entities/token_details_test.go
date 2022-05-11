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
				AtExpires: time.Now().Add(1 * time.Minute).Unix(),
			},
		},
	}

	for _, tc := range tests {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			actual, err := NewTokenDetails(tc.userUID, tc.duration)

			require.NoError(t, err)

			require.GreaterOrEqual(t, actual.AtExpires, tc.expected.AtExpires)
			require.InDelta(t, tc.expected.AtExpires, actual.AtExpires, time.Second.Seconds())

			require.Equal(t, uuidV4Len, len(actual.AccessUUID))
		})
	}
}
