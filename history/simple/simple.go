package simple

import (
	"gitlab.com/pangold/goim/protocol"
)

type Simple struct {

}

func NewSimple() *Simple {
	return &Simple {

	}
}

// cid could be cids
func (this *Simple) Add(message *protocol.Message, cids []string) {

}

func (this *Simple) Find(uid, cid string) (res []*protocol.Message) {
	return res
}

func (this *Simple) Hit(uid, tid, gid string, action int32) bool {
	return true
}
