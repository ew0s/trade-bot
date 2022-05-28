package mapper

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ew0s/trade-bot/internal/domain/entities"
	"github.com/ew0s/trade-bot/internal/repos/postgres/models"
)

func TestMakeEntityUser(t *testing.T) {
	tests := []struct {
		name     string
		model    models.User
		expected entities.User
	}{
		{
			name: "successfully make entity user",
			model: models.User{
				UID:          "123e4567-e89b-12d3-a456-426614174000",
				Name:         "Joe",
				Username:     "admin",
				PasswordHash: "$2a$04$ZUv1MyDmv3NCXykjuMruAuos8t6sKI/1/bzPIH4QC9jow6yU9LEZm",
			},
			expected: entities.User{
				UID:          "123e4567-e89b-12d3-a456-426614174000",
				Name:         "Joe",
				Username:     "admin",
				PasswordHash: "$2a$04$ZUv1MyDmv3NCXykjuMruAuos8t6sKI/1/bzPIH4QC9jow6yU9LEZm",
			},
		},
	}

	for _, tc := range tests {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			actual := MakeEntityUser(tc.model)

			require.Equal(t, tc.expected, actual)
		})
	}
}

func TestMakeModelUser(t *testing.T) {
	tests := []struct {
		name     string
		entity   entities.User
		expected models.User
	}{
		{
			name: "successfully make entity user",
			entity: entities.User{
				UID:          "123e4567-e89b-12d3-a456-426614174000",
				Name:         "Joe",
				Username:     "admin",
				PasswordHash: "$2a$04$ZUv1MyDmv3NCXykjuMruAuos8t6sKI/1/bzPIH4QC9jow6yU9LEZm",
			},
			expected: models.User{
				UID:          "123e4567-e89b-12d3-a456-426614174000",
				Name:         "Joe",
				Username:     "admin",
				PasswordHash: "$2a$04$ZUv1MyDmv3NCXykjuMruAuos8t6sKI/1/bzPIH4QC9jow6yU9LEZm",
			},
		},
	}

	for _, tc := range tests {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			actual := MakeModelUser(tc.entity)

			require.Equal(t, tc.expected, actual)
		})
	}
}
