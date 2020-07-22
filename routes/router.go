package routes

import (
	"github.com/foolin/goview/supports/ginview"
	"github.com/gin-gonic/gin"
	"github.com/jasongauvin/wikiPattern/controllers"
)

func SetupRouter(router *gin.Engine) {
	//new template engine
	router.HTMLRender = ginview.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Poil au message",
		})
	})

	router.GET("/articles", controllers.GetArticles)
	router.GET("/articles/:id", controllers.GetArticleById)
}