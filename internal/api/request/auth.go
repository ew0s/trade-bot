package request

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ew0s/trade-bot/pkg/httputils"
	"github.com/go-chi/chi/v5"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type SignIn struct {
	UID      string `json:"-"`
	Password string `json:"password"`
}

func (r *SignIn) Bind(req *http.Request) error {
	uid := chi.URLParam(req, "uid")

	if err := json.NewDecoder(req.Body).Decode(r); err != nil {
		return fmt.Errorf("decoding body: %w", err)
	}

	r.UID = uid

	if err := r.validate(); err != nil {
		return fmt.Errorf("validating SignIn: %w", err)
	}

	return nil
}

func (r *SignIn) validate() error {
	if err := validation.ValidateStruct(r,
		validation.Field(&r.UID, validation.Required),
		validation.Field(&r.Password, validation.Required),
	); err != nil {
		return fmt.Errorf("validating struct: %w", err)
	}

	return nil
}

type SignUp struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (r *SignUp) Bind(req *http.Request) error {
	if err := json.NewDecoder(req.Body).Decode(r); err != nil {
		return fmt.Errorf("decoding body: %w", err)
	}

	if err := r.validate(); err != nil {
		return fmt.Errorf("validating struct: %w", err)
	}

	return nil
}

func (r *SignUp) validate() error {
	if err := validation.ValidateStruct(r,
		validation.Field(&r.Name, validation.Required),
		validation.Field(&r.Username, validation.Required),
		validation.Field(&r.Password, validation.Required),
	); err != nil {
		return fmt.Errorf("validating struct: %w", err)
	}

	return nil
}

type Logout struct {
	AccessToken string
}

func (r *Logout) Bind(req *http.Request) error {
	token, err := httputils.GetBearerToken(req)
	if err != nil {
		return fmt.Errorf("getting bearer token: %w", err)
	}

	r.AccessToken = token

	return nil
}
