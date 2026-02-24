package routes

import (
	"github.com/Jason-cqtan/simple-blog/config"
	"github.com/Jason-cqtan/simple-blog/handlers"
	"github.com/Jason-cqtan/simple-blog/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(router *gin.Engine, db *gorm.DB, cfg *config.Config) {
	userHandler := handlers.NewUserHandler(db, cfg)
	postHandler := handlers.NewPostHandler(db)
	commentHandler := handlers.NewCommentHandler(db)

	// Public routes
	router.GET("/", postHandler.Home)
	router.GET("/posts", postHandler.List)
	router.GET("/posts/:id", postHandler.Show)
	router.GET("/login", userHandler.ShowLoginForm)
	router.GET("/register", userHandler.ShowRegisterForm)
	router.POST("/login", userHandler.Login)
	router.POST("/register", userHandler.Register)
	router.POST("/logout", userHandler.Logout)

	// Protected routes
	auth := router.Group("/")
	auth.Use(middleware.JWTAuthMiddleware(cfg.JWTSecret))
	{
		auth.GET("/posts/new", postHandler.ShowCreateForm)
		auth.POST("/posts", postHandler.Create)
		auth.GET("/posts/:id/edit", postHandler.ShowEditForm)
		auth.POST("/posts/:id/update", postHandler.Update)
		auth.POST("/posts/:id/delete", postHandler.Delete)
		auth.POST("/posts/:id/comments", commentHandler.Create)
		auth.GET("/profile", userHandler.ShowProfile)
	}
}
