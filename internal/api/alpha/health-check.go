package alpha

import (
	"github.com/gin-gonic/gin"
)

// HealthCheck godoc
//	@Summary		Health Check
//	@Description	Check if the server is up and running
//	@Tags			Health Check
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	ApiResponse
//	@Router			/health-check [get]
func HealthCheck(c *gin.Context) {
	res := &ApiResponse{
		Code:    200,
		Message: "OK",
		Data:    nil,
	}

	res.Send(c)
}
