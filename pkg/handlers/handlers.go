package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/ochom/api-design/pkg/controllers"
	"github.com/ochom/api-design/pkg/handlers/graph"
	"github.com/ochom/api-design/pkg/handlers/rest"
)

// Handler is a handler interface
type Handler struct {
	ctl           *controllers.Controller
	graphResolver *graph.Resolver
	rest          *rest.API
}

// NewHandler creates a new handler instance
func NewHandler() *Handler {
	ctl := controllers.NewController()
	gr := graph.NewResolver(ctl)
	rest := rest.NewRestAPI(ctl)
	return &Handler{
		ctl:           ctl,
		graphResolver: gr,
		rest:          rest,
	}
}

// Routes returns the routes for the handler
func (h *Handler) Routes() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.GET("/ping", h.Ping)

	// rest api routes
	rest := router.Group("/v1")
	{
		rest.GET("/ping", h.rest.Ping())
	}

	// graphql routes
	graphql := router.Group("/graphql")
	{
		graphql.GET("/play", h.graphResolver.Play())
		graphql.POST("/query", h.graphResolver.Serve())
	}

	return router
}

// Ping is a ping handler
func (h *Handler) Ping(c *gin.Context) {
	c.String(200, "pong")
}
