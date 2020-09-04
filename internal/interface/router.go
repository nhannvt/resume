package api

import (
	"net/http"

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

	r.app.GET("/", func(c *gin.Context) {
		r.app.LoadHTMLGlob("templates/*")
		// Call the HTML method of the Context to render a template
		c.HTML(
			// Set the HTTP status to 200 (OK)
			http.StatusOK,
			// Use the index.html template
			"index.html",
			// Pass the data that the page uses (in this case, 'title')
			gin.H{
				"title": "Home Page",
			},
		)

	})
}
