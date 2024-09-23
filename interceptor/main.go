package main

import (
	"bytes"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Notification struct {
	ID     string `json:"id"`
	Status string `json:"status"`
}

func main() {
	r := gin.Default()
	r.Use(LogMiddleware())
	r.POST("/notification", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	})
	r.Run(":3000") // 8080
}

func LogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		bodyBytes, err := io.ReadAll(c.Request.Body)
		if err != nil {
			log.Println("error reading body", err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		log.Println("request body:", string(bodyBytes))

		c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

		c.Next()
	}
}
