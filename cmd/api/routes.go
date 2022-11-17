// Filename: cmd/api/routes

package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	// Create a new httprouter router instance
	router := httprouter.New()
	router.NotFound = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)
	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)
	router.HandlerFunc(http.MethodGet, "/v1/toasts", app.listToastsHandler)
	router.HandlerFunc(http.MethodPost, "/v1/toasts", app.createToastHandler)
	router.HandlerFunc(http.MethodGet, "/v1/toasts/:id", app.showToastHandler)
	router.HandlerFunc(http.MethodPatch, "/v1/toasts/:id", app.updateToastHandler)
	router.HandlerFunc(http.MethodDelete, "/v1/toasts/:id", app.deleteToastHandler)

	return app.recoverPanic(router)
}
