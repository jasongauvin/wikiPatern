package routes

import (
	"github.com/foolin/goview/supports/ginview"
	"github.com/gin-gonic/gin"
	"github.com/gobuffalo/packr/v2"
	"github.com/jasongauvin/wikiPattern/controllers"
	"net/http"
)

func loadTemplate(templatePath string) string {
	box := packr.New("templates", "./views")
	html, err := box.FindString(templatePath + ".html")
	if err != nil {
		return ""
	}

	return html
}

func SetupRouter(router *gin.Engine) {
	//new template engine
	router.HTMLRender = ginview.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Poil au message",
		})
	})

	router.GET("/articles", func(c *gin.Context) {
		c.HTML(http.StatusOK,
			"article/index.html",
			gin.H{
				"title":   "Home Page",
				"payload": controllers.GetAllArticles(),
			})
	})
}
