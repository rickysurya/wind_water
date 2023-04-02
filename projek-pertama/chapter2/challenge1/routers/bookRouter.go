package routers

import (
	"github.com/gin-gonic/gin"
	"projek-pertama/chapter2/challenge1/controllers"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	router.POST("/books", controllers.CreateBook)
	router.GET("/books", controllers.GetAllBook)
	router.PUT("/books/:bookID", controllers.UpdateBook)
	router.GET("/books/:bookID", controllers.GetBook)
	router.DELETE("/books/:bookID", controllers.DeleteBook)

	return router
}
