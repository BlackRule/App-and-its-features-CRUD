package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/BlackRule/App-and-its-features-CRUD/common/hmachash"
	"github.com/BlackRule/App-and-its-features-CRUD/common/randomstring"
	"github.com/BlackRule/App-and-its-features-CRUD/configs"
	pwdDomain "github.com/BlackRule/App-and-its-features-CRUD/models/passwordreset"
	"github.com/BlackRule/App-and-its-features-CRUD/models/user"

	"github.com/BlackRule/App-and-its-features-CRUD/gql"
	"github.com/BlackRule/App-and-its-features-CRUD/middlewares"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"

	_ "github.com/BlackRule/App-and-its-features-CRUD/docs" // docs is generated by Swag CLI

	"github.com/BlackRule/App-and-its-features-CRUD/controllers"
	"github.com/BlackRule/App-and-its-features-CRUD/repositories/passwordreset"
	"github.com/BlackRule/App-and-its-features-CRUD/repositories/userrepo"
	"github.com/BlackRule/App-and-its-features-CRUD/services/authservice"
	"github.com/BlackRule/App-and-its-features-CRUD/services/userservice"

	_ "github.com/lib/pq" // For Postgres setup
)

var (
	router = gin.Default()
)

// @title Go API Boilerplate Swagger
// @version 1.0
// @description This is Go API Boilerplate
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @host localhost:3000
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func Run() {

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
	userRepo := userrepo.NewUserRepo(db)
	pwdRepo := passwordreset.NewPasswordResetRepo(db)

	/*
		====== Setup services ===========
	*/
	userService := userservice.NewUserService(userRepo, pwdRepo, rds, hm, config.Pepper)
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

	router.GET("/graphql", gql.PlaygroundHandler("/query"))
	router.POST("/query",
		middlewares.SetUserContext(config.JWTSecret),
		gql.GraphqlHandler(userService, authService))

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
