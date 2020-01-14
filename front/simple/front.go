package simple

import (
	"gitlab.com/pangold/goim/protocol"
	"gitlab.com/pangold/goimb/config"
)

type FrontApiAdapter struct {

}

func NewFrontApi(conf config.Host) *FrontApiAdapter {
	front := &FrontApiAdapter{

	}
	return front
}

func (this *FrontApiAdapter) Send(message *protocol.Message) error {
	return nil
}

func (this *FrontApiAdapter) SendEx(messages []*protocol.Message) error {
	return nil
}

func (this *FrontApiAdapter) Broadcast(message *protocol.Message) error {
	return nil
}

func (this *FrontApiAdapter) BroadcastEx(messages []*protocol.Message) error {
	return nil
}

func (this *FrontApiAdapter) GetConnections() []string {
	return nil
}

func (this *FrontApiAdapter) Online(uid string) bool {
	return true
}

func (this *FrontApiAdapter) Kick(uid string) bool {
	return true
}

