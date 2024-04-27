package yodel

import "github.com/labstack/echo/v4"

// Middleware wraps a request Handler, accessing and modifying the
// request and response values.
type Middleware interface {
    Invoke(next echo.HandlerFunc) echo.HandlerFunc
}

// MiddlewareFunc provides a function implementation of Middleware.
type MiddlewareFunc func(next echo.HandlerFunc) echo.HandlerFunc

// Invoke implements Middleware.
func (fn MiddlewareFunc) Invoke(next echo.HandlerFunc) echo.HandlerFunc {
    return fn(next)
}

func mapMiddlewares(middleware []Middleware) []echo.MiddlewareFunc {
    mws := make([]echo.MiddlewareFunc, len(middleware))
    for _, mw := range middleware {
        mws = append(mws, mw.Invoke)
    }

    return mws
}
