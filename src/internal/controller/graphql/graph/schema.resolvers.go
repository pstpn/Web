package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.49

import (
	"context"
	"time"
)

// Healthcheck is the resolver for the healthcheck field.
func (r *queryResolver) Healthcheck(ctx context.Context) (*string, error) {
	now := time.Now().String()
	return &now, nil
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }