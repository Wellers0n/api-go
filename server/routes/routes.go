package routes

import (
	"webapi-with-go/controllers"

	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		books := main.Group("books")
		{
			books.GET("/", controllers.ShowAllBooks)
			books.GET("/:id", controllers.ShowBook)
			books.POST("/", controllers.CreateBook)
			books.PUT("/:id", controllers.EditBook)
			books.DELETE("/:id", controllers.DeleteBook)
		}
	}

	return router
}
