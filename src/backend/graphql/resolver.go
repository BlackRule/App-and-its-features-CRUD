package graphql

//go:generate go run github.com/99designs/gqlgen generate
import (
	"github.com/BlackRule/App-and-its-features-CRUD/entities/user"
	"github.com/BlackRule/App-and-its-features-CRUD/graphql/gen"
	"github.com/BlackRule/App-and-its-features-CRUD/services/authservice"
	"github.com/BlackRule/App-and-its-features-CRUD/services/emailservice"
	"github.com/jinzhu/gorm"
)

type Resolver struct {
	UserService  user.UserService
	AuthService  authservice.AuthService
	EmailService emailservice.EmailService
	db           *gorm.DB
}

func (r *Resolver) Mutation() gen.MutationResolver {
	return &mutationResolver{r}
}

func (r *Resolver) Query() gen.QueryResolver {
	return &queryResolver{r}
}

type queryResolver struct{ *Resolver }

type mutationResolver struct{ *Resolver }
