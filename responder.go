package yodel

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type (
    // Responder prepares the response using the provided model type.
	Responder[T any] interface {
		Respond(echo.Context, T) error
	}

    // ResponderFunc provides a functions implementation of Responder.
	ResponderFunc[T any] func(echo.Context, T) error

    // JSONResponder responds to the request by encoding the provided
    // model type into JSON before serializing it to the response body.
	JSONResponder[T any] struct{}

    // RedirectResponder responds to the request via redirect using the
    // parameters provided in the RedirectResponse value.
    RedirectResponder struct{}

    // RedirectResponse provides the HTTP status code and URL necessary
    // to perform the redirect.
	RedirectResponse struct {
		Code int
		Url  string
	}
)

// Respond implements Responder.
func (fn ResponderFunc[T]) Respond(c echo.Context, value T) error { return fn(c, value) }

// Respond implements Responder.
func (JSONResponder[T]) Respond(c echo.Context, value T) error {
	return c.JSON(http.StatusOK, value)
}

// Respond implements Responder.
func (RedirectResponder) Respond(c echo.Context, value RedirectResponse) error {
	return c.Redirect(value.Code, value.Url)
}
