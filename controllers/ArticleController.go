package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jasongauvin/wikiPattern/models"
	"time"
)

var articleList = []models.Article{
	{
		ID:        0,
		Title:     "Article 1",
		Content:   "Body article 1",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	},
	{
		ID:        1,
		Title:     "Article 2",
		Content:   "Body article 2",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	},
}

func GetAllArticles() []models.Article {
	return articleList
}

func GetArticles(c *gin.Context) []models.Article {
	var articles []models.Article
	var err error
	articles, err = models.FindArticles()

	if err != nil {
		fmt.Println("Error: ", err)
	}

	return articles
}
