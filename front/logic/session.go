package logic

import (
	"gitlab.com/pangold/goim/protocol"
	"gitlab.com/pangold/goimb/config"
	"gitlab.com/pangold/goimb/front/grpc"
	"gitlab.com/pangold/goimb/front/interfaces"
	"gitlab.com/pangold/goimb/front/simple"
)

// A session pool that in only one single node.
// to process session together.
type SessionAdapter struct {
	adapters []interfaces.SessionAdapter
	sessions map[string]map[string]*protocol.Session
	sessionIn *func(*protocol.Session)
	sessionOut *func(*protocol.Session)
}

func NewSessionAdapter(configs []config.Host) *SessionAdapter {
	var adapters []interfaces.SessionAdapter
	for _, conf := range configs {
		if conf.Disabled {
			adapters = append(adapters, simple.NewSessionAdapter(conf))
		} else {
			// grpc implement
			adapters = append(adapters, grpc.NewSessionAdapter(conf))
		}
	}
	session := &SessionAdapter{
		adapters: adapters,
		sessions: make(map[string]map[string]*protocol.Session),
		sessionIn: nil,
		sessionOut: nil,
	}
	for _, adapter := range session.adapters {
		adapter.SetSessionInHandler(session.on_session_in)
		adapter.SetSessionOutHandler(session.on_session_out)
	}
	return session
}

func (this *SessionAdapter) SetSessionInHandler(handler func(*protocol.Session)) {
	this.sessionIn = &handler
}

func (this *SessionAdapter) SetSessionOutHandler(handler func(*protocol.Session)) {
	this.sessionOut = &handler
}

// get all connecting session, include all kinds of client
func (this *SessionAdapter) GetSessionList() (res []*protocol.Session) {
	for _, sessions := range this.sessions {
		for _, session := range sessions {
			res = append(res, session)
		}
	}
	return res
}

// get all connecting session, one user one session
func (this *SessionAdapter) GetOnlineUsers() (res []string) {
	for _, sessions := range this.sessions {
		// get first
		for _, session := range sessions {
			res = append(res, session.UserId)
			break
		}
	}
	return res
}

func (this *SessionAdapter) Find(uid string) (res []*protocol.Session) {
	if sessions, ok := this.sessions[uid]; ok {
		for _, session := range sessions {
			res = append(res, session)
		}
	}
	return res
}

func (this *SessionAdapter) on_session_in(session *protocol.Session) {
	s, ok := this.sessions[session.UserId]
	if !ok {
		s = make(map[string]*protocol.Session)
	}
	s[session.ClientId] = session
	if this.sessionIn != nil {
		(*this.sessionIn)(session)
	}
}

func (this *SessionAdapter) on_session_out(session *protocol.Session) {
	if sss, ok := this.sessions[session.UserId]; ok {
		if _, ok2 := sss[session.ClientId]; ok2 {
			delete(sss, session.ClientId)
		}
		if len(sss) == 0 {
			delete(this.sessions, session.UserId)
		}
	}
	if this.sessionOut != nil {
		(*this.sessionOut)(session)
	}
}

