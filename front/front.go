package front

import (
	"gitlab.com/pangold/goim/protocol"
	"gitlab.com/pangold/goimb/config"
	"gitlab.com/pangold/goimb/front/logic"
)

type FrontApi struct {
	message *logic.MessageAdapter
	session *logic.SessionAdapter
	front   *logic.FrontApiAdapter
}

func NewFrontApi(conf config.FrontConfig) *FrontApi {
	// FIXME: config for cluster
	session := logic.NewSessionAdapter(append([]config.Host{}, conf.SessionApi))
	return &FrontApi{
		message: logic.NewMessageAdapter(append([]config.Host{}, conf.Adapter.MessageDispatcher)),
		session: session,
		front:   logic.NewFrontApiAdapter(append([]config.Host{}, conf.FrontApi), session),
	}
}

func (this *FrontApi) SetMessageHandler(handler func(*protocol.Message) error) {
	this.message.SetMessageHandler(handler)
}

func (this *FrontApi) SetSessionInHandler(handler func(*protocol.Session)) {
	this.session.SetSessionInHandler(handler)
}

func (this *FrontApi) SetSessionOutHandler(handler func(*protocol.Session)) {
	this.session.SetSessionOutHandler(handler)
}

func (this *FrontApi) Send(message *protocol.Message, tid string) []string {
	return this.front.Send(message, tid)
}

func (this *FrontApi) SendEx(messages []*protocol.Message, tid string) []string {
	return this.front.SendEx(messages, tid)
}

func (this *FrontApi) Broadcast(message *protocol.Message) {
	this.front.Broadcast(message)
}

func (this *FrontApi) BroadcastEx(messages []*protocol.Message) {
	this.front.BroadcastEx(messages)
}

func (this *FrontApi) GetConnections() []string {
	return this.front.GetConnections()
}

func (this *FrontApi) Online(uid string) []string {
	return this.front.Online(uid)
}

func (this *FrontApi) Kick(uid string) bool {
	return this.front.Kick(uid)
}