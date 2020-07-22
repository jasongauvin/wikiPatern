package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jasongauvin/wikiPattern/models"
)

func GetArticles(c *gin.Context) []models.Article {
	var articles []models.Article
	var err error
	articles, err = models.FindArticles()

	if err != nil {
		fmt.Println("Error: ", err)
	}

	return articles
}
