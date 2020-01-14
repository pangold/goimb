package rest

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"gitlab.com/pangold/goimb/proto"
	"gitlab.com/pangold/goimb/service"
	"net/http"
)

type GroupController struct {
	impl *service.GroupService
}

func NewGroupController(impl *service.GroupService) *GroupController {
	return &GroupController{
		impl: impl,
	}
}

func (this *GroupController) GetGroups(ctx *gin.Context) {
	request := &imb.QueryGroupsRequest{}
	if err := ctx.ShouldBindJSON(request); err != nil {
		this.failure(ctx, http.StatusBadRequest, err.Error())
		return
	}
	friend := this.impl.GetGroups(request)
	data, err := json.Marshal(friend)
	if err != nil {
		this.failure(ctx, http.StatusBadRequest, err.Error())
		return
	}
	this.success(ctx, data)
}

func (this *GroupController) GetGroup(ctx *gin.Context) {
	request := &imb.QueryGroupRequest{}
	if err := ctx.ShouldBindJSON(request); err != nil {
		this.failure(ctx, http.StatusBadRequest, err.Error())
		return
	}
	friend := this.impl.GetGroup(request)
	data, err := json.Marshal(friend)
	if err != nil {
		this.failure(ctx, http.StatusBadRequest, err.Error())
		return
	}
	this.success(ctx, data)
}

func (this *GroupController) GetGroupMembers(ctx *gin.Context) {
	request := &imb.QueryGroupRequest{}
	if err := ctx.ShouldBindJSON(request); err != nil {
		this.failure(ctx, http.StatusBadRequest, err.Error())
		return
	}
	friend := this.impl.GetGroupMembers(request)
	data, err := json.Marshal(friend)
	if err != nil {
		this.failure(ctx, http.StatusBadRequest, err.Error())
		return
	}
	this.success(ctx, data)
}

func (this *GroupController) GroupCreate(ctx *gin.Context) {
	request := &imb.GroupInfo{}
	if err := ctx.ShouldBindJSON(request); err != nil {
		this.failure(ctx, http.StatusBadRequest, err.Error())
		return
	}
	resp := this.impl.GroupCreate(request)
	this.response(ctx, resp)
}

func (this *GroupController) GroupDismiss(ctx *gin.Context) {
	request := &imb.GroupDismissRequest{}
	if err := ctx.ShouldBindJSON(request); err != nil {
		this.failure(ctx, http.StatusBadRequest, err.Error())
		return
	}
	resp := this.impl.GroupDismiss(request)
	this.response(ctx, resp)
}

func (this *GroupController) GroupMasterChange(ctx *gin.Context) {
	request := &imb.GroupMasterChangeRequest{}
	if err := ctx.ShouldBindJSON(request); err != nil {
		this.failure(ctx, http.StatusBadRequest, err.Error())
		return
	}
	resp := this.impl.GroupMasterChange(request)
	this.response(ctx, resp)
}

func (this *GroupController) GroupAdminPromote(ctx *gin.Context) {
	request := &imb.GroupAdminPromoteRequest{}
	if err := ctx.ShouldBindJSON(request); err != nil {
		this.failure(ctx, http.StatusBadRequest, err.Error())
		return
	}
	resp := this.impl.GroupAdminPromote(request)
	this.response(ctx, resp)
}

func (this *GroupController) GroupAdminDemote(ctx *gin.Context) {
	request := &imb.GroupAdminDemoteRequest{}
	if err := ctx.ShouldBindJSON(request); err != nil {
		this.failure(ctx, http.StatusBadRequest, err.Error())
		return
	}
	resp := this.impl.GroupAdminDemote(request)
	this.response(ctx, resp)
}

func (this *GroupController) GroupJoinApply(ctx *gin.Context) {
	request := &imb.GroupJoinApplyRequest{}
	if err := ctx.ShouldBindJSON(request); err != nil {
		this.failure(ctx, http.StatusBadRequest, err.Error())
		return
	}
	resp := this.impl.GroupJoinApply(request)
	this.response(ctx, resp)
}

func (this *GroupController) GroupJoinAccept(ctx *gin.Context) {
	request := &imb.GroupJoinHandleRequest{}
	if err := ctx.ShouldBindJSON(request); err != nil {
		this.failure(ctx, http.StatusBadRequest, err.Error())
		return
	}
	resp := this.impl.GroupJoinAccept(request)
	this.response(ctx, resp)
}

func (this *GroupController) GroupJoinReject(ctx *gin.Context ) {
	request := &imb.GroupJoinHandleRequest{}
	if err := ctx.ShouldBindJSON(request); err != nil {
		this.failure(ctx, http.StatusBadRequest, err.Error())
		return
	}
	resp := this.impl.GroupJoinReject(request)
	this.response(ctx, resp)
}

func (this *GroupController) GroupMemberTake(ctx *gin.Context) {
	request := &imb.GroupMemberTakeRequest{}
	if err := ctx.ShouldBindJSON(request); err != nil {
		this.failure(ctx, http.StatusBadRequest, err.Error())
		return
	}
	resp := this.impl.GroupMemberTake(request)
	this.response(ctx, resp)
}

func (this *GroupController) GroupMemberLeave(ctx *gin.Context) {
	request := &imb.GroupMemberLeaveRequest{}
	if err := ctx.ShouldBindJSON(request); err != nil {
		this.failure(ctx, http.StatusBadRequest, err.Error())
		return
	}
	resp := this.impl.GroupMemberLeave(request)
	this.response(ctx, resp)
}

func (this *GroupController) GroupMemberKick(ctx *gin.Context) {
	request := &imb.GroupMemberKickRequest{}
	if err := ctx.ShouldBindJSON(request); err != nil {
		this.failure(ctx, http.StatusBadRequest, err.Error())
		return
	}
	resp := this.impl.GroupMemberKick(request)
	this.response(ctx, resp)
}

func (this *GroupController) success(ctx *gin.Context, data []byte)  {
	resp := &imb.GroupResponse{
		Code: http.StatusOK,
		Message: "",
		Data: data,
	}
	this.response(ctx, resp)
}

func (this *GroupController) failure(ctx *gin.Context, code int, message string) {
	resp := &imb.GroupResponse{
		Code: int32(code),
		Message: message,
	}
	this.response(ctx, resp)
}

func (this *GroupController) response(ctx *gin.Context, resp *imb.GroupResponse) {
	buf, err := json.Marshal(resp)
	code := http.StatusOK
	if err != nil {
		code = http.StatusBadRequest
	}
	ctx.String(code, string(buf))
}

