package main

import (
	"bytes"
	"encoding/json"
	"io"
	"message_center/input/prometheus"
	"os"

	botApi "github.com/electricbubble/wecom-bot-api"
	"github.com/electricbubble/wecom-bot-api/md"
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

		alertMessage := prometheus.AlertMessage{}
		err := json.Unmarshal(body, &alertMessage)
		if err != nil {
			logger.Error("Failed to unmarshal alert message",
				zap.Error(err),
				zap.String("body", string(body)))
			c.JSON(500, err.Error())
		}

		bot := botApi.NewWeComBot(botKey)
		alerts := alertMessage.Alerts
		if len(alerts) == 0 {
			c.JSON(500, "no alerts found")
		}

		// Process alerts and format them into markdown message
		// Each alert contains annotations and labels with details about the alert
		// We'll format these into a readable message with namespace, node, and pod info
		content := bytes.NewBufferString("")
		for _, alert := range alerts {
			if len(content.String()) > 4000 {
				content.WriteString(md.QuoteText("更多: " + md.CommentText("更多信息请查看日志")))
				content.WriteString("\n")
				continue
			}
			content.WriteString("标题: " + alert.Annotations.Description + " \n")
			content.WriteString(md.QuoteText("命名空间: " + md.CommentText(alert.Labels.Namespace)))
			content.WriteString(md.QuoteText("服务器节点: " + md.CommentText(alert.Labels.Node)))
			content.WriteString(md.QuoteText("pod名称: " + md.CommentText(alert.Labels.Pod)))
			content.WriteString(md.QuoteText("状态: " + md.CommentText(alert.Status)))
			content.WriteString("\n")
		}
		logger.Info("Formatted markdown message content",
			zap.String("content", content.String()))
		err = bot.PushMarkdownMessage(content.String())
		if err != nil {
			logger.Error("Failed to send message to wecom",
				zap.Error(err),
			)
		}

		c.JSON(200, "ok")
	})
	r.Run(":" + port) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
