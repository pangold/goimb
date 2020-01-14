package simple

import (
	"gitlab.com/pangold/goim/protocol"
	"gitlab.com/pangold/goimb/config"
)

type MessageAdapter struct {

}

func NewMessageAdapter(conf config.Host) *MessageAdapter {
	return &MessageAdapter{

	}
}

func (this *MessageAdapter) SetMessageHandler(handler func(*protocol.Message) error) {

}