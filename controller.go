package yodel

import "github.com/labstack/echo/v4"

// Controller provides a composable implementation of Handler, using a Binder,
// Resolver, and Responder.
type Controller[Req, Res any] struct {
	Binder    Binder[Req]
	Responder Responder[Res]
	Resolver  Resolver[Req, Res]
}

// Handle implements Handler.
func (controller *Controller[Req, Res]) Handle(c echo.Context) error {
	req, err := controller.Binder.Bind(c)
	if err != nil {
		return err
	}

	res, err := controller.Resolver.Resolve(c.Request().Context(), req)
	if err != nil {
		return err
	}

	if err := controller.Responder.Respond(c, res); err != nil {
		return err
	}

	return nil
}
