/*
 api package is the package which provides Sforum API server.
 API server is boot via command under cmd directory.
*/
package api

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nhannvt/resume/configs"
	"github.com/nhannvt/resume/internal/interface/middleware"
	"github.com/nhannvt/resume/internal/registry"
	"github.com/nhannvt/resume/pkg/log"
	"github.com/nhannvt/resume/pkg/util"
)

// Server type
type Server struct {
	Options *Options
}

// Type Options contains the information for server behavior
type Options struct {
	// Port to listen
	Port string
	// Stage Name
	Stage string
	// The list of middleware name which will be turned off with this list
	Without []string
	// The environment name like prod, test, dev.
	// When debug is set, Debug message will be displayed
	Debug bool
	// The seconds to wait since application starts shutdown.
	TerminationPeriod int64
}

// NewServer returns server instance with given options
func NewServer(options *Options) *Server {
	return &Server{Options: options}
}

// Start runs application while importing configuration files
func (server *Server) Start() {

	app := server.CreateApplication()

	srv := &http.Server{
		Addr:    ":" + server.Options.Port,
		Handler: app,
	}

	// Followings are logic for graceful showdown

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of specified seconds.
	quit := make(chan os.Signal)
	// kill (no param) default send syscanll.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	log.Infof("SIGNAL %d received, Shutdown Server...", <-quit)

	duration := time.Duration(server.Options.TerminationPeriod)
	ctx, cancel := context.WithTimeout(context.Background(), duration*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server Shutdown: %s", err)
	}

	select {
	case <-ctx.Done():
		log.Infof("timeout of %d seconds.", server.Options.TerminationPeriod)
	}
	log.Info("Server exiting")
}

// CreateApplication returns new gin.Engine includes custom middlewares.
func (server *Server) CreateApplication() *gin.Engine {
	if !server.Options.Debug {
		gin.SetMode(gin.ReleaseMode)
		log.SetDefaultLoggerDebugMode(false)
	}

	app := gin.New()
	server.loadMiddlewares(app)
	config := configs.GetConfig()
	registry := registry.NewRegistry(config)
	router := NewRouter(app, registry)
	router.Setup()

	return app
}

func (server *Server) CreateTestApplication() *gin.Engine {
	if !server.Options.Debug {
		gin.SetMode(gin.ReleaseMode)
		log.SetDefaultLoggerDebugMode(false)
	}

	app := gin.New()
	server.loadMiddlewares(app)
	config := configs.GetConfig(configs.TestAPIEndpoint())
	registry := registry.NewRegistry(config)
	router := NewRouter(app, registry)
	router.Setup()

	return app
}

// loadMiddlewares registers custom middlewares.
// Using 'without' option, You can skip registration of these middleware.
func (server *Server) loadMiddlewares(app *gin.Engine) {

	app.Use(gin.Recovery())

	if !util.ContainString(server.Options.Without, "log") {
		app.Use(middleware.Logger())
	}

	// if !util.ContainString(server.Options.Without, "authenticate") {
	// 	app.Use(middleware.Authenticate(server.Options.Stage))
	// }

	// if !util.ContainString(server.Options.Without, "authorize") {
	// 	app.Use(middleware.Authorize())
	// }

}
