package handler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/ew0s/trade-bot/pkg/constant"
	"github.com/ew0s/trade-bot/pkg/httputils/baseresponse"
	"github.com/ew0s/trade-bot/pkg/httputils/request"
)

type UserIdentityService interface {
	ValidAccessToken(ctx context.Context, accessToken string) (bool, error)
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

		valid, err := h.service.ValidAccessToken(r.Context(), bearerToken)
		if err != nil {
			baseresponse.RenderErr(w, r, fmt.Errorf("%w: %s", constant.ErrUnauthorized, err))
			return
		}

		if !valid {
			baseresponse.RenderErr(w, r, fmt.Errorf(
				"%w: %s", constant.ErrUnauthorized, fmt.Errorf("invalid access token")),
			)
			return
		}

		next.ServeHTTP(w, r)
	})
}
