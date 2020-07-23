package services

import (
	"github.com/jasongauvin/wikiPattern/models"
	"time"
)

type CommentForm struct {
	Content   string `form:"commentText" binding:"required"`
	ArticleId uint64 `form:"articleId"`
}

func SaveComment(content string, articleId uint64) (*models.Comment, error) {
	var comment models.Comment
	comment.Text = content
	comment.ArticleId = articleId
	comment.CreatedAt = time.Now()
	err := models.CreateComment(&comment)
	if err != nil {
		return nil, err
	}
	return &comment, nil
}
