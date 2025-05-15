package main

import (
	"io"
	"os"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {

	gin.SetMode(gin.DebugMode)
	port := ""
	if port = os.Getenv("HTTP_PORT"); port == "" {
		port = "80"
	}

	// Initialize logger
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	// Check if WECOM_ROBOT environment variable is set
	botKey := ""
	if botKey = os.Getenv("WECOM_ROBOT"); botKey == "" {
		logger.Fatal("WECOM_ROBOT environment variable not set")
		os.Exit(1)
	}

	r := gin.Default()
	r.POST("/adapter/wx", func(c *gin.Context) {
		body, _ := io.ReadAll(c.Request.Body)
		logger.Info("Received alert message body",
			zap.String("body", string(body)))

		c.JSON(200, "ok")
	})
	r.Run(":" + port) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
