package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jasongauvin/wikiPattern/services"
	"net/http"
	"strconv"
)

func CreateComment(c *gin.Context) {
	var form services.CommentForm
	if err := c.ShouldBind(&form); err != nil {
		fmt.Println("error:", err)
		c.HTML(
			http.StatusBadRequest,
			"errors/error.html",
			gin.H{"error": err.Error()})
		return
	}
	comment, err := services.SaveComment(form.Content, form.ArticleId)
	if err != nil {
		fmt.Println("Error: ", err)
		c.AbortWithStatus(http.StatusUnprocessableEntity)
	}
	c.Redirect(http.StatusMovedPermanently, "/articles/"+strconv.FormatUint(comment.ArticleId, 10))
}
