package mongodb

import (
	"gitlab.com/pangold/goim/protocol"
	"time"
)

// message content
type Record struct {
	protocol.Message
	PC bool
	Android bool
	IOS bool
	Web bool
	CreatedTime time.Time
	UpdatedTime time.Time
}
