package interfaces

import (
	"gitlab.com/pangold/goim/protocol"
)

// Wrap grpc session synchronizing client of front service
type SessionAdapter interface {
	// get new session that connected to im service
	// mark down belonging's node.
	// GetSessionIn(func(*protocol.Session) error)
	SetSessionInHandler(func(*protocol.Session))
	// disconnected session that needs to be removed
	// mark down belonging's node
	// GetSessionOut(func(*protocol.Session) error)
	SetSessionOutHandler(func(*protocol.Session))
}

// Wrap grpc dispatch client of front service
type MessageAdapter interface {
	// get dispatched message from front im service
	// GetMessage(func(*protocol.Message) error)
	SetMessageHandler(func(*protocol.Message)error)
}

// Wrap grpc api client of front service
type FrontApiAdapter interface {
	//
	// GetNodeName() string
	// Send(message *protocol.Message)
	Send(message *protocol.Message) error
	SendEx(message []*protocol.Message) error
	// SendMessages(messages []*protocol.Message)
	Broadcast(message *protocol.Message) error
	BroadcastEx(message []*protocol.Message) error
	// BroadcastMessages(messages []*protocol.Message)
	GetConnections() []string
	Online(uid string) bool
	Kick(uid string) bool
}
