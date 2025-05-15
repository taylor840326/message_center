package wecom

import (
	"bytes"

	botApi "github.com/electricbubble/wecom-bot-api"
)

func SendWecomMessage() {
	bot := botApi.NewWeComBot("")
	alerts := ""
	if len(alerts) == 0 {
	}

	// Process alerts and format them into markdown message
	// Each alert contains annotations and labels with details about the alert
	// We'll format these into a readable message with namespace, node, and pod info
	content := bytes.NewBufferString("")
	// for _, alert := range alerts {
	// 	if len(content.String()) > 4000 {
	// 		content.WriteString(md.QuoteText("更多: " + md.CommentText("更多信息请查看日志")))
	// 		content.WriteString("\n")
	// 		continue
	// 	}
	// content.WriteString("标题: " + alert.Annotations.Description + " \n")
	// content.WriteString(md.QuoteText("命名空间: " + md.CommentText(alert.Labels.Namespace)))
	// content.WriteString(md.QuoteText("服务器节点: " + md.CommentText(alert.Labels.Node)))
	// content.WriteString(md.QuoteText("pod名称: " + md.CommentText(alert.Labels.Pod)))
	// switch alert.Status {
	// case "firing":
	// 	content.WriteString(md.QuoteText("状态: " + md.WarningText(alert.Status)))
	// case "resolved":
	// 	content.WriteString(md.QuoteText("状态: " + md.InfoText(alert.Status)))
	// default:
	// 	content.WriteString(md.QuoteText("状态: " + md.CommentText(alert.Status)))
	// }
	// 	content.WriteString("\n")
	// }
	_ = bot.PushMarkdownMessage(content.String())

}
