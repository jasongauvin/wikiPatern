package routes

import (
	"github.com/foolin/goview/supports/ginview"
	"github.com/gin-gonic/gin"
	"github.com/jasongauvin/wikiPattern/controllers"
	"github.com/jasongauvin/wikiPattern/middlewares"
	"github.com/gin-gonic/contrib/sessions"
)

func SetupRouter(router *gin.Engine) {
	// Sessions
	router.Use(sessions.Sessions("gosess", sessions.NewCookieStore([]byte("secret"))))
	//New template engine
	router.HTMLRender = ginview.Default()
	// Home
	router.GET("/", controllers.GetHomePage)

	// Articles
	router.GET("/articles", controllers.GetArticles)
	router.GET("/articles/:id", controllers.GetArticleById)
	router.GET("/new_article", controllers.GetNewArticleForm)
	router.POST("/articles", controllers.CreateArticle)
	router.GET("/edit_article/:id", controllers.GetEditArticleForm)
	router.POST("/edit_article/:id", controllers.EditArticleById)
	router.GET("/delete_article/:id", controllers.DeleteArticleById)

	// Comments
	router.POST("/comment", controllers.CreateComment)

	// Auth
	router.GET("/register", controllers.GetRegistrationForm)
	router.GET("/login", controllers.GetLoginForm)
	router.POST("/register", controllers.Registration)
	router.POST("/login", controllers.Login)

	// Auth required
	authorized := router.Group("/auth")
	authorized.Use(middlewares.CheckAuthorization)
	{
		authorized.GET("/profile", controllers.GetUserProfile)
	}

}
