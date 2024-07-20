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
	r.GET("/articles/all", controllers.GetAllArticles)
	r.GET("/articles/:id", controllers.GetArticle)
	r.POST("/articles/create", controllers.CreateArticle)
	r.PUT("/articles/edit/:id", controllers.UpdateArticle)
	r.DELETE("/articles/delete/:id", controllers.DeleteArticle)

	r.Run()
}