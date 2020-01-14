package rest

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"gitlab.com/pangold/goimb/proto"
	"gitlab.com/pangold/goimb/service"
	"net/http"
)

type FriendController struct {
	impl *service.FriendService
}

func NewFriendController(impl *service.FriendService) *FriendController {
	return &FriendController{
		impl: impl,
	}
}

func (this *FriendController) GetFriendInfo(ctx *gin.Context) {
	request := &imb.QueryFriendRequest{}
	if err := ctx.ShouldBindJSON(request); err != nil {
		this.failure(ctx, http.StatusBadRequest, err.Error())
		return
	}
	friend := this.impl.GetFriendInfo(request)
	data, err := json.Marshal(friend)
	if err != nil {
		this.failure(ctx, http.StatusBadRequest, err.Error())
		return
	}
	this.success(ctx, data)
}

func (this *FriendController) GetFriends(ctx *gin.Context) {
	request := &imb.QueryFriendRequest{}
	if err := ctx.ShouldBindJSON(request); err != nil {
		this.failure(ctx, http.StatusBadRequest, err.Error())
		return
	}
	friends := this.impl.GetFriends(request)
	data, err := json.Marshal(friends)
	if err != nil {
		this.failure(ctx, http.StatusBadRequest, err.Error())
		return
	}
	this.success(ctx, data)
}

func (this *FriendController) FriendMakingApply(ctx *gin.Context) {
	request := &imb.FriendMakingRequest{}
	if err := ctx.ShouldBindJSON(request); err != nil {
		this.failure(ctx, http.StatusBadRequest, err.Error())
		return
	}
	resp := this.impl.FriendMakingApply(request)
	this.response(ctx, resp)
}

func (this *FriendController) FriendMakingAccept(ctx *gin.Context) {
	request := &imb.FriendMakingRequest{}
	if err := ctx.ShouldBindJSON(request); err != nil {
		this.failure(ctx, http.StatusBadRequest, err.Error())
		return
	}
	resp := this.impl.FriendMakeAccept(request)
	this.response(ctx, resp)
}

func (this *FriendController) FriendMakingReject(ctx *gin.Context) {
	request := &imb.FriendMakingRequest{}
	if err := ctx.ShouldBindJSON(request); err != nil {
		this.failure(ctx, http.StatusBadRequest, err.Error())
		return
	}
	resp := this.impl.FriendMakeReject(request)
	this.response(ctx, resp)
}

func (this *FriendController) FriendBreakup(ctx *gin.Context) {
	request := &imb.FriendBreakupRequest{}
	if err := ctx.ShouldBindJSON(request); err != nil {
		this.failure(ctx, http.StatusBadRequest, err.Error())
		return
	}
	resp := this.impl.FriendBreakup(request)
	this.response(ctx, resp)
}

func (this *FriendController) FriendRecommendation(ctx *gin.Context) {
	request := &imb.FriendRecommendationRequest{}
	if err := ctx.ShouldBindJSON(request); err != nil {
		this.failure(ctx, http.StatusBadRequest, err.Error())
		return
	}
	resp := this.impl.FriendRecommendation(request)
	this.response(ctx, resp)
}

func (this *FriendController) success(ctx *gin.Context, data []byte)  {
	resp := &imb.FriendResponse{
		Code: http.StatusOK,
		Message: "",
		Data: data,
	}
	this.response(ctx, resp)
}

func (this *FriendController) failure(ctx *gin.Context, code int, message string) {
	resp := &imb.FriendResponse{
		Code: int32(code),
		Message: message,
	}
	this.response(ctx, resp)
}

func (this *FriendController) response(ctx *gin.Context, resp *imb.FriendResponse) {
	buf, err := json.Marshal(resp)
	code := http.StatusOK
	if err != nil {
		code = http.StatusBadRequest
	}
	ctx.String(code, string(buf))
}