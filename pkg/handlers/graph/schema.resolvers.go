package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/ochom/api-design/pkg/handlers/graph/generated"
	"github.com/ochom/api-design/pkg/handlers/graph/model"
	"github.com/ochom/api-design/pkg/models"
)

// PingMe is the resolver for the pingMe field.
func (r *mutationResolver) PingMe(ctx context.Context, input model.Pang) (*models.Ping, error) {
	return &models.Ping{Message: "pong"}, nil
}

// Ping is the resolver for the ping field.
func (r *queryResolver) Ping(ctx context.Context) (string, error) {
	return "pong", nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
