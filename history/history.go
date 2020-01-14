package history

import (
	"gitlab.com/pangold/goim/protocol"
	"gitlab.com/pangold/goimb/config"
	"gitlab.com/pangold/goimb/history/mongodb"
	"gitlab.com/pangold/goimb/history/simple"
)

type History interface {
	// #### Why ####
	// There are 4 kinds of client for each users: PC, Android, IOS, Web(H5)
	// Reason 1:
	// Sync to 4 kinds of client
	// Each messages need to dispatch to all kinds of client
	// But not all the kinds are online
	// So, messages needs to be stored and dispatches when the client login.
	// ------------------
	// Reason 2:
	// To fetch historical messages
	// ==================
	//
	// #### Who ####
	// 1. dispatches to targets(all kinds of client)
	// 2. dispatches to yourself(other kinds of client)
	// ==================
	//
	// #### How ####
	// Situation 1: p2p message
	// From one user
	// To one user
	// ------------------
	// Situation 2: group message
	// From one user
	// To many users
	//     query database or cache to get members(users)
	//     1. clone message for each members
	//     2. set targetId field to member's id(targetId is empty originally)
	// ------------------
	// Condition: check if groupId is empty?
	//     T: p2p message(targetId must not be empty)
	//     F: group message
	// ==================
	//
	// Finally, each messages need to be stored via this func(Push)
	Add(message *protocol.Message, cids []string)
	// For synchronizing to specific client(when session in)
	//     Where is the beginning index of the list for each kind of clients?
	//     [inside the implements, invokers don't need to know about that]
	Find(uid, cid string) []*protocol.Message
	// Check if record is history,
	// Group's application message, only one needs to handle it
	Hit(uid, tid, gid string, action int32) bool
}

func NewHistory(conf config.Host) History {
	if conf.Disabled {
		return simple.NewSimple()
	}
	return mongodb.NewMongoDB(conf)
}
