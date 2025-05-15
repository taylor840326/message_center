package wecom

import (
	"bytes"
	"message_center/plugins/input/prometheus"

	botApi "github.com/electricbubble/wecom-bot-api"
	"github.com/electricbubble/wecom-bot-api/md"
)

func SendWecomMessage(msg prometheus.AlertMessage) {
	botKey := "WeCom_Bot_Key" // 只填 key= 后边的内容

	bot := botApi.NewWeComBot(botKey)

	content := bytes.NewBufferString(md.Heading(1, "H1"))
	content.WriteString("实时新增用户反馈" + md.WarningText("132例") + "，请相关同事注意。\n")
	content.WriteString(md.QuoteText("类型:" + md.CommentText("用户反馈")))
	content.WriteString(md.QuoteText("普通用户反馈:" + md.CommentText("117例")))
	content.WriteString(md.QuoteText("VIP用户反馈:" + md.CommentText("15例")))
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
