package main

import (
	"fmt"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/mvrilo/go-redoc"
	ginredoc "github.com/mvrilo/go-redoc/gin"

	"golang-program-structure/common/middleware"
	"golang-program-structure/service"
)

func NewHandler(config *service.Config) *gin.Engine {
	h := gin.New()

	// Header preservation during redirect differs between Google and Apple
	// To be safe, do no automatic redirects
	h.RedirectTrailingSlash = false
	h.RedirectFixedPath = false

	// The status route is open
	h.GET("/status", config.Status)

	// Add a CORS middleware handler
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	h.Use(cors.New(corsConfig))

	// Add the fallback handler to catch any panic and return a 500 to the client
	h.Use(gin.CustomRecovery(config.PanicRecovery))

	// Add the request logging middleware handler to all service routes
	requestLogger := middleware.NewRequestLogger(config.Logger)
	h.Use(requestLogger.Middleware)

	// Setup default routes
	service.SetupDefaultEndpoints(h, config)
	// Setup custom routes
	service.AddRoutes(h, config)

	// Add the handler to serve the redoc
	specFile := "./docs/swagger.json"
	if _, err := os.Stat(specFile); err == nil {
		docs := redoc.Redoc{
			Title:       "Docs",
			Description: "Documentation",
			SpecFile:    "./docs/swagger.json",
			SpecPath:    fmt.Sprintf("/v%s/%s/docs/openapi.json", config.Version, config.ServiceName),
			DocsPath:    fmt.Sprintf("/v%s/%s/docs", config.Version, config.ServiceName),
		}
		h.Use(ginredoc.New(docs))
	} else {
		config.Logger.Warnf("Swagger file not found at %s, skipping redoc init", specFile)
	}

	return h
}
