package handler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"

	"github.com/ew0s/trade-bot/internal/api/request"
	"github.com/ew0s/trade-bot/internal/api/response"
	"github.com/ew0s/trade-bot/pkg/constant"
	"github.com/ew0s/trade-bot/pkg/httputils/baseresponse"
)

type UserAuthService interface {
	SignUp(ctx context.Context, req request.SignUp) (response.SignUp, error)
	SignIn(ctx context.Context, req request.SignIn) (response.SignIn, error)
	Logout(ctx context.Context, req request.Logout) error
}

type Auth struct {
	service      UserAuthService
	userIdentity *UserIdentity
}

func NewAuth(service UserAuthService, identity *UserIdentity) *Auth {
	return &Auth{
		service:      service,
		userIdentity: identity,
	}
}

func (h *Auth) Routes() chi.Router {
	r := chi.NewRouter()

	r.Route("/auth", func(r chi.Router) {
		r.Post("/sign-in", h.SignIn)
		r.Post("/sign-up", h.SignUp)
		r.With(h.userIdentity.identify).Delete("/logout", h.Logout)
	})

	return r
}

// SignUp
// @Summary      sign up user
// @Tags         auth
// @Description  method stands for signing up user
// @ID           sign-up-user
// @Accept       json
// @Param        input  body  request.SignUp  true  "account info"
// @Success      201    "successfully create user"
// @Header       201    {string}  Location                  "user uid"
// @Failure      400    {object}  baseresponse.ErrResponse  "bad request"
// @Failure      500    {object}  baseresponse.ErrResponse  "internal server error"
// @Router       /auth/sign-up [post]
func (h *Auth) SignUp(w http.ResponseWriter, r *http.Request) {
	var req request.SignUp
	if err := req.Bind(r); err != nil {
		baseresponse.RenderErr(w, r, fmt.Errorf("%w: %s", constant.ErrBadRequest, err))
		return
	}

	resp, err := h.service.SignUp(r.Context(), req)
	if err != nil {
		baseresponse.RenderErr(w, r, err)
		return
	}

	w.Header().Add("Location", resp.UID)
	w.WriteHeader(http.StatusCreated)
}

// SignIn
// @Summary      sign in user
// @Tags         auth
// @Description  method stands for signing in user
// @ID           sign-in-user
// @Accept       json
// @Produce      json
// @Param        uid    query     string                    true  "user uid"
// @Param        input  body      request.SignIn            true  "credentials"
// @Success      200    {object}  response.SignIn           "access token"
// @Failure      400    {object}  baseresponse.ErrResponse  "bad request"
// @Failure      500    {object}  baseresponse.ErrResponse  "internal server error"
// @Router       /auth/sign-in [post]
func (h *Auth) SignIn(w http.ResponseWriter, r *http.Request) {
	var req request.SignIn
	if err := req.Bind(r); err != nil {
		baseresponse.RenderErr(w, r, fmt.Errorf("%w: %s", constant.ErrBadRequest, err))
		return
	}

	resp, err := h.service.SignIn(r.Context(), req)
	if err != nil {
		baseresponse.RenderErr(w, r, err)
		return
	}

	render.JSON(w, r, resp)
}

// Logout
// @Summary      logout user
// @Security     ApiKeyAuth
// @Tags         auth
// @Description  method stands for logging out user
// @ID           logout-user
// @Success      204  "ok"
// @Failure      400  {object}  baseresponse.ErrResponse  "bad request"
// @Failure      404  {object}  baseresponse.ErrResponse  "not found"
// @Failure      500  {object}  baseresponse.ErrResponse  "internal server error"
// @Router       /auth/logout [delete]
func (h *Auth) Logout(w http.ResponseWriter, r *http.Request) {
	var req request.Logout
	if err := req.Bind(r); err != nil {
		baseresponse.RenderErr(w, r, fmt.Errorf("%w: %s", constant.ErrUnauthorized, err))
		return
	}

	if err := h.service.Logout(r.Context(), req); err != nil {
		baseresponse.RenderErr(w, r, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
