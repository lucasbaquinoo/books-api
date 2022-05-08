package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/lucasbaquinoo/webapi-with-go/controllers"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		user := main.Group("user")
		{
			user.POST("/", controllers.CreateUser)
		}
		books := main.Group("books")
		{
			books.GET("/:id", controllers.ShowBook)
			books.GET("/", controllers.ShowAllBooks)
			books.POST("/", controllers.CreateBook)
			books.PUT("/", controllers.EditBook)
			books.DELETE("/:id", controllers.DeleteBook)
		}
	}

	return router
}
