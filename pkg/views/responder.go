package views

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Responder responds to the request by rendering the template
// specified by view to the response body as HTML.
type Responder struct{}

// Respond implements Responder.
func (Responder) Respond(c echo.Context, view View) error {
    if view.StatusCode == 0 {
        view.StatusCode = http.StatusOK
    }

    return c.Render(view.StatusCode, view.Name, view.Data)
}
