package graphql

import (
	"github.com/BlackRule/App-and-its-features-CRUD/entities/user"
	"github.com/BlackRule/App-and-its-features-CRUD/graphql/gen"
	"github.com/BlackRule/App-and-its-features-CRUD/services/authservice"
	"github.com/BlackRule/App-and-its-features-CRUD/services/emailservice"
)

// Resolver struct
type Resolver struct {
	UserService  user.UserService
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
