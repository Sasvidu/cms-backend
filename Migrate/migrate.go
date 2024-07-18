package main

import (
	initializers "github.com/Sasvidu/cms-backend/Initializers"
	models "github.com/Sasvidu/cms-backend/Models"
)

func init() { 
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Article{})
}