package main

import (
	"net/http"
)

func (app *application) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"status":      "available",
		"environment": app.config.env,
		"version":     app.config.version,
	}

	if err := writeJSON(w, http.StatusOK, data); err != nil {
		app.badRequestResponse(w, r, err)
	}
}
