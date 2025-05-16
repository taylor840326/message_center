package wecom

import (
	"bytes"
	"message_center/message"
	"message_center/template"

	botApi "github.com/electricbubble/wecom-bot-api"
)

type WeCom struct {
	RobotKey string            `json:"robot_key"`
	Messages []message.Message `json:"messages"`
}

func NewWeCom(messages []message.Message) (*WeCom, error) {
	return &WeCom{
		RobotKey: "a15aeacd-2f81-49e7-ad90-2293f1a086d5",
		Messages: messages,
	}, nil
}

func (wecom *WeCom) Output() error {
	bot := botApi.NewWeComBot(wecom.RobotKey)

	// 企业微信消息模板
	const wecomTemplate = `标题: {{.Description.Default}} 
		> 命名空间: <font color="comment">{{.Labels.namespace.Default}}</font>
		> 服务器节点: <font color="comment">{{.Labels.node.Default}}</font>
		> pod名称: <font color="comment">{{.Labels.pod.Default}}</font>
		> 状态: <font color="info">{{.Status.Default}}</font>
		> 严重程度: <font color="info">{{.Severity.Default}}</font>
		`

	content := bytes.NewBufferString("")
	for _, message := range wecom.Messages {
		// 获取文本渲染器（因为企业微信消息实际上是文本格式）
		renderer, err := template.GetRenderer(wecomTemplate, template.FormatText)
		if err != nil {
			return err
		}

		// 渲染消息
		result, err := renderer.Render(message)
		if err != nil {
			return err
		}
		content.WriteString(result)
		content.WriteString("\n")
	}
	_ = bot.PushMarkdownMessage(content.String())

	return nil
}
