package wecom

import (
	"bytes"
	"message_center/input/prometheus"

	botApi "github.com/electricbubble/wecom-bot-api"
	"github.com/electricbubble/wecom-bot-api/md"
)

func SendWeComMessage(message prometheus.AlertMessage) {
	botKey := "a15aeacd-2f81-49e7-ad90-2293f1a086d5" // 只填 key= 后边的内容

	bot := botApi.NewWeComBot(botKey)

	alerts := message.Alerts
	for _, alertItem := range alerts {
		labels := alertItem.Labels
		print(labels)
		content := bytes.NewBufferString(md.Heading(1, "数据库报警"))
		content.WriteString("实时新增用户反馈" + md.WarningText("132例") + "，请相关同事注意。\n")

	}
	// 👆效果等同于👇
	/*
		# H1
		实时新增用户反馈 <font color="warning">132例</font>，请相关同事注意。\n
		> 类型:<font color="comment">用户反馈</font>
		> 普通用户反馈:<font color="comment">117例</font>
		> VIP用户反馈:<font color="comment">15例</font>
	*/

	// 仅发送 `markdown` 格式的文本
	_ = bot.PushMarkdownMessage(content.String())

}
