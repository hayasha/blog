package main

import (
	"github.com/hayasha/blog/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hayasha/blog/initializers"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectToMysql()
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.POST("/posts", controllers.PostCreate)

	r.GET("/posts", controllers.PostsIndex)
	r.GET("/posts/:id", controllers.PostDetail)

	r.PUT("/posts/:id", controllers.PostUpdate)

	r.DELETE("/posts/:id", controllers.PostDelete)

	r.Run()
}
