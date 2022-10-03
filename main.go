package main

import (
	"go-sample/config"
	v1 "go-sample/handler/v1"
	"go-sample/middleware"
	"go-sample/repo"
	"go-sample/service"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db           *gorm.DB             = config.SetupDatabaseConnection()
	userRepo     repo.UserRepository  = repo.NewUserRepo(db)
	tweetRepo    repo.TweetRepository = repo.NewTweetRepo(db)
	authService  service.AuthService  = service.NewAuthService(userRepo)
	jwtService   service.JWTService   = service.NewJWTService()
	userService  service.UserService  = service.NewUserService(userRepo)
	tweetService service.TweetService = service.NewTweetService(tweetRepo)
	authHandler  v1.AuthHandler       = v1.NewAuthHandler(authService, jwtService, userService)
	userHandler  v1.UserHandler       = v1.NewUserHandler(userService, jwtService)
	tweetHandler v1.TweetHandler      = v1.NewTweetHandler(tweetService, jwtService)
)

func main() {
	defer config.CloseDatabaseConnection(db)
	server := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	config.AllowCredentials = true
	server.Use(cors.New(config))

	authRoutes := server.Group("api/auth")
	{
		authRoutes.POST("/login", authHandler.Login)
		authRoutes.POST("/register", authHandler.Register)
	}

	userRoutes := server.Group("api/user", middleware.AuthorizeJWT(jwtService))
	{
		userRoutes.GET("/profile", userHandler.Profile)
		userRoutes.PUT("/profile", userHandler.Update)
	}

	tweetRoutes := server.Group("api/tweet", middleware.AuthorizeJWT(jwtService))
	{
		tweetRoutes.GET("/", tweetHandler.All)
		tweetRoutes.POST("/", tweetHandler.CreateTweet)
		tweetRoutes.GET("/:id", tweetHandler.FindOneTweetByID)
		tweetRoutes.PUT("/:id", tweetHandler.UpdateTweet)
		tweetRoutes.DELETE("/:id", tweetHandler.DeleteTweet)
	}

	checkRoutes := server.Group("api/check")
	{
		checkRoutes.GET("health", v1.Health)
	}

	server.Run()
}
