package model

import (
	"gitlab.com/pangold/goimb/proto"
	"log"
	"testing"
)

var (
	database *Database
)

func init() {
	database = NewSimple()
}

func TestFriend(t *testing.T) {
	friends := database.Friend.GetFriends("10000").(*imb.Friends)
	for _, f := range friends.Friends {
		log.Printf("%s %s %s\n", f.UserId, f.NickName, f.Url)
	}
	friend := database.Friend.GetFriendInfo("10001")
	if friend == nil {
		t.Error("unexpected result")
	}
	log.Printf("%v\n", friend)
	if err := database.Friend.FriendDelete("10002", "10003"); err == nil {
		t.Error("unexpected result")
	}
	if err := database.Friend.FriendCreate("10002", "10003"); err != nil {
		t.Errorf(err.Error())
	}
	if err := database.Friend.FriendDelete("10002", "10003"); err != nil {
		t.Error(err.Error())
	}
}

func TestGroup(t *testing.T) {
	groups := database.Group.GetGroups("10000")
	if groups == nil {
		t.Error("unexpected result")
	}
	log.Println(groups)
	group := database.Group.GetGroupInfo(groups.(*imb.Groups).Groups[0].GroupId)
	if group == nil {
		t.Error("unexpected result")
	}
	log.Println(group)
	//
	members := database.Group.GetGroupMembers(group.(*imb.GroupInfo).GroupId)
	if members == nil {
		t.Error("unexpected result")
	}
	log.Println(members)
	//
	memberIds := database.Group.GetGroupMemberIds(group.(*imb.GroupInfo).GroupId)
	if memberIds == nil {
		t.Error("unexpected result")
	}
	log.Println(memberIds)
	//
	memberAdmins := database.Group.GetGroupAdminIds(group.(*imb.GroupInfo).GroupId)
	if memberIds == nil {
		t.Error("unexpected result")
	}
	log.Println(memberAdmins)
	//
	group2 := &imb.GroupInfo{GroupName: "Test", Script: "Test", Members: []*imb.GroupMember{
		{User: &imb.UserInfo{UserId: "10000"}},
		{User: &imb.UserInfo{UserId: "10001"}},
		{User: &imb.UserInfo{UserId: "10002"}},
		{User: &imb.UserInfo{UserId: "10003"}},
	}}
	if err := database.Group.GroupCreate(group2); err != nil {
		t.Error(err.Error())
	}
	log.Println(group2)
	//
	//if err := database.Group.GroupDismiss(group2.GroupId); err != nil {
	//	t.Error(err.Error())
	//}
	//
	if err := database.Group.GroupMasterChange(group2.GroupId, "10001"); err != nil {
		t.Error(err.Error())
	}
	log.Println(group2)
	//
	if err := database.Group.GroupAdminPromote(group2.GroupId, "10002"); err != nil {
		t.Error(err.Error())
	}
	log.Println(group2)
	//
	if err := database.Group.GroupAdminDemote(group2.GroupId, "10002"); err != nil {
		t.Error(err.Error())
	}
	log.Println(group2)
	//
	if err := database.Group.GroupMemberJoin(group2.GroupId, "10004"); err != nil {
		t.Error(err.Error())
	}
	log.Println(group2)
	//
	if err := database.Group.GroupMemberLeave(group2.GroupId, "10004"); err != nil {
		t.Error(err.Error())
	}
	log.Println(group2)
}