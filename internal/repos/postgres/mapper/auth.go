package mapper

import (
	"github.com/ew0s/trade-bot/internal/domain/entities"
	"github.com/ew0s/trade-bot/internal/repos/postgres/models"
)

func MakeModelUser(e entities.User) models.User {
	return models.User{
		UID:          e.UID,
		Name:         e.Name,
		Username:     e.Username,
		PasswordHash: e.PasswordHash,
	}
}

func MakeEntityUser(m models.User) entities.User {
	return entities.User{
		UID:          m.UID,
		Name:         m.Name,
		Username:     m.Username,
		PasswordHash: m.PasswordHash,
	}
}
