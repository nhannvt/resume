package api

import (
	"net/http"
	"os"
	"path/filepath"

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

	r.app.Static("/assets", filepath.Join(os.Getenv("GOPATH"),
	"src/github.com/nhannvt/resume/static/assets"))


	r.app.StaticFile("/NhanNguyen.pdf", filepath.Join(os.Getenv("GOPATH"),
	"src/github.com/nhannvt/resume/static/NhanNguyen.pdf"))


	r.app.GET("/", func(c *gin.Context) {
		r.app.LoadHTMLGlob(filepath.Join(os.Getenv("GOPATH"),
			"src/github.com/nhannvt/resume/internal/interface/templates/*"))
		// Call the HTML method of the Context to render a template
		c.HTML(
			// Set the HTTP status to 200 (OK)
			http.StatusOK,
			// Use the index.html template
			"index.tmpl",
			// Pass the data that the page uses (in this case, 'title')
			gin.H{
				"title": "Nguyen Van Trong Nhan",
			},
		)

	})
}
