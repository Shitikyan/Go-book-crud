package router

import (
	"github.com/Shitikyan/Go-book-crud/controllers"
	"github.com/gin-gonic/gin"
)

func (r RouterShema) BookRoutes(rg *gin.RouterGroup) {
	book := rg.Group("/books")

	book.POST("", controllers.CreateBook)
	book.GET("", controllers.GetAllBooks)
	book.GET("/:id", controllers.GetBook)
	book.PUT("/:id", controllers.UpdateBook)
	book.DELETE("/:id", controllers.DeleteBook)
}
