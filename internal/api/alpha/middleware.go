package alpha

import "github.com/gin-gonic/gin"

var dev = true

func DisableCacheInDevMode(c *gin.Context) {
	if dev {

		c.Header("Cache-Control", "no-store")
	}
	c.Next()

}
