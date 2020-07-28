package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/jasongauvin/wikiPattern/services"
	"net/http"
)

func CheckAuthorization(c *gin.Context) {
	cookie, err := c.Cookie("session_token")
	if err != nil {
		if err == http.ErrNoCookie {
			c.Redirect(http.StatusPermanentRedirect, "/login")
			c.Abort()
			return
		}
		c.Redirect(http.StatusPermanentRedirect, "/login")
		c.Abort()
		return
	}
	_, err = services.CheckSessionExistence(cookie)
	if err != nil {
		c.Redirect(http.StatusPermanentRedirect, "/login")
		c.Abort()
		return
	}
	c.Next()
}
