package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/jasongauvin/wikiPattern/services"
	"net/http"
)

// CheckAuthorization verify that the user got the session key.
func CheckAuthorization(c *gin.Context) {
	var err error
	_, err = services.CheckSessionExistence(c)
	if err != nil {
		c.Redirect(http.StatusPermanentRedirect, "/login")
		return
	}
	c.Next()
}
