package main

import (
	"github.com/albizzy/go-crud/controllers"
	"github.com/albizzy/go-crud/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDatabase()
}

func main() {
	r := gin.Default()
	r.POST("/posts", controllers.PostCreate)
	r.PUT("/posts/:id", controllers.PostUpdate)
	r.GET("/posts", controllers.PostIndex)
	r.GET("/posts/:id", controllers.PostShow)
	r.DELETE("/posts/:id", controllers.PostDelete)
	
	r.Run()
}
