package gql

import (
	"github.com/BlackRule/App-and-its-features-CRUD/gql/gen"
	"github.com/BlackRule/App-and-its-features-CRUD/services/authservice"
	"github.com/BlackRule/App-and-its-features-CRUD/services/emailservice"
	"github.com/BlackRule/App-and-its-features-CRUD/services/userservice"
)

// Resolver struct
type Resolver struct {
	UserService  userservice.UserService
	AuthService  authservice.AuthService
	EmailService emailservice.EmailService
}

// Mutation graphql
func (r *Resolver) Mutation() gen.MutationResolver {
	return &mutationResolver{r}
}

// Query graphql
func (r *Resolver) Query() gen.QueryResolver {
	return &queryResolver{r}
}

type queryResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
