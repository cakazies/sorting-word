package cmd

import (
	"sorting-word/handler"

	"github.com/gin-gonic/gin"
)

func Run() {
	r := gin.Default()
	r.Use(gin.Recovery())

	r.POST("/api/sorting", handler.SortingMap)

	r.Run(":8080")
}
