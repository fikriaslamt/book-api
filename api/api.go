package api

import (
	"book-api/controller/bookController"
	"book-api/controller/userController"
	"book-api/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

var r = gin.Default()

func Routes() {

	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("session_token", store))

	config := cors.DefaultConfig()
	config.AllowCredentials = true
	config.AllowOrigins = []string{"http://114.79.6.174:3000/", "http://google.com"}

	r.Use(cors.New(config))
	//user routes
	r.POST("/api/login", userController.Login)
	r.GET("/api/logout", userController.Logout)

	//book routes
	r.GET("/api/book", bookController.GetAllBooks)
	r.GET("/api/book/:id", bookController.GetBook)

	r.POST("/api/book/add", middleware.Auth, bookController.AddBook)
	r.PUT("/api/book/update/:id", middleware.Auth, bookController.UpdateBook)
	r.DELETE("/api/book/delete/:id", middleware.Auth, bookController.DeleteBook)

}

func Start() {

	r.Run()

}
