package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hayasha/blog/initializers"
	"github.com/hayasha/blog/models"
	"html/template"
)

func PostCreate(c *gin.Context) {
	if initializers.DB == nil {
		c.JSON(500, gin.H{
			"error": "Database not initialized",
		})
		return
	}

	post := models.Post{Title: "hello", Content: "Post Content"}
	result := initializers.DB.Create(&post)

	c.JSON(200, gin.H{
		"data":    result,
		"message": "post created",
	})
}

func PostsIndex(c *gin.Context) {
	var posts []models.Post
	initializers.DB.Find(&posts)

	t, _ := template.ParseFiles("templates/index.html")
	err := t.Execute(c.Writer, posts)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}

func PostDetail(c *gin.Context) {
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

func PostUpdate(c *gin.Context) {
	id := c.Param("id")

	// Request Body
	var body struct {
		Title   string
		Content string
	}
	c.Bind(&body)

	var post models.Post
	initializers.DB.First(&post, id)

	initializers.DB.Model(&post).Updates(models.Post{
		Title:   body.Title,
		Content: body.Content,
	})

	c.JSON(200, gin.H{
		"data":    post,
		"message": "OK",
	})
}

func PostDelete(c *gin.Context) {
	id := c.Param("id")

	var post models.Post
	initializers.DB.First(&post, id)
	initializers.DB.Delete(&post)

	c.JSON(200, gin.H{
		"data":    post,
		"message": "DELETED",
	})
}
