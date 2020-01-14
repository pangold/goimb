package service

import (
	"github.com/golang/protobuf/proto"
	"gitlab.com/pangold/goim/protocol"
	"gitlab.com/pangold/goimb/model"
	"gitlab.com/pangold/goimb/proto"
	"gitlab.com/pangold/goimb/utils"
)

type FriendService struct {
	database model.FriendDB
	notifier *NotificationService
}

func NewFriendService(n *NotificationService, d model.FriendDB) *FriendService {
	return &FriendService{
		database: d,
		notifier: n,
	}
}

func (this *FriendService) GetFriendInfo(request *imb.QueryFriendRequest) *imb.UserInfo {
	return this.database.GetFriendInfo(request.UserId).(*imb.UserInfo)
}

func (this *FriendService) GetFriends(request *imb.QueryFriendRequest) *imb.Friends {
	return this.database.GetFriends(request.UserId).(*imb.Friends)
}

func (this *FriendService) FriendMakingApply(request *imb.FriendMakingRequest) *imb.FriendResponse {
	message := &protocol.Message{
		Id:       utils.GenerateMessageId(),
		UserId:   request.UserId,
		TargetId: request.TargetId,
		Action:   int32(ACTION_FRIEND_APPLIED),
		Body:     []byte(request.Postscript),
	}
	if err := this.notifier.SendP2P(message); err != nil {
		return this.failure(CODE_FRIEND_NOTIFY_MAKING_APPLIED_ERROR, err.Error())
	}
	return this.success()
}

func (this *FriendService) FriendMakeAccept(request *imb.FriendMakingRequest) *imb.FriendResponse {
	message := &protocol.Message{
		Id:       utils.GenerateMessageId(),
		UserId:   request.UserId,
		TargetId: request.TargetId,
		Action:   int32(ACTION_FRIEND_ACCEPTED),
		Body:     []byte(request.Postscript),
	}
	if err := this.database.FriendCreate(message.UserId, message.TargetId); err != nil {
		return this.failure(CODE_FRIEND_MAKING_ERROR, err.Error())
	}
	if err := this.notifier.SendP2P(message); err != nil {
		return this.failure(CODE_FRIEND_NOTIFY_MAKING_ACCEPTED_ERROR, err.Error())
	}
	return this.success()
}

func (this *FriendService) FriendMakeReject(request *imb.FriendMakingRequest) *imb.FriendResponse {
	message := &protocol.Message{
		Id:       utils.GenerateMessageId(),
		UserId:   request.UserId,
		TargetId: request.TargetId,
		Action:   int32(ACTION_FRIEND_REJECTED),
		Body:     []byte(request.Postscript),
	}
	if err := this.notifier.SendP2P(message); err != nil {
		return this.failure(CODE_FRIEND_NOTIFY_MAKING_REJECTED_ERROR, err.Error())
	}
	return this.success()
}

func (this *FriendService) FriendBreakup(request *imb.FriendBreakupRequest) *imb.FriendResponse {
	// TODO: check if userId and targetId are friends
	message := &protocol.Message{
		Id:       utils.GenerateMessageId(),
		UserId:   request.UserId,
		TargetId: request.TargetId,
		Action:   int32(ACTION_FRIEND_BROKEUP),
		Body:     []byte(request.Postscript),
	}
	if err := this.database.FriendDelete(message.UserId, message.TargetId); err != nil {
		return this.failure(CODE_FRIEND_BREAKUP_ERROR, err.Error())
	}
	if err := this.notifier.SendP2P(message); err != nil {
		return this.failure(CODE_FRIEND_NOTIFY_BROKEUP_ERROR, err.Error())
	}
	return this.success()
}

func (this *FriendService) FriendRecommendation(request *imb.FriendRecommendationRequest) *imb.FriendResponse {
	// TODO: check if userId and recommendId are friends
	user := this.database.GetFriendInfo(request.RecommendId)
	buf, err := proto.Marshal(user.(*imb.UserInfo))
	if err != nil {
		return this.failure(CODE_FRIEND_RECOMMENDED_USER_ERROR, "no such user")
	}
	message := &protocol.Message{
		Id:       utils.GenerateMessageId(),
		UserId:   request.UserId,
		TargetId: request.TargetId,
		Action:   int32(ACTION_FRIEND_RECOMMENDED),
		Body:     buf,
	}
	if err := this.notifier.SendP2P(message); err != nil {
		return this.failure(CODE_FRIEND_NOTIFY_RECOMMENDED_ERROR, err.Error())
	}
	return this.success()
}

func (this *FriendService) success() *imb.FriendResponse {
	return &imb.FriendResponse{
		Code: CODE_SUCCESS,
		Data: nil,
	}
}

func (this *FriendService) failure(code int, msg string) *imb.FriendResponse {
	return &imb.FriendResponse{
		Code: int32(code),
		Message: msg,
	}
}
