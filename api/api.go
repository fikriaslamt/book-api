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

	//user routes
	r.POST("/api/login", userController.Login)
	r.GET("/api/logout", userController.Logout)

	//book routes
	r.GET("/api/book", bookController.GetAllBooks)
	r.GET("/api/book/:id", bookController.GetBook)

	auth := r.Group("/api/book")
	auth.Use(middleware.Authentication())
	{

		auth.POST("/add", bookController.AddBook)
		auth.PUT("/update/:id", bookController.UpdateBook)
		auth.DELETE("/delete/:id", bookController.DeleteBook)
	}

}

func Start() {
	r.Run()
	r.Use(cors.Default())
}
