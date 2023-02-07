package graph

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/ochom/api-design/pkg/controllers"
	"github.com/ochom/api-design/pkg/handlers/graph/generated"
)

//go:generate go get github.com/99designs/gqlgen@v0.17.18
//go:generate go run github.com/99designs/gqlgen generate

// Resolver ...
// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.
type Resolver struct {
	ctl *controllers.Controller
}

// NewResolver ...
func NewResolver(ctl *controllers.Controller) *Resolver {
	return &Resolver{ctl}
}

// Play ...
func (r *Resolver) Play() gin.HandlerFunc {
	return func(c *gin.Context) {
		h := playground.Handler("GraphQL", "/query")
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// Serve ...
func (r *Resolver) Serve() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		h := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: r}))
		h.ServeHTTP(ctx.Writer, ctx.Request)
	}
}
