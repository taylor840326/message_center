package input

import (
	"message_center/message"
	"message_center/plugins/input/prometheus"
)

type InputProcessor interface {
	ToString() (string, error)
	ToMessages() ([]message.Message, error)
}

const (
	Prometheus = "prometheus"
	Text       = "text"
)

// GetInputProcessor 根据格式获取对应的渲染器
func GetInputProcessor(content []byte, ptype string) (InputProcessor, error) {
	switch ptype {
	case Prometheus:
		return prometheus.NewPrometheusInputProcessor(content)
	default:
		return prometheus.NewPrometheusInputProcessor(content)
	}
}
