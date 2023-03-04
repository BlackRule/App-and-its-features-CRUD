package gql

import (
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/BlackRule/App-and-its-features-CRUD/entities/user"
	"github.com/BlackRule/App-and-its-features-CRUD/gql/gen"

	"github.com/99designs/gqlgen/handler"
	"github.com/BlackRule/App-and-its-features-CRUD/services/authservice"
	"github.com/gin-gonic/gin"
)

// GraphqlHandler defines the GQLGen GraphQL server handler
func GraphqlHandler(
	us user.UserService,
	as authservice.AuthService) gin.HandlerFunc {
	// NewExecutableSchema and Config are in the generated.go file
	conf := gen.Config{
		Resolvers: &Resolver{
			UserService: us,
			AuthService: as,
		},
	}
	exec := gen.NewExecutableSchema(conf)
	h := handler.GraphQL(exec)
	return func(c *gin.Context) { h.ServeHTTP(c.Writer, c.Request) }
}

// PlaygroundHandler Defines the Playground handler to expose our playground
func PlaygroundHandler(path string) gin.HandlerFunc {
	h := playground.Handler("GraphQL Playground", path)
	return func(c *gin.Context) { h.ServeHTTP(c.Writer, c.Request) }
}
