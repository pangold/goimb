package rest

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/pangold/goimb/config"
	"gitlab.com/pangold/goimb/service"
	"log"
)

type Router struct {
	conf config.Host
	engine *gin.Engine
	friendController *FriendController
	groupController *GroupController
}

func NewRouter(conf config.Host, impl *service.Service) *Router {
	router := &Router{
		conf: conf,
		engine: gin.Default(),
		friendController: NewFriendController(impl.FriendService),
		groupController: NewGroupController(impl.GroupService),
	}
	router.friendRouter(router.engine, router.friendController)
	router.groupRouter(router.engine, router.groupController)
	return router
}

func (this *Router) Run() {
	log.Printf("HTTP server start serving on %s", this.conf.Address)
	if err := this.engine.Run(this.conf.Address); err != nil {
		panic(err)
	}
}

func (this *Router) friendRouter(router *gin.Engine, f *FriendController) {
	router.GET   ("/friends/{uid}",       f.GetFriends)
	router.GET   ("/friend/{uid}",        f.GetFriendInfo)
	router.POST  ("/friend/{uid}/apply",  f.FriendMakingApply)
	router.POST  ("/friend/{uid}/accept", f.FriendMakingAccept)
	router.POST  ("/friend/{uid}/reject", f.FriendMakingReject)
	router.DELETE("/friend/{uid}",        f.FriendBreakup)
	router.POST  ("/friend/{uid}/recommend", f.FriendRecommendation)
}

func (this *Router) groupRouter(router *gin.Engine, g *GroupController) {
	router.GET   ("/groups/{uid}",       g.GetGroups)
	router.GET   ("/group/{id}",         g.GetGroup)
	router.GET   ("/group/{id}/members", g.GetGroupMembers)
	router.POST  ("/group",              g.GroupCreate)
	router.DELETE("/group/{id}",         g.GroupDismiss)
	router.POST  ("/group/{id}/change",  g.GroupMasterChange)
	router.POST  ("/group/{id}/promote", g.GroupAdminPromote)
	router.POST  ("/group/{id}/demote",  g.GroupAdminDemote)
	router.POST  ("/group/{id}/apply",   g.GroupJoinApply)
	router.POST  ("/group/{id}/accept",  g.GroupJoinAccept)
	router.POST  ("/group/{id}/reject",  g.GroupJoinReject)
	router.POST  ("/group/{id}/take",    g.GroupMemberTake)
	router.POST  ("/group/{id}/leave",   g.GroupMemberLeave)
	router.POST  ("/group/{id}/kick",    g.GroupMemberKick)

}
