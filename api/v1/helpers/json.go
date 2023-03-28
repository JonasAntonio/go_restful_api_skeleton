package helpers

import "github.com/gin-gonic/gin"

type Response struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func Respond(c *gin.Context, status int, message string, data ...any) {
	payload := Response{
		Error:   status > 399,
		Message: message,
	}

	if data != nil {
		payload.Data = data
	}

	c.JSON(status, payload)
}
