package controllers

import (
	"github.com/albizzy/go-crud/initializers"
	"github.com/albizzy/go-crud/models"
	"github.com/gin-gonic/gin"
)

func PostCreate(c *gin.Context) {
	// Get data off req body
	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	// Create a post
	post := models.Post{Title: body.Title, Body: body.Body}

	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	// Return it
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostIndex(c *gin.Context) {
	// Get the post
	var posts []models.Post
	initializers.DB.Find(&posts)

	// Respond with them
	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func PostShow(c *gin.Context) {
	// get id off url
	id := c.Param("id")

	// Get the post
	var post models.Post
	initializers.DB.First(&post, id)

	// Respond with them
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostUpdate(c *gin.Context) {
	// get id of the url
	id := c.Param("id")

	// get the data off req body
	var body struct {
		Body string
		Title string
	}

	c.Bind(&body)
	// find the post to be updated
	var post models.Post
	initializers.DB.First(&post, id)

	// update it
	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body: body.Body,
	})

	// respond with it
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostDelete(c *gin.Context) {
	// get id of the url
	id := c.Param("id")

	// delete the post
	initializers.DB.Delete(&models.Post{}, id)

	// respond
	c.Status(200)
}