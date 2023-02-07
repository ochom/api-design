package graph

import "github.com/ochom/api-design/pkg/controllers"

//go:generate go get github.com/99designs/gqlgen@v0.17.18
//go:generate go run github.com/99designs/gqlgen generate

// Resolver ...
// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.
type Resolver struct {
	ctl *controllers.Controller
}
