package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/ochom/api-design/pkg/controllers"
)

// API is the interface for the API
type API struct {
	ctl *controllers.Controller
}

// NewRestAPI creates a new API
func NewRestAPI(ctl *controllers.Controller) *API {
	return &API{ctl: ctl}
}

// Ping is a simple ping handler
func (a *API) Ping() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(200, "pong")
	}
}
