package routers

import (
	"github.com/gin-gonic/gin"
)

func InitializeRouter(r *gin.RouterGroup, env ...string) {

	environment := "Production"

	if len(env) > 0 {
		environment = env[0]
	}

	var routes = r.Group("/library")
	{
		// Get Book Info
		routes.GET("book/:id", func(c *gin.Context) {
			var bookId string = c.Param("id")

			c.JSON(200, gin.H{
				"BookID":      bookId,
				"data":        "Book Demo !!!",
				"Environment": environment,
			})
		})

		// Add Book to Library
		routes.POST("add-book", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"data": "Add Book to Library !!!",
			})
		})

		// Remove Book From Library
		routes.DELETE("remove-book", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"data": "Remove Book from Library !!!",
			})
		})

	}

}
