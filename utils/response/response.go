package utils

import (
	"github.com/gin-gonic/gin"
)

type Resp struct {
	Message string      `json:"message"`
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
}

func Response(c *gin.Context, message string, code int, data interface{}) {
	c.Header("Content-Type", "application/json")
	c.Set("responseBody", data)
	c.JSON(code, Resp{
		Message: message,
		Data:    data,
	})
}
