package model

import (
	"gitlab.com/pangold/goimb/config"
	"gitlab.com/pangold/goimb/model/db"
	"gitlab.com/pangold/goimb/model/simple"
)

type FriendDB interface {
	//
	GetFriends(uid string) interface{}               // return imb.Friends
	GetFriendInfo(uid string) interface{}            // return imb.FriendInfo
	//
	FriendCreate(userId1, userId2 string) error      //
	FriendDelete(userId1, userId2 string) error      //
}

type GroupDB interface {
	CheckAdminRole(groupId, userId string) bool
	CheckMasterRole(groupId, userId string) bool
	//
	GetGroups(userId string) interface{}             // return imb.Groups
	GetGroupInfo(groupId string) interface{}         // return imb.GroupInfo
	GetGroupMembers(groupId string) interface{}      // return imb.Members
	GetGroupMemberIds(groupId string) []string       // return []string
	GetGroupAdminIds(groupId string) []string        // return []string
	//
	GroupCreate(interface{}) error                   // param GroupInfo
	GroupDismiss(groupId string) error
	GroupMasterChange(groupId, userId string) error
	GroupAdminPromote(groupId, userId string) error
	GroupAdminDemote(groupId, userId string) error
	GroupMemberJoin(groupId, userId string) error
	GroupMemberLeave(groupId, userId string) error
}

type Database struct {
	Friend FriendDB
	Group GroupDB
}

func NewDatabase(c config.MySQL) *Database {
	if c.Disabled {
		return NewSimple()
	}
	return NewDB(c)
}

func NewDB(c config.MySQL) *Database {
	conn := db.NewDB(c.User, c.Password, c.Host, c.DBName, c.Port)
	mysql := &Database {
		Friend: &db.Friend{DB: conn},
		Group: &db.Group{DB: conn},
	}
	return mysql
}

func NewSimple() *Database {
	return &Database{
		Friend: &simple.Friend{},
		Group: &simple.Group{},
	}
}