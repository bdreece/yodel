package yodel

import "github.com/labstack/echo/v4"

type (
	// A Binder binds values from the request to the generic
    // model type.
	Binder[T any] interface {
		Bind(echo.Context) (T, error)
	}

	// BinderFunc provides a function implementation of Binder.
	BinderFunc[T any] func(echo.Context) (T, error)

	// DefaultBinder binds the model type using the built-in
    // [echo.DefaultBinder].
	DefaultBinder[T any] struct{}
)

// Bind implements Binder.
func (fn BinderFunc[T]) Bind(c echo.Context) (T, error) { return fn(c) }

// Bind implements Binder.
func (DefaultBinder[T]) Bind(c echo.Context) (value T, err error) {
	err = c.Bind(value)
	return
}
