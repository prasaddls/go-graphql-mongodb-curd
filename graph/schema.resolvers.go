package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.40

import (
	"context"
	"fmt"

	"github.com/prasaddls/go-graphql-mongodb-curd/graph/model"
)

// CreateJobListing is the resolver for the createJobListing field.
func (r *mutationResolver) CreateJobListing(ctx context.Context, input model.CreateJobListingInput) (*model.JobListing, error) {
	panic(fmt.Errorf("not implemented: CreateJobListing - createJobListing"))
}

// UpdateJobListing is the resolver for the updateJobListing field.
func (r *mutationResolver) UpdateJobListing(ctx context.Context, id string, input model.UpdateJobListingInput) (*model.JobListing, error) {
	panic(fmt.Errorf("not implemented: UpdateJobListing - updateJobListing"))
}

// DeleteJobListing is the resolver for the deleteJobListing field.
func (r *mutationResolver) DeleteJobListing(ctx context.Context, id string) (*model.DeleteJobResponse, error) {
	panic(fmt.Errorf("not implemented: DeleteJobListing - deleteJobListing"))
}

// Jobs is the resolver for the jobs field.
func (r *queryResolver) Jobs(ctx context.Context) ([]*model.JobListing, error) {
	panic(fmt.Errorf("not implemented: Jobs - jobs"))
}

// Job is the resolver for the job field.
func (r *queryResolver) Job(ctx context.Context, id string) (*model.JobListing, error) {
	panic(fmt.Errorf("not implemented: Job - job"))
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
