package main

import (
	controllers "github.com/Sasvidu/cms-backend/Controllers"
	initializers "github.com/Sasvidu/cms-backend/Initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables() 
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

	initializers.SetupCORS(r)

	r.GET("/articles", controllers.GetAllArticles)
	r.GET("/articles/:id", controllers.GetArticle)
	r.POST("/articles", controllers.CreateArticle)
	r.PUT("/articles/:id", controllers.UpdateArticle)
	r.DELETE("/articles/:id", controllers.DeleteArticle)

	r.Run()
}