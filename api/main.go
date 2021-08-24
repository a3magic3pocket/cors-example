package main

import (
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(cors.New(
		cors.Config{
			AllowOrigins:     []string{"http://localhost:8888"},
			AllowMethods:     []string{"POST"},
			AllowHeaders:     []string{"Origin", "custom-header"},
			AllowCredentials: true,
			MaxAge:           12 * time.Hour,
		}))

	router.POST("test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "success"})
	})

	router.Run(":8889")
}
