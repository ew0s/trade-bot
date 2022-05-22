package mapper

import (
	"github.com/ew0s/trade-bot/internal/api/request"
	"github.com/ew0s/trade-bot/internal/api/response"
	"github.com/ew0s/trade-bot/internal/domain/entities"
)

type Auth struct{}

func (m Auth) MakeUser(req request.SignUp, passwordHash string) entities.User {
	return entities.User{
		Name:         req.Name,
		Username:     req.Username,
		PasswordHash: passwordHash,
	}
}

func (m Auth) MakeSignUpResponse(uid string) response.SignUp {
	return response.SignUp{
		UID: uid,
	}
}

func (m Auth) MakeSignInResponse(token string) response.SignIn {
	return response.SignIn{AccessToken: token}
}
