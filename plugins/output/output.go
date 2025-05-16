package output

import (
	"message_center/message"
	"message_center/plugins/output/wecom"
)

type OutputProcessor interface {
	Output() error
}

const (
	WeCom = "wecom"
	Text  = "text"
)

// GetInputProcessor 根据格式获取对应的渲染器
func GetOutputProcessor(messages []message.Message, ptype string) (OutputProcessor, error) {
	switch ptype {
	case WeCom:
		return wecom.NewWeCom(messages)
	default:
		return wecom.NewWeCom(messages)
	}
}
