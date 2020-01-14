package service

import (
	"errors"
	"gitlab.com/pangold/goim/protocol"
	"gitlab.com/pangold/goimb/config"
	"gitlab.com/pangold/goimb/model"
	"gitlab.com/pangold/goimb/front"
	"gitlab.com/pangold/goimb/history"
)

// See the design thoughts(Why Who How) written in history
// It needs to
// 1. query database/cache to get members of group
// 2. query database/cache to confirm their relationship(does it matter?)
//
// This module integrate adapter, database service, history, and conn together
// adapter: to receive
// database service: to get relationship if group message to know who to dispatch
// history: to save history message(for {expire} day)
// conn: to send(dispatch) message(s) out
//
type Service struct {
	// for business logic(service itself)
	Database *model.Database
	History history.History
	FrontApi *front.FrontApi
	// for api
	NotifyService *NotificationService
	FriendService *FriendService
	GroupService *GroupService
}

func NewService(conf config.ServiceConfig, frontApi *front.FrontApi) *Service {
	logic := &Service{
		Database: model.NewDatabase(conf.MySQL),
		History:  history.NewHistory(conf.MongoDB),
		FrontApi: frontApi,
	}
	logic.FrontApi.SetSessionInHandler(logic.on_session_in)
	logic.FrontApi.SetSessionOutHandler(logic.on_session_out)
	logic.FrontApi.SetMessageHandler(logic.on_message)
	logic.NotifyService = NewNotifierService(logic.Database, logic.History, logic.FrontApi)
	logic.FriendService = NewFriendService(logic.NotifyService, logic.Database.Friend)
	logic.GroupService = NewGroupService(logic.NotifyService, logic.Database.Group)
	return logic
}

// pull history message
func (this *Service) on_session_in(session *protocol.Session) {
	this.fetchHistory(session.UserId, session.ClientId)
	this.NotifyService.FriendOnline(session.UserId)
}

func (this *Service) on_session_out(session *protocol.Session) {
	this.NotifyService.FriendOffline(session.UserId)
}

func (this *Service) on_message(message *protocol.Message) error {
	switch message.Action {
	case ACTION_CHAT:
		return this.chat(message)
	}
	return errors.New("invalid action")
}

func (this *Service) fetchHistory(uid, cid string) {
	messages := this.History.Find(uid, cid)
	for _, message := range messages {
		this.NotifyService.SendWithoutHistory(message)
	}
}

func (this *Service) chat(message *protocol.Message) error {
	if message.TargetId != "" {
		return this.NotifyService.SendP2P(message)
	} else if message.GroupId != "" {
		return this.GroupService.SendGroup(message)
	}
	return errors.New("invalid chat message")
}