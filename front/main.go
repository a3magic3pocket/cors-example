package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.LoadHTMLGlob("view/*")
	router.GET("", func(c *gin.Context) {
		c.HTML(http.StatusOK, "test.html", gin.H{})
	})
	router.Run(":8888")
}
