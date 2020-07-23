package services

import (
	"fmt"
	"github.com/jasongauvin/wikiPattern/models"
	"time"
)

type ArticleForm struct {
	Title   string `form:"articleTitle" binding:"required"`
	Content string `form:"articleContent" binding:"required"`
	//Published bool   `form:"articlePublish"`
}

func LoadArticles() (*[]models.Article, error) {
	var articles []models.Article
	var err error
	articles, err = models.FindArticles()

	if err != nil {
		fmt.Println("Error: ", err)
		return nil, err
	}
	return &articles, nil
}

func LoadArticleByName(name string) (*models.Article, error) {
	var article models.Article
	var err error
	article, err = models.FindArticleByName(name)
	if err != nil {
		fmt.Println("Error: ", err)
		return nil, err
	}
	return &article, nil
}

func LoadArticleById(id string) (*models.Article, error) {
	convertedId := ConvertStringToInt(id)
	var article models.Article
	var err error
	article, err = models.FindArticleByID(convertedId)
	if err != nil {
		fmt.Println("Error: ", err)
		return nil, err
	}
	return &article, nil
}

func SaveArticle(title string, content string) (*models.Article, error) {
	var article models.Article
	article.Title = title
	article.Content = content
	article.CreatedAt = time.Now()
	err := models.CreateArticle(&article)
	if err != nil {
		return nil, err
	}
	return &article, nil
}

func EditArticle(id string, title string, content string) (*models.Article, error) {
	convertedId := ConvertStringToInt(id)
	var article models.Article
	article.Title = title
	article.Content = content
	article.UpdatedAt = time.Now()
	err := models.EditArticleByID(&article, convertedId)
	if err != nil {
		return nil, err
	}
	return &article, nil
}

func DeleteArticle(id string) error {
	convertedId := ConvertStringToInt(id)
	var err error
	err = models.DeleteArticleByID(convertedId)
	if err != nil {
		return err
	}
	return nil
}

func LoadArticleByOrderedDate() (*[]models.Article, error) {
	var articles []models.Article
	var err error
	articles, err = models.FindArticleByOrderedDate()
	if err != nil {
		return nil, err
	}
	return &articles, nil
}
