package simple

import (
	"gitlab.com/pangold/goim/protocol"
	"gitlab.com/pangold/goimb/config"
)

type SessionAdapter struct {

}

func NewSessionAdapter(conf config.Host) *SessionAdapter {
	return &SessionAdapter{

	}
}

func (this *SessionAdapter) SetSessionInHandler(handler func(*protocol.Session)) {

}

func (this *SessionAdapter) SetSessionOutHandler(handler func(*protocol.Session)) {

}
