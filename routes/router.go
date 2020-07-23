package routes

import (
	"github.com/foolin/goview/supports/ginview"
	"github.com/gin-gonic/gin"
	"github.com/jasongauvin/wikiPattern/controllers"
)

func SetupRouter(router *gin.Engine) {
	//new template engine
	router.HTMLRender = ginview.Default()

	router.GET("/", controllers.GetHomePage)
	router.GET("/articles", controllers.GetArticles)
	router.GET("/articles/:id", controllers.GetArticleById)
	router.GET("/new_article", controllers.CreateArticle)
	router.POST("/articles", controllers.CreateArticle)
	router.GET("/edit_article/:id", controllers.EditArticleById)
	router.POST("/edit_article/:id", controllers.EditArticleById)
	router.GET("/delete_article/:id", controllers.DeleteArticleById)
	router.POST("/comment", controllers.CreateComment)
}
