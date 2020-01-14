package grpc

import (
	"gitlab.com/pangold/goimb/proto"
	"gitlab.com/pangold/goimb/service"
	"golang.org/x/net/context"
)

type GroupService struct {
	impl *service.GroupService
}

func NewGroupService(impl *service.GroupService) *GroupService {
	return &GroupService{
		impl: impl,
	}
}

func (this *GroupService) GetGroups(ctx context.Context, request *imb.QueryGroupsRequest) (*imb.Groups, error) {
	return this.impl.GetGroups(request), nil
}

func (this *GroupService) GetGroup(ctx context.Context, request *imb.QueryGroupRequest) (*imb.GroupInfo, error) {
	return this.impl.GetGroup(request), nil
}

func (this *GroupService) GetGroupMembers(ctx context.Context, request *imb.QueryGroupRequest) (*imb.GroupMembers, error) {
	return this.impl.GetGroupMembers(request), nil
}


func (this *GroupService) GroupCreate(ctx context.Context, request *imb.GroupInfo) (*imb.GroupResponse, error) {
	return this.impl.GroupCreate(request), nil
}

func (this *GroupService) GroupDismiss(ctx context.Context, request *imb.GroupDismissRequest) (*imb.GroupResponse, error) {
	return this.impl.GroupDismiss(request), nil
}

func (this *GroupService) GroupMasterChange(ctx context.Context, request *imb.GroupMasterChangeRequest) (*imb.GroupResponse, error) {
	return this.impl.GroupMasterChange(request), nil
}

func (this *GroupService) GroupAdminPromote(ctx context.Context, request *imb.GroupAdminPromoteRequest) (*imb.GroupResponse, error) {
	return this.impl.GroupAdminPromote(request), nil
}

func (this *GroupService) GroupAdminDemote(ctx context.Context, request *imb.GroupAdminDemoteRequest) (*imb.GroupResponse, error) {
	return this.impl.GroupAdminDemote(request), nil
}

func (this *GroupService) GroupJoinApply(ctx context.Context, request *imb.GroupJoinApplyRequest) (*imb.GroupResponse, error) {
	return this.impl.GroupJoinApply(request), nil
}

func (this *GroupService) GroupJoinAccept(ctx context.Context, request *imb.GroupJoinHandleRequest) (*imb.GroupResponse, error) {
	return this.impl.GroupJoinAccept(request), nil
}

func (this *GroupService) GroupJoinReject(rctx context.Context, request *imb.GroupJoinHandleRequest) (*imb.GroupResponse, error) {
	return this.impl.GroupJoinReject(request), nil
}

func (this *GroupService) GroupMemberTake(ctx context.Context, request *imb.GroupMemberTakeRequest) (*imb.GroupResponse, error) {
	return this.impl.GroupMemberTake(request), nil
}

func (this *GroupService) GroupMemberLeave(ctx context.Context, request *imb.GroupMemberLeaveRequest) (*imb.GroupResponse, error) {
	return this.impl.GroupMemberLeave(request), nil
}

func (this *GroupService) GroupMemberKick(ctx context.Context, request *imb.GroupMemberKickRequest) (*imb.GroupResponse, error) {
	return this.impl.GroupMemberKick(request), nil
}
