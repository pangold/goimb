package service

import (
	"gitlab.com/pangold/goim/protocol"
	"gitlab.com/pangold/goimb/model"
	"gitlab.com/pangold/goimb/front"
	"gitlab.com/pangold/goimb/history"
	imb "gitlab.com/pangold/goimb/proto"
)

type NotificationService struct {
	database *model.Database
	history history.History
	frontApi *front.FrontApi
}

func NewNotifierService(db *model.Database, h history.History, a *front.FrontApi) *NotificationService {
	return &NotificationService{
		database: db,
		history: h,
		frontApi: a,
	}
}

func (this *NotificationService) SendWithoutHistory(message *protocol.Message) error {
	// message will be send properly by Send to
	// 1. its own targetId
	// 2. itself(the rest of clients)
	this.frontApi.Send(message, message.UserId)
	this.frontApi.Send(message, message.TargetId)
	return nil
}

func (this *NotificationService) SendP2P(message *protocol.Message) error {
	// message will be send properly by Send to
	// 1. its own targetId
	// 2. itself(the rest of clients)
	this.history.Add(message, this.frontApi.Send(message, message.UserId))
	this.history.Add(message, this.frontApi.Send(message, message.TargetId))
	return nil
}

func (this *NotificationService) FriendOnline(uid string) {
	resp := this.database.Friend.GetFriends(uid).(imb.Friends)
	for _, friend := range resp.Friends {
		// filter: online friends
		// FIXME: Not list, Map will be more efficient
		//sessions := this.frontApi.GetConnections()
		//for _, session := range sessions {
		//	if session == friend.UserId && uid != friend.UserId {
		//		this.online(uid, friend.UserId)
		//	}
		//}
		// if offline friend, send will be failure automatically
		if uid != friend.UserId {
			this.online(uid, friend.UserId)
		}
	}
}

func (this *NotificationService) FriendOffline(uid string) {
	resp := this.database.Friend.GetFriends(uid).(imb.Friends)
	for _, friend := range resp.Friends {
		// filter: tell friends
		// FIXME: Not list, Map will be more efficient
		//sessions := this.frontApi.GetConnections()
		//for _, session := range sessions {
		//	if session == friend.UserId && uid != friend.UserId {
		//		this.offline(uid, friend.UserId)
		//	}
		//}
		// FIXME: will be
		// if offline friend, send will be failure automatically
		if uid != friend.UserId {
			this.offline(uid, friend.UserId)
		}
	}
}

func (this *NotificationService) online(uid, tid string) {
	message := &protocol.Message{
		UserId: uid,
		TargetId: tid,
		Action: int32(ACTION_FRIEND_ONLINE),
	}
	this.frontApi.Send(message, uid)
	this.frontApi.Send(message, tid)
}

func (this *NotificationService) offline(uid, tid string) {
	message := &protocol.Message{
		UserId: uid,
		TargetId: tid,
		Action: int32(ACTION_FRIEND_OFFLINE),
	}
	this.frontApi.Send(message, uid)
	this.frontApi.Send(message, tid)
}

