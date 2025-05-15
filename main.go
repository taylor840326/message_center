package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"message_center/input/prometheus"

	botApi "github.com/electricbubble/wecom-bot-api"
	"github.com/electricbubble/wecom-bot-api/md"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	r.POST("/adapter/wx", func(c *gin.Context) {
		body, _ := io.ReadAll(c.Request.Body)
		fmt.Println(string(body))

		alertMessage := prometheus.AlertMessage{}
		err := json.Unmarshal(body, &alertMessage)
		if err != nil {
			fmt.Println(err.Error())
			c.JSON(500, err.Error())
		}

		botKey := "a15aeacd-2f81-49e7-ad90-2293f1a086d5" // 只填 key= 后边的内容
		bot := botApi.NewWeComBot(botKey)
		alerts := alertMessage.Alerts

		content := bytes.NewBufferString("")
		if len(alerts) == 0 {
			c.JSON(500, "收到的报警消息内容为空")
		}

		for _, alert := range alerts {

			content.WriteString("标题: " + alert.Annotations.Summary + " \n")
			content.WriteString(md.QuoteText("命名空间:" + md.CommentText(alert.Labels.Namespace)))
			content.WriteString(md.QuoteText("服务器节点:" + md.CommentText(alert.Labels.Node)))
			content.WriteString(md.QuoteText("pod名称:" + md.CommentText(alert.Labels.Pod)))
			content.WriteString("\n")
		}
		_ = bot.PushMarkdownMessage(content.String())
		c.JSON(200, "ok")
	})
	r.Run(":80") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
