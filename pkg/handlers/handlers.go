package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/ochom/api-design/pkg/controllers"
	// "github.com/ochom/api-design/pkg/handlers/graphql"
)

// Handler is a handler interface
type Handler struct {
	ctl *controllers.Controller
}

// NewHandler creates a new handler instance
func NewHandler() *Handler {
	ctl := controllers.NewController()
	return &Handler{
		ctl: ctl,
	}
}

// Routes returns the routes for the handler
func (h *Handler) Routes() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.GET("/ping", h.Ping)

	// rest api routes

	// graphql routes
	// graph := router.Group("/graphql")
	// graph.GET("/play", graphql.PlaygroundHandler())
	// graph.POST("/query", graphql.RequestHandler(h.ctl))

	return router
}

// Ping is a ping handler
func (h *Handler) Ping(c *gin.Context) {
	c.String(200, "pong")
}
