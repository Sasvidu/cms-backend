package controllers

import (
	initializers "github.com/Sasvidu/cms-backend/Initializers"
	models "github.com/Sasvidu/cms-backend/Models"
	"github.com/gin-gonic/gin"
)

func CreateArticle(c *gin.Context) {
	var body struct {
		Title string
		Body  string
	}

	c.Bind(&body)

	article := models.Article{Title: body.Title, Body: body.Body}
	result := initializers.DB.Create(&article)

	if result.Error != nil {
		c.Status(500)
		return
	}

	c.JSON(200, gin.H{
		"article": article,
	})
}

func GetAllArticles(c *gin.Context) {
	var articles []models.Article
	initializers.DB.Find(&articles)

	c.JSON(200, gin.H{
		"articles": articles,
	})
}

func GetArticle(c *gin.Context) {
	id := c.Param("id")

	var article models.Article
	initializers.DB.First(&article, id)

	c.JSON(200, gin.H{
		"article": article,
	})
}

func UpdateArticle(c *gin.Context) {
	id := c.Param("id")
	var article models.Article
	initializers.DB.First(&article, id)

	var body struct {
		Title string
		Body  string
	}
	c.Bind(&body)
	initializers.DB.Model(&article).Updates(models.Article{Title: body.Title, Body: body.Body})

	c.JSON(200, gin.H{
		"article": article,
	})
}

func DeleteArticle(c *gin.Context) {
	id := c.Param("id")
	initializers.DB.Delete(&models.Article{}, id)

	if initializers.DB.Error != nil {
		c.Status(500)
	} else {
		c.Status(200)
	}
}