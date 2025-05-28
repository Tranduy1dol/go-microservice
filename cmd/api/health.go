package main

import (
	"net/http"
)

// HealthCheck godoc
//
//	@Summary		Health check
//	@Description	Check service health
//	@Tags			health
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	error
//	@Router			/health [get]
func (app *application) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"status":      "available",
		"environment": app.config.env,
		"version":     app.config.version,
	}

	if err := app.jsonResponse(w, http.StatusOK, data); err != nil {
		app.badRequestResponse(w, r, err)
	}
}
