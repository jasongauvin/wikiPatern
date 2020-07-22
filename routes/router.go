package routes

import (
	"github.com/foolin/goview/supports/ginview"
	"github.com/gin-gonic/gin"
	"github.com/jasongauvin/wikiPattern/controllers"
	"net/http"
)


func SetupRouter(router *gin.Engine) {
	//new template engine
	router.HTMLRender = ginview.Default()
	router.LoadHTMLGlob("views/*")

	router.GET("/", func(c *gin.Context){
		c.JSON(200, gin.H{
			"message": "Poil au message",
		})
	})

	router.GET("/articles", func(ctx *gin.Context) {
		//render only file, must full name with extension
		ctx.HTML(http.StatusOK, "article/index.html", gin.H{
			"title": "Home Page",
			"articles": controllers.GetArticles,
		})
	})
}