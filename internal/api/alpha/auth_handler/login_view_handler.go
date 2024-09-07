package login

import (
	"net/http"

	"github.com/euklid-dev/deep_chained/internal/views"
	"github.com/gin-gonic/gin"
)

func LoginViewHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "", views.LoginScreen())
}
