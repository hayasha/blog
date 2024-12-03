package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hayasha/blog/initializers"
	"github.com/hayasha/blog/models"
	"html/template"
)

func SetPostRouter(router *gin.Engine) *gin.Engine {
	posts := router.Group("/posts")

	posts.GET("/index", postsIndex)
	posts.GET("/:id", postDetail)

	return router
}

func postsIndex(c *gin.Context) {
	var posts []models.Post
	initializers.DB.Find(&posts)

	t, _ := template.ParseFiles("templates/index.html")
	err := t.Execute(c.Writer, posts)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}

func postDetail(c *gin.Context) {
	id := c.Param("id")

	var post models.Post
	initializers.DB.First(&post, id)

	t, _ := template.ParseFiles("templates/post.html")
	err := t.Execute(c.Writer, post)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}
