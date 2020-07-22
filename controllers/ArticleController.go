package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jasongauvin/wikiPattern/models"
	"github.com/jasongauvin/wikiPattern/services"
	"net/http"
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

func GetArticles(c *gin.Context) {
	var articles []models.Article
	var err error
	articles, err = models.FindArticles()

	if err != nil {
		fmt.Println("Error: ", err)
	}
	c.HTML(http.StatusOK,
		"article/index.html",
		gin.H{
			"title":   "Post index",
			"payload": articles,
		})
}

func GetArticleById(c *gin.Context){
	id := services.ConvertStringToInt(c.Param("id"))

	var article models.Article
	var err error
	article, err = models.FindArticleByID(id)

	if err != nil {
		fmt.Println("Error: ", err)
	}

	c.HTML(http.StatusOK,
		"article/show.html",
		gin.H{
			"title":   "Post Page",
			"article": article,
		})
}
