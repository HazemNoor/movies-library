package app

import (
	"github.com/HazemNoor/movies-library/app/controllers"
	"github.com/HazemNoor/movies-library/app/middleware"
	"github.com/gin-gonic/gin"
)

// Dispatch is handle routing
func Dispatch() error {
	r := gin.Default()

	authMiddleware := middleware.NewAuthMiddleware()

	userController := controllers.NewUserController()
	authController := controllers.NewAuthController()
	movieController := controllers.NewMovieController()
	watchedListController := controllers.NewWatchedListController()

	r.POST("user", userController.Register)

	r.POST("auth/login", authController.Login)
	r.GET("auth/check", authMiddleware.Handle, authController.Check)

	r.GET("movies", authMiddleware.Handle, movieController.List)
	r.POST("movie", authMiddleware.Handle, movieController.Create)
	r.PATCH("movie/:id", authMiddleware.Handle, movieController.Update)
	r.GET("movie/:id", authMiddleware.Handle, movieController.Show)
	r.DELETE("movie/:id", authMiddleware.Handle, movieController.Delete)

	r.GET("watched-list", authMiddleware.Handle, watchedListController.List)
	r.POST("watched-list", authMiddleware.Handle, watchedListController.Add)
	r.PATCH("watched-list/:id", authMiddleware.Handle, watchedListController.Update)
	r.DELETE("watched-list/:id", authMiddleware.Handle, watchedListController.Delete)

	return r.Run()
}
