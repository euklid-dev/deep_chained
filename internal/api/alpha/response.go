package alpha

import (
	"github.com/gin-gonic/gin"
)

type ApiResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (r *ApiResponse) Send(c *gin.Context) {
	c.JSON(r.Code, r)
}

func SendError(c *gin.Context, code int, message string) {
	response := &ApiResponse{
		Code:    code,
		Message: message,
		Data:    nil,
	}
	response.Send(c)
}
