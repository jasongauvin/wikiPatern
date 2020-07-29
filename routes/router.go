package routes

import (
	"github.com/foolin/goview/supports/ginview"
	"github.com/gin-gonic/gin"
	"github.com/jasongauvin/wikiPattern/controllers"
)

func SetupRouter(router *gin.Engine) {
	//new template engine
	router.HTMLRender = ginview.Default()

	// Home----
	router.GET("/", controllers.GetHomePage)
	// ----HOME

	// Articles----
	router.GET("/articles", controllers.GetArticles)
	router.GET("/articles/:id", controllers.GetArticleById)
	router.GET("/articles/create", controllers.GetArticleForm)
	router.POST("/articles/create", controllers.CreateArticle)
	router.GET("/edit_article/:id", controllers.GetArticleEditForm)
	router.POST("/edit_article/:id", controllers.EditArticleById)
	router.GET("/delete_article/:id", controllers.DeleteArticleById)
	// Export article
	router.GET("/export/:id", controllers.ExportArticle)
	// ----Articles

	// Comments----
	router.POST("/comment", controllers.CreateComment)
	// ----Comments

}
