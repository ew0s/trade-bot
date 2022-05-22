package handler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/ew0s/trade-bot/pkg/constant"
	"github.com/ew0s/trade-bot/pkg/httputils/baseresponse"
	"github.com/ew0s/trade-bot/pkg/httputils/request"
)

type UserIdentityKey string

var userUID UserIdentityKey = "user_uid"

type UserIdentityService interface {
	GetUserUID(accessToken string) (string, error)
}

type UserIdentity struct {
	service UserIdentityService
}

func NewUserIdentity(service UserIdentityService) *UserIdentity {
	return &UserIdentity{
		service: service,
	}
}

func (h *UserIdentity) identify(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bearerToken, err := request.GetBearerToken(r)
		if err != nil {
			baseresponse.RenderErr(w, r, fmt.Errorf("%w: %s", constant.ErrBadRequest, err))
			return
		}

		uid, err := h.service.GetUserUID(bearerToken)
		if err != nil {
			baseresponse.RenderErr(w, r, fmt.Errorf("%w: %s", constant.ErrUnauthorized, err))
			return
		}

		ctx := context.WithValue(r.Context(), userUID, uid)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
