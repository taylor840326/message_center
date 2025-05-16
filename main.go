package main

import (
	"io/ioutil"
	"os"

	"message_center/plugins/input"
	"message_center/plugins/output"

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
	// botKey := ""
	// if botKey = os.Getenv("WECOM_ROBOT"); botKey == "" {
	// 	logger.Fatal("WECOM_ROBOT environment variable not set")
	// 	os.Exit(1)
	// }

	r := gin.Default()
	r.POST("/adapter/wx", func(c *gin.Context) {
		// body, _ := io.ReadAll(c.Request.Body)
		body, _ := ioutil.ReadFile("./plugins/input/prometheus/body.json")
		logger.Info("Received alert message body",
			zap.String("body", string(body)))

		// TODO: 解析消息体，发送到企业微信
		inputProcessor, err := input.GetInputProcessor(body, input.Prometheus)
		if err != nil {
			logger.Error("Failed to get input processor", zap.Error(err))
			c.JSON(500, "failed")
			return
		}
		messages, err := inputProcessor.ToMessages()
		if err != nil {
			logger.Error("Failed to get messages", zap.Error(err))
			c.JSON(500, "failed")
			return
		}
		outputProcessor, err := output.GetOutputProcessor(messages, output.WeCom)
		if err != nil {
			logger.Error("Failed to get output processor", zap.Error(err))
			c.JSON(500, "failed")
			return
		}
		outputProcessor.Output()

		c.JSON(200, "ok")
	})
	r.Run(":" + port) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
