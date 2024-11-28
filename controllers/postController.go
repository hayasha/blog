package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/hayasha/blog/initializers"
	"github.com/hayasha/blog/models"
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
	c.JSON(200, gin.H{
		"data":    posts,
		"message": "OK",
	})
}

func PostDetail(c *gin.Context) {
	id := c.Param("id")

	var post models.Post
	initializers.DB.First(&post, id)

	c.JSON(200, gin.H{
		"data":    post,
		"message": "OK",
	})
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
