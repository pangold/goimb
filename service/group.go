package service

import (
	"github.com/golang/protobuf/proto"
	"gitlab.com/pangold/goim/protocol"
	"gitlab.com/pangold/goimb/model"
	"gitlab.com/pangold/goimb/proto"
	"log"
)

type GroupService struct {
	database model.GroupDB
	notifier *NotificationService
}

func NewGroupService(n *NotificationService, d model.GroupDB) *GroupService {
	return &GroupService{
		database: d,
		notifier: n,
	}
}

func (this *GroupService) SendGroup(message *protocol.Message) error {
	ids := this.database.GetGroupMemberIds(message.GroupId)
	for _, id := range ids {
		// FIXME: clone message to msg
		msg := message
		msg.TargetId = id
		this.notifier.SendP2P(msg)
	}
	return nil
}

func (this *GroupService) SendGroupForAdmin(message *protocol.Message) error {
	ids := this.database.GetGroupAdminIds(message.GroupId)
	for _, id := range ids {
		// FIXME: clone message to msg
		msg := message
		msg.TargetId = id
		this.notifier.SendP2P(msg)
	}
	return nil
}

func (this *GroupService) GetGroups(request *imb.QueryGroupsRequest) *imb.Groups {
	return this.database.GetGroups(request.UserId).(*imb.Groups)
}

func (this *GroupService) GetGroup(request *imb.QueryGroupRequest) *imb.GroupInfo {
	return this.database.GetGroupInfo(request.GroupId).(*imb.GroupInfo)
}

func (this *GroupService) GetGroupMembers(request *imb.QueryGroupRequest) *imb.GroupMembers {
	return this.database.GetGroupMembers(request.GroupId).(*imb.GroupMembers)
}

func (this *GroupService) GroupCreate(request *imb.GroupInfo) *imb.GroupResponse {
	if len(request.Members) == 0 {
		return this.failure(CODE_GROUP_CREATE_ERROR, "invalid param")
	}
	if err := this.database.GroupCreate(request); err != nil {
		return this.failure(CODE_GROUP_CREATE_ERROR, err.Error())
	}
	msg := &protocol.Message{
		UserId: request.Members[0].User.UserId,       // who creates group
		GroupId: request.GroupId,
		Action: int32(ACTION_GROUP_CREATED),
		Body: this.marshal(request), // group info
	}
	if err := this.SendGroup(msg); err != nil {
		return this.failure(CODE_GROUP_NOTIFY_CREATED_ERROR, err.Error())
	}
	return this.success()
}

func (this *GroupService) GroupDismiss(request *imb.GroupDismissRequest) *imb.GroupResponse {
	// check userId's permission, master role is require
	if !this.database.CheckMasterRole(request.GroupId, request.UserId) {
		return this.failure(CODE_GROUP_DISMISS_ERROR, "unauthorized")
	}
	if err := this.database.GroupDismiss(request.GroupId); err != nil {
		return this.failure(CODE_GROUP_DISMISS_ERROR, err.Error())
	}
	message := &protocol.Message{
		UserId: request.UserId,   // Who dismisses group?
		GroupId: request.GroupId,
		Action: int32(ACTION_GROUP_DISMISSED),
	}
	if err := this.SendGroup(message); err != nil {
		return this.failure(CODE_GROUP_NOTIFY_DISMISSED_ERROR, err.Error())
	}
	return this.success()
}

func (this *GroupService) GroupMasterChange(request *imb.GroupMasterChangeRequest) *imb.GroupResponse {
	// check userId's permission, master role is require
	if !this.database.CheckMasterRole(request.GroupId, request.UserId) {
		return this.failure(CODE_GROUP_MASTER_CHANGE_ERROR, "unauthorized")
	}
	if err := this.database.GroupMasterChange(request.GroupId, request.TargetId); err != nil {
		return this.failure(CODE_GROUP_MASTER_CHANGE_ERROR, err.Error())
	}
	message := &protocol.Message{
		UserId: request.UserId,       // operator(old master)
		TargetId: request.TargetId,   // NEW MASTER
		GroupId: request.GroupId,
		Action: int32(ACTION_GROUP_MASTER_CHANGED),
	}
	if err := this.SendGroup(message); err != nil {
		return this.failure(CODE_GROUP_NOTIFY_MASTER_CHANGED_ERROR, err.Error())
	}
	return this.success()
}

func (this *GroupService) GroupAdminPromote(request *imb.GroupAdminPromoteRequest) *imb.GroupResponse {
	// check userId's permission, admin role is require
	if !this.database.CheckAdminRole(request.GroupId, request.UserId) {
		return this.failure(CODE_GROUP_ADMIN_PROMOTE_ERROR, "unauthorized")
	}
	if err := this.database.GroupAdminPromote(request.GroupId, request.TargetId); err != nil {
		return this.failure(CODE_GROUP_ADMIN_PROMOTE_ERROR, err.Error())
	}
	message := &protocol.Message{
		UserId: request.UserId,         // operator
		TargetId: request.TargetId,     // new promoted admin
		GroupId: request.GroupId,
		Action: int32(ACTION_GROUP_ADMIN_PROMOTED),
	}
	if err := this.SendGroup(message); err != nil {
		return this.failure(CODE_GROUP_NOTIFY_ADMIN_PROMOTED_ERROR, err.Error())
	}
	return this.success()
}

func (this *GroupService) GroupAdminDemote(request *imb.GroupAdminDemoteRequest) *imb.GroupResponse {
	// check userId's permission, master role is require
	if !this.database.CheckMasterRole(request.GroupId, request.UserId) {
		return this.failure(CODE_GROUP_ADMIN_DEMOTE_ERROR, "unauthorized")
	}
	if err := this.database.GroupAdminDemote(request.GroupId, request.TargetId); err != nil {
		return this.failure(CODE_GROUP_ADMIN_DEMOTE_ERROR, err.Error())
	}
	message := &protocol.Message{
		UserId: request.UserId,          // operator
		TargetId: request.TargetId,      // admin who is being demoted
		GroupId: request.GroupId,
		Action: int32(ACTION_GROUP_ADMIN_DEMOTED),
	}
	if err := this.SendGroup(message); err != nil {
		return this.failure(CODE_GROUP_NOTIFY_ADMIN_DEMOTED_ERROR, err.Error())
	}
	return this.success()
}

func (this *GroupService) GroupJoinApply(request *imb.GroupJoinApplyRequest) *imb.GroupResponse {
	message := &protocol.Message{
		UserId: request.UserId,        // user who applies to join
		GroupId: request.GroupId,
		Action: int32(ACTION_GROUP_JOIN_APPLIED),
		Body: []byte(request.Postscript),
	}
	// just notify admins to accept/reject
	if err := this.SendGroupForAdmin(message); err != nil {
		return this.failure(CODE_GROUP_NOTIFY_JOIN_APPLIED_ERROR, err.Error())
	}
	return this.success()
}

func (this *GroupService) GroupJoinAccept(request *imb.GroupJoinHandleRequest) *imb.GroupResponse {
	// check userId's permission, admin role is require
	if !this.database.CheckAdminRole(request.GroupId, request.UserId) {
		return this.failure(CODE_GROUP_JOIN_ACCEPT_ERROR, "unauthorized")
	}
	// TODO: check if is handled? no matter accept or reject?
	if err := this.database.GroupMemberJoin(request.GroupId, request.TargetId); err != nil {
		return this.failure(CODE_GROUP_JOIN_ACCEPT_ERROR, err.Error())
	}
	message := &protocol.Message{
		UserId: request.UserId,      // operator
		TargetId: request.TargetId,  // user who is accepted to join
		GroupId: request.GroupId,
		Action: int32(ACTION_GROUP_JOIN_ACCEPTED),
		Body: []byte(request.Postscript),
	}
	if err := this.SendGroup(message); err != nil {
		return this.failure(CODE_GROUP_NOTIFY_JOIN_ACCEPTED_ERROR, err.Error())
	}
	return this.success()
}

func (this *GroupService) GroupJoinReject(request *imb.GroupJoinHandleRequest) *imb.GroupResponse {
	// check userId's permission, admin role is require
	if !this.database.CheckAdminRole(request.GroupId, request.UserId) {
		return this.failure(CODE_GROUP_JOIN_REJECT_ERROR, "unauthorized")
	}
	// TODO: check if is handled? no matter accept or reject?
	message := &protocol.Message{
		UserId: request.UserId,      // operator
		TargetId: request.TargetId,  // being rejected
		GroupId: request.GroupId,
		Action: int32(ACTION_GROUP_JOIN_REJECTED),
		Body: []byte(request.Postscript),
	}
	if err := this.SendGroupForAdmin(message); err != nil {
		return this.failure(CODE_GROUP_NOTIFY_JOIN_REJECTED_ERROR, err.Error())
	}
	if err := this.notifier.SendP2P(message); err != nil {
		return this.failure(CODE_GROUP_NOTIFY_JOIN_REJECTED_ERROR, err.Error())
	}
	return this.success()
}

func (this *GroupService) GroupMemberTake(request *imb.GroupMemberTakeRequest) *imb.GroupResponse {
	if err := this.database.GroupMemberJoin(request.GroupId, request.TargetId); err != nil {
		return this.failure(CODE_GROUP_MEMBER_TAKE_ERROR, err.Error())
	}
	message := &protocol.Message{
		UserId: request.UserId,      // operator
		TargetId: request.TargetId,  // being token
		GroupId: request.GroupId,
		Action: int32(ACTION_GROUP_MEMBER_TOKEN),
	}
	if err := this.SendGroup(message); err != nil {
		return this.failure(CODE_GROUP_NOTIFY_MEMBER_TOKEN_ERROR, err.Error())
	}
	return this.success()
}

func (this *GroupService) GroupMemberLeave(request *imb.GroupMemberLeaveRequest) *imb.GroupResponse {
	if err := this.database.GroupMemberLeave(request.GroupId, request.UserId); err != nil {
		return this.failure(CODE_GROUP_MEMBER_LEAVE_ERROR, err.Error())
	}
	message := &protocol.Message{
		UserId: request.UserId,     // the member who left
		GroupId: request.GroupId,
		Action: int32(ACTION_GROUP_MEMBER_LEFT),
	}
	// notify to all members that someone left
	if err := this.SendGroup(message); err != nil {
		return this.failure(CODE_GROUP_NOTIFY_MEMBER_LEFT_ERROR, err.Error())
	}
	return this.success()
}

func (this *GroupService) GroupMemberKick(request *imb.GroupMemberKickRequest) *imb.GroupResponse {
	// check userId's permission, admin role is require
	if !this.database.CheckAdminRole(request.GroupId, request.UserId) {
		return this.failure(CODE_GROUP_MEMBER_KICK_ERROR, "unauthorized")
	}
	if err := this.database.GroupMemberLeave(request.GroupId, request.TargetId); err != nil {
		return this.failure(CODE_GROUP_MEMBER_KICK_ERROR, err.Error())
	}
	message := &protocol.Message{
		UserId: request.UserId,     // operator
		TargetId: request.TargetId, // being kicked
		GroupId: request.GroupId,
		Action: int32(ACTION_GROUP_MEMBER_KICKED),
	}
	// notify to all members that someone's been kicked out of the group
	if err := this.SendGroup(message); err != nil {
		return this.failure(CODE_GROUP_NOTIFY_MEMBER_KICKED_ERROR, err.Error())
	}
	return this.success()
}

func (this *GroupService) marshal(message proto.Message) []byte {
	buf, err := proto.Marshal(message)
	if err != nil {
		log.Printf(err.Error())
		return nil
	}
	return buf
}

func (this *GroupService) success() *imb.GroupResponse {
	return &imb.GroupResponse{
		Code: 200,
		Message: "",
		Data: nil,
	}
}

func (this *GroupService) failure(code int, msg string) *imb.GroupResponse {
	return &imb.GroupResponse{
		Code: int32(code),
		Message: msg,
		Data: nil,
	}
}