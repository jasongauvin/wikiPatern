package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jasongauvin/wikiPattern/services"
	"net/http"
)

func GetArticles(c *gin.Context) {
	articles, err := services.LoadArticles()
	if err != nil {
		fmt.Println("Error: ", err)
		c.HTML(
			http.StatusNoContent,
			"errors/error.html",
			gin.H{"error": err.Error()})
		return
	}
	c.HTML(
		http.StatusOK,
		"article/index.html",
		gin.H{
			"title":   "Article index",
			"payload": articles,
		})
}

func GetArticleById(c *gin.Context) {
	article, err := services.LoadArticleById(c.Param("id"))
	if err != nil {
		fmt.Println("Error: ", err)
		c.HTML(
			http.StatusBadRequest,
			"errors/error.html",
			gin.H{"error": err.Error()})
		return
	}
	c.HTML(
		http.StatusOK,
		"article/show.html",
		gin.H{
			"title":         "Article Page",
			"article":       article,
			"comments":      article.Comments,
			"commentsCount": len(article.Comments),
		})
}

func CreateArticle(c *gin.Context) {
	switch c.Request.Method {
	case "GET":
		c.HTML(http.StatusOK,
			"article/new.html",
			gin.H{
				"title": "Article form",
				"url":   c.Request.URL.Path,
			})
	case "POST":
		var form services.ArticleForm
		if err := c.ShouldBind(&form); err != nil {
			fmt.Println("error:", err)
			c.HTML(
				http.StatusBadRequest,
				"errors/error.html",
				gin.H{"error": err.Error()})
			return
		}
		article, err := services.SaveArticle(form.Title, form.Content)
		if err != nil {
			fmt.Println("Error: ", err)
			c.AbortWithStatus(http.StatusUnprocessableEntity)
		}
		c.HTML(
			http.StatusCreated,
			"article/show.html",
			gin.H{
				"title":   "Article Page",
				"article": article,
			})
	default:
		c.AbortWithStatus(http.StatusMethodNotAllowed)
	}
}

func EditArticleById(c *gin.Context) {
	switch c.Request.Method {
	case "GET":
		article, err := services.LoadArticleById(c.Param("id"))
		if err != nil {
			fmt.Println("Error:", err)
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.HTML(
			http.StatusOK,
			"article/edit.html",
			gin.H{
				"title":   "Article edit form",
				"article": article,
			})
	case "POST":
		var form services.ArticleForm
		if err := c.ShouldBind(&form); err != nil {
			c.HTML(
				http.StatusBadRequest,
				"errors/error.html",
				gin.H{"error": err.Error()})
			return
		}
		article, err := services.EditArticle(c.Param("id"), form.Title, form.Content)
		if err != nil {
			fmt.Println("Error: ", err)
			c.AbortWithStatus(http.StatusUnprocessableEntity)
		}
		c.HTML(
			http.StatusOK,
			"article/show.html",
			gin.H{
				"title":   "Article Page",
				"article": article,
			})
	default:
		c.AbortWithStatus(http.StatusMethodNotAllowed)
	}
}

func DeleteArticleById(c *gin.Context) {
	err := services.DeleteArticle(c.Param("id"))
	if err != nil {
		fmt.Println("Error: ", err)
		c.AbortWithStatus(http.StatusNotModified)
	}
	c.Redirect(http.StatusMovedPermanently, "/articles")
}
