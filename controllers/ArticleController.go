package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jasongauvin/wikiPattern/models"
	"github.com/jasongauvin/wikiPattern/services"
	"github.com/jasongauvin/wikiPattern/strategies/export"
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

func GetArticleForm(c *gin.Context) {
	c.HTML(http.StatusOK,
		"article/new.html",
		gin.H{
			"title": "Article form",
			"url":   c.Request.URL.Path,
		})
}

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
}

func GetArticleEditForm(c *gin.Context) {
	var err error
	var article *models.Article
	article, err = services.LoadArticleById(c.Param("id"))
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
}

func EditArticleById(c *gin.Context) {
	var article *models.Article
	var err error
	var form services.ArticleForm
	if err := c.ShouldBind(&form); err != nil {
		c.HTML(
			http.StatusBadRequest,
			"errors/error.html",
			gin.H{"error": err.Error()})
		return
	}
	article, err = services.EditArticle(c.Param("id"), form.Title, form.Content)
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

func DeleteArticleById(c *gin.Context) {
	err := services.DeleteArticle(c.Param("id"))
	if err != nil {
		fmt.Println("Error: ", err)
		c.AbortWithStatus(http.StatusNotModified)
	}
	c.Redirect(http.StatusMovedPermanently, "/articles")
}

func ExportArticle(c *gin.Context) {
	queryParams := c.Request.URL.Query()
	exportFormat := queryParams["format"][0]

	csv := &export.Csv{}
	xlsx := &export.Xlsx{}
	var articleExportFile *export.ArticleExportFile
	var exportContext *export.ExportContext

	// Select export Format
	if exportFormat == "csv" {
		exportContext = export.InitExportContext(csv)
	} else if exportFormat == "xlsx" {
		exportContext = export.InitExportContext(xlsx)
	}
	articleExportFile = exportContext.Export(c.Param("id"))
	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Disposition", "attachment; filename="+articleExportFile.FileName)
	c.Data(http.StatusOK, articleExportFile.MimeType, articleExportFile.FileBytes)
	//c.Redirect(http.StatusTemporaryRedirect, "/articles/"+c.Param("id"))
}
