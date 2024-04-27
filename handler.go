package yodel

import "github.com/labstack/echo/v4"

// Handler handles a given HTTP(S) request.
type Handler interface {
    Handle(echo.Context) error
}

// HandlerFunc provides a function implementation of Handler.
type HandlerFunc func(echo.Context) error

// Handle implements Handler.
func (fn HandlerFunc) Handle(c echo.Context) error {
    return fn(c)
}
