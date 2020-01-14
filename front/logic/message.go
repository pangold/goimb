package logic

import (
	"gitlab.com/pangold/goim/protocol"
	"gitlab.com/pangold/goimb/config"
	"gitlab.com/pangold/goimb/front/grpc"
	"gitlab.com/pangold/goimb/front/interfaces"
	"gitlab.com/pangold/goimb/front/simple"
)

type MessageAdapter struct {
	adapters []interfaces.MessageAdapter
}

func NewMessageAdapter(configs []config.Host) *MessageAdapter {
	var adapters []interfaces.MessageAdapter
	for _, conf := range configs {
		if conf.Disabled {
			adapters = append(adapters, simple.NewMessageAdapter(conf))
		} else {
			// grpc implement
			adapters = append(adapters, grpc.NewMessageAdapter(conf))
		}
	}
	return &MessageAdapter {
		adapters: adapters,
	}
}

func (this *MessageAdapter) SetMessageHandler(handler func(*protocol.Message) error) {
	for _, adapter := range this.adapters {
		adapter.SetMessageHandler(handler)
	}
}
