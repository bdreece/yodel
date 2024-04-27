package views

type (
    // M provides a simple key-value mapping type
    M map[string]any

    // View provides the response type for view resolvers.
	View struct {
        // The name of the template to render.
		Name string
        // The data to pass to the template.
		Data M

        // An optional status code. Defaults to [http.StatusOK].
        StatusCode int
	}
)
