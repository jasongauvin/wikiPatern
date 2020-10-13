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
	articlesView := router.Group("articles")

	articlesView.GET("/", controllers.GetArticles)
	articlesView.GET("/:id", controllers.GetArticleById)
	// Articles Forms ----
	articleForms := router.Group("/")
	articleForms.GET("/new", controllers.GetArticleForm)
	articleForms.POST("/new-article", controllers.CreateArticle)
	articleForms.GET("/edit/:id", controllers.GetArticleEditForm)
	articleForms.POST("/edit-article/:id", controllers.EditArticleById)
	router.GET("/delete-article/:id", controllers.DeleteArticleById)

	// Export article
	router.GET("/export-article/:id", controllers.ExportArticle)
	// Comments----
	router.POST("/comment", controllers.CreateComment)
	// ----Comments

}
