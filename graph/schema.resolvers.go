package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"cat-turner/outerbanks-api/graph/generated"
	"cat-turner/outerbanks-api/graph/model"
	"context"
	"fmt"
)

func (r *mutationResolver) UpsertCharacter(ctx context.Context, input model.CharacterInput) (*model.Character, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Character(ctx context.Context, id string) (*model.Character, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Pogues(ctx context.Context) ([]*model.Character, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Kooks(ctx context.Context) ([]*model.Character, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Search(ctx context.Context, text string) ([]*model.SearchResult, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
