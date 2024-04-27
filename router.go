package yodel

import (
	"log/slog"

	"github.com/labstack/echo/v4"
	slogecho "github.com/samber/slog-echo"
)

// Router provides a thin wrapper around [echo.Echo]
type Router struct{ *echo.Echo }

// Add registers a new Route.
func (r Router) Add(route Route) *echo.Route {
	return r.Echo.Add(route.Method, route.Path, route.Handler.Handle, mapMiddlewares(route.Middleware)...)
}

// Use add a Middleware to the chain.
func (r Router) Use(middleware ...Middleware) {
    r.Echo.Use(mapMiddlewares(middleware)...)
}

// UseLogger provides the [slogecho] middleware to the Router.
func (r Router) UseLogger(logger *slog.Logger) {
    r.Echo.Use(slogecho.New(logger))
}

// New creates a new Router.
func New() Router {
	e := echo.New()
	return Router{e}
}
