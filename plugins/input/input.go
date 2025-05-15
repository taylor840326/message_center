package input

import "message_center/message"

type Input interface {
	ToString() (string, error)
	ToMessage() (message.Message, error)
}
