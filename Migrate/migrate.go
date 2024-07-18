package main

import (
	initializers "github.com/Sasvidu/cms-backend/Initializers"
	"github.com/Sasvidu/cms-backend/models"
)

func init() { 
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Article{})
}