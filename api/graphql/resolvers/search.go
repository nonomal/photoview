package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.64

import (
	"context"

	"github.com/photoview/photoview/api/graphql/auth"
	"github.com/photoview/photoview/api/graphql/models"
	"github.com/photoview/photoview/api/graphql/models/actions"
)

// Search is the resolver for the search field.
func (r *queryResolver) Search(ctx context.Context, query string, limitMedia *int, limitAlbums *int) (*models.SearchResult, error) {
	user := auth.UserFromContext(ctx)
	if user == nil {
		return nil, auth.ErrUnauthorized
	}

	return actions.Search(r.DB(ctx), query, user.ID, limitMedia, limitAlbums)
}
