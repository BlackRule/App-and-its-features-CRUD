package App_and_its_features_CRUD

import (
	"fmt"
	"github.com/BlackRule/App-and-its-features-CRUD/entities/passwordreset"
	"github.com/BlackRule/App-and-its-features-CRUD/entities/user"
	"github.com/BlackRule/App-and-its-features-CRUD/rest/controllers"
	"log"
	"net/http"

	"github.com/BlackRule/App-and-its-features-CRUD/common/hmachash"
	"github.com/BlackRule/App-and-its-features-CRUD/common/randomstring"
	"github.com/BlackRule/App-and-its-features-CRUD/configs"
	pwdDomain "github.com/BlackRule/App-and-its-features-CRUD/entities/passwordreset"
	"github.com/BlackRule/App-and-its-features-CRUD/graphql"
	"github.com/BlackRule/App-and-its-features-CRUD/middlewares"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"

	_ "github.com/BlackRule/App-and-its-features-CRUD/rest/docs" // docs is generated by Swag CLI

	"github.com/BlackRule/App-and-its-features-CRUD/services/authservice"
	_ "github.com/lib/pq" // For Postgres setup
)

var (
	router = gin.Default()
)

func main() {

	/*
		====== Setup configs ============
	*/
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	config := configs.GetConfig()

	// Connects to PostgresDB
	db, err := gorm.Open(
		config.Postgres.Dialect(),
		config.Postgres.GetPostgresConnectionInfo(),
	)
	if err != nil {
		panic(err)
	}

	// Migration
	// db.DropTableIfExists(&user.User{})
	db.AutoMigrate(&user.User{}, &pwdDomain.PasswordReset{})
	defer db.Close()

	/*
		=======  Setup Common ============
	*/
	rds := randomstring.NewRandomString()
	hm := hmachash.NewHMAC(config.HMACKey)

	/*
		====== Setup repositories =======
	*/
	userRepo := user.NewUserRepo(db)
	pwdRepo := passwordreset.NewPasswordResetRepo(db)

	/*
		====== Setup services ===========
	*/
	userService := user.NewUserService(userRepo, pwdRepo, rds, hm, config.Pepper)
	authService := authservice.NewAuthService(config.JWTSecret)

	/*
		====== Setup controllers ========
	*/
	userCtl := controllers.NewUserController(userService, authService)

	/*
		====== Setup middlewares ========
	*/
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	/*
		====== Setup routes =============
	*/
	router.GET("/ping", func(c *gin.Context) { c.String(http.StatusOK, "pong") })

	router.GET("/graphql", graphql.PlaygroundHandler("/query"))
	router.POST("/query",
		middlewares.SetUserContext(config.JWTSecret),
		graphql.GraphqlHandler(userService, authService))

	api := router.Group("/api")

	api.POST("/register", userCtl.Register)
	api.POST("/login", userCtl.Login)
	api.POST("/forgot_password", userCtl.ForgotPassword)
	api.POST("/update_password", userCtl.ResetPassword)

	user := api.Group("/users")

	user.GET("/:id", userCtl.GetByID)

	account := api.Group("/account")
	account.Use(middlewares.RequireLoggedIn(config.JWTSecret))
	{
		account.GET("/profile", userCtl.GetProfile)
		account.PUT("/profile", userCtl.Update)
	}

	// Run
	// port := fmt.Sprintf(":%s", viper.Get("APP_PORT"))
	port := fmt.Sprintf(":%s", config.Port)
	router.NoRoute(gin.WrapH(http.FileServer(http.Dir("static"))))
	router.Run(port)
}