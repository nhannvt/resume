package api

import (
	"github.com/gin-gonic/gin"
	"github.com/nhannvt/resume/internal/registry"
)

type Router interface {
	Setup()
}

type router struct {
	app      *gin.Engine
	registry registry.Registry
}

func NewRouter(app *gin.Engine, registry registry.Registry) Router {
	router := &router{app, registry}
	return router
}

func (r *router) Setup() {

	healthCheckHandler := r.registry.NewHealthCheckHandler()
	r.app.GET("/health", healthCheckHandler.HealthCheck)

	notFoundHandler := r.registry.NewNotFoundHandler()
	r.app.NoRoute(notFoundHandler.NotFound)
}
