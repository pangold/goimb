package logic

import (
	"gitlab.com/pangold/goim/protocol"
	"gitlab.com/pangold/goimb/config"
	"gitlab.com/pangold/goimb/front/grpc"
	"gitlab.com/pangold/goimb/front/interfaces"
	"gitlab.com/pangold/goimb/front/simple"
	"log"
)

type FrontApiAdapter struct {
	adapters map[string]interfaces.FrontApiAdapter
	sessions *SessionAdapter
}

func NewFrontApiAdapter(configs []config.Host, session *SessionAdapter) *FrontApiAdapter {
	adapters := make(map[string]interfaces.FrontApiAdapter)
	for _, conf := range configs {
		if conf.Disabled {
			adapters[conf.Address] = simple.NewFrontApi(conf)
		} else {
			// grpc implement
			adapters[conf.Address] = grpc.NewFrontApi(conf)
		}
	}
	return &FrontApiAdapter {
		adapters: adapters,
		sessions: session,
	}
}

func (this *FrontApiAdapter) Send(message *protocol.Message, tid string) (res []string) {
	sessions := this.sessions.Find(tid)
	for _, session := range sessions {
		// get target front node that session located at
		if adapter, ok := this.adapters[session.NodeName]; ok {
			if err := adapter.Send(message); err != nil {
				log.Printf(err.Error())
				return res
			}
			res = append(res, session.ClientId)
		}
	}
	return res
}

func (this *FrontApiAdapter) SendEx(messages []*protocol.Message, tid string) (res []string) {
	sessions := this.sessions.Find(tid)
	for _, session := range sessions {
		// get target front node that session located at
		if adapter, ok := this.adapters[session.NodeName]; ok {
			if err := adapter.SendEx(messages); err != nil {
				log.Printf(err.Error())
				return res
			}
			res = append(res, session.ClientId)
		}
	}
	return res
}

func (this *FrontApiAdapter) Broadcast(message *protocol.Message) {
	for _, api := range this.adapters {
		if err := api.Broadcast(message); err != nil {
			log.Printf(err.Error())
			return
		}
	}
}

func (this *FrontApiAdapter) BroadcastEx(messages []*protocol.Message) {
	for _, api := range this.adapters {
		if err := api.BroadcastEx(messages); err != nil {
			log.Printf(err.Error())
			return
		}
	}
}

func (this *FrontApiAdapter) GetConnections() (res []string) {
	return this.sessions.GetOnlineUsers()
}

func (this *FrontApiAdapter) Online(uid string) (res []string) {
	sessions := this.sessions.Find(uid)
	for _, session := range sessions {
		res = append(res, session.ClientId)
	}
	return res
}

func (this *FrontApiAdapter) Kick(uid string) (res bool) {
	sessions := this.sessions.Find(uid)
	for _, session := range sessions {
		if adapter, ok := this.adapters[session.NodeName]; ok {
			res = adapter.Kick(uid)
		}
	}
	return res
}
