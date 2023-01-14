package controller

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.22

import (
	"context"
	"log"
	"thegame/graph"
	"thegame/model"

	"github.com/google/uuid"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	log.Println("Creating a new USER")
	uuidValue := uuid.NewString()
	user := &model.User{ID: uuidValue, Name: input.Name}

	_, err := user.Create(ctx, user)
	if err != nil {
		log.Printf("error while creating user: %v", err)
		return nil, err
	}

	return user, nil
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	var users model.User
	return users.GetAll(ctx)
}

// Mutation returns graph.MutationResolver implementation.
func (r *Resolver) Mutation() graph.MutationResolver { return &mutationResolver{r} }

// Query returns graph.QueryResolver implementation.
func (r *Resolver) Query() graph.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
