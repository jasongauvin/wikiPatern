package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jasongauvin/wikiPattern/services"
	"net/http"
)

func GetHomePage(c *gin.Context) {
	articles, err := services.LoadArticleByOrderedDate()
	if err != nil {
		c.HTML(
			http.StatusNoContent,
			"errors/error.html",
			gin.H{"error": err.Error()})
		return
	}
	c.HTML(
		http.StatusOK,
		"homepage.html",
		gin.H{
			"title":    "Home page",
			"articles": articles,
		})
}
