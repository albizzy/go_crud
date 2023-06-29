package main

import (
	"github.com/albizzy/go-crud/initializers"
	"github.com/albizzy/go-crud/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDatabase()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
