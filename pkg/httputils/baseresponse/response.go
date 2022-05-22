package baseresponse

import (
	"encoding/json"
	"net/http"

	"github.com/ew0s/trade-bot/pkg/constant"
	"github.com/go-chi/render"
	"github.com/pkg/errors"
)

// ErrResponse renderer type for handling all sorts of errors.
//
// In the best case scenario, the excellent github.com/pkg/errors package
// helps reveal information on the error, setting it on Err, and in the Render()
// method, using it to set the application-specific error code in AppCode.
type ErrResponse struct {
	Err            error `json:"-"` // low-level runtime error
	HTTPStatusCode int   `json:"-"` // http response status code

	ErrorText string `json:"error,omitempty"` // application-level error message, for debugging
}

// Render error response
func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

// ErrUnknown describes an unknown error
func ErrUnknown(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: http.StatusInternalServerError,
		ErrorText:      err.Error(),
	}
}

// ErrNotFound describes an error not found
func ErrNotFound(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: http.StatusNotFound,
		ErrorText:      err.Error(),
	}
}

// ErrForbidden describes an unknown error forbidden
func ErrForbidden(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: http.StatusForbidden,
		ErrorText:      err.Error(),
	}
}

func ErrUnauthorized(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: http.StatusUnauthorized,
		ErrorText:      err.Error(),
	}
}

// ErrBadRequest describes an unknown error bad request
func ErrBadRequest(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: http.StatusBadRequest,
		ErrorText:      err.Error(),
	}
}

// ErrConflict describes an unknown error conflict
func ErrConflict(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: http.StatusConflict,
		ErrorText:      err.Error(),
	}
}

// ErrTooManyRequests describes exceeding rate limit requests error
func ErrTooManyRequests(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: http.StatusTooManyRequests,
		ErrorText:      err.Error(),
	}
}

// SuccessResponse describes an success work
type SuccessResponse struct {
	Result interface{} `json:"result"`
}

// Render success response
func (e *SuccessResponse) Render(w http.ResponseWriter, r *http.Request) error {
	resp, err := json.Marshal(e.Result)
	if err != nil {
		return errors.Wrap(err, "can't marshal response")
	}

	_, err = w.Write(resp)
	if err != nil {
		return errors.Wrap(err, "can't write response")
	}

	render.Status(r, http.StatusOK)

	return nil
}

func RenderErr(w http.ResponseWriter, r *http.Request, err error) {
	var respErr render.Renderer

	switch {
	case errors.Is(err, constant.ErrBadRequest):
		respErr = ErrBadRequest(err)

	case errors.Is(err, constant.ErrUnauthorized):
		respErr = ErrUnauthorized(err)

	case errors.Is(err, constant.ErrNotFound):
		respErr = ErrNotFound(err)

	default:
		respErr = ErrUnknown(err)
	}

	_ = render.Render(w, r, respErr)
}
