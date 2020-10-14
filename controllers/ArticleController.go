package controllers

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jasongauvin/wikiPattern/models"
	"github.com/jasongauvin/wikiPattern/services"
	"github.com/jasongauvin/wikiPattern/validators"
	"net/http"
)

// GetArticles render all the articles
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

// GetArticleById render the article from its id.
func GetArticleById(c *gin.Context) {
	article, err := services.LoadArticleById(c.Param("id"))
	if err != nil {
		err = errors.New("No article found")
		c.HTML(
			http.StatusNoContent,
			"errors/error.html",
			gin.H{
				"error": err,
			})
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

// GetNewArticleForm render the view to create an article
func GetNewArticleForm(c *gin.Context)  {
	tags, err := models.FindTags()
	if err != nil {
		c.HTML(
			http.StatusNoContent,
			"errors/error.html",
			gin.H{
				"error": errors.New("No article found"),
			})
	}

	c.HTML(http.StatusOK,
		"article/new.html",
		gin.H{
			"title": "Article form",
			"url":   c.Request.URL.Path,
			"tags": tags,
		})
}


// CreateArticle gets the posted form, bind it with model and saves it in db.
func CreateArticle(c *gin.Context) {
	var form services.ArticleForm
	if err := c.ShouldBind(&form); err != nil {
		fmt.Println("error:", err)
		c.HTML(
			http.StatusBadRequest,
			"errors/error.html",
			gin.H{"error": err.Error()})
		return
	}
	article, err := services.SaveArticle(form.Title, form.Content, form.Tags)
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
}

// GetEditArticleForm renders the view to edit an article
func GetEditArticleForm(c *gin.Context)  {
	article, err := services.LoadArticleById(c.Param("id"))
	if err != nil {
		fmt.Println("Error:", err)
		c.HTML(
			http.StatusBadRequest,
			"errors/error.html",
			gin.H{"error": err.Error()})
		return
	}
	c.HTML(
		http.StatusOK,
		"article/edit.html",
		gin.H{
			"title":   "Article edit form",
			"article": article,
		})
}

// EditArticleById gets the edit article form bind it with model, and update the article in db.
func EditArticleById(c *gin.Context) {
	var form services.ArticleForm
	err := validators.ValidateArticle(&form)
	if err != nil {
		fmt.Println("error:", err)
		c.HTML(
			http.StatusBadRequest,
			"errors/error.html",
			gin.H{"error": err.Error()})
		return
	}
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
}

// DeleteArticleById remove the article from db.
func DeleteArticleById(c *gin.Context) {
	err := services.DeleteArticle(c.Param("id"))
	if err != nil {
		fmt.Println("Error: ", err)
		c.AbortWithStatus(http.StatusNotModified)
	}
	c.Redirect(http.StatusMovedPermanently, "/articles")
}
