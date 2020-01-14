package grpc

import (
	"gitlab.com/pangold/goimb/proto"
	"gitlab.com/pangold/goimb/service"
	"golang.org/x/net/context"
)

type FriendService struct {
	impl *service.FriendService
}

func NewFriendService(impl *service.FriendService) *FriendService {
	return &FriendService{
		impl: impl,
	}
}

func (this *FriendService) GetFriendInfo(ctx context.Context, request *imb.QueryFriendRequest) (*imb.UserInfo, error) {
	return this.impl.GetFriendInfo(request), nil
}

func (this *FriendService) GetFriends(ctx context.Context, request *imb.QueryFriendRequest) (*imb.Friends, error) {
	return this.impl.GetFriends(request), nil
}

func (this *FriendService) FriendMakingApply(ctx context.Context, request *imb.FriendMakingRequest) (*imb.FriendResponse, error) {
	return this.impl.FriendMakingApply(request), nil
}

func (this *FriendService) FriendMakingAccept(ctx context.Context, request *imb.FriendMakingRequest) (*imb.FriendResponse, error) {
	return this.impl.FriendMakeAccept(request), nil
}

func (this *FriendService) FriendMakingReject(ctx context.Context, request *imb.FriendMakingRequest) (*imb.FriendResponse, error) {
	return this.impl.FriendMakeReject(request), nil
}

func (this *FriendService) FriendBreakup(ctx context.Context, request *imb.FriendBreakupRequest) (*imb.FriendResponse, error) {
	return this.impl.FriendBreakup(request), nil
}

func (this *FriendService) FriendRecommendation(ctx context.Context, request *imb.FriendRecommendationRequest) (*imb.FriendResponse, error) {
	return this.impl.FriendRecommendation(request), nil
}
