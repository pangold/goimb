package db

import (
	"database/sql"
	"gitlab.com/pangold/goimb/proto"
)

type Group struct {
	DB *sql.DB
}

func (this *Group) CheckAdminRole(groupId, userId string) bool {
	row := this.DB.QueryRow("SELECT id FROM group_members WHERE user_id = ? AND group_id = ? AND role = ?",
		userId, groupId, imb.Role_ROLE_ADMIN)
	var id int32
	err := row.Scan(&id)
	return err != nil
}

func (this *Group) CheckMasterRole(groupId, userId string) bool {
	row := this.DB.QueryRow("SELECT id FROM group_members WHERE user_id = ? AND group_id = ? AND role = ?",
		userId, groupId, imb.Role_ROLE_MASTER)
	var id int32
	err := row.Scan(&id)
	return err != nil
}

// return imb.Groups
func (this *Group) GetGroups(userId string) interface{} {
	return nil
}

// return imb.GroupInfo
func (this *Group) GetGroupInfo(groupId string) interface{} {
	return nil
}

// return imb.Members
func (this *Group) GetGroupMembers(groupId string) interface{} {
	return nil
}

// return []string
func (this *Group) GetGroupMemberIds(groupId string) []string {
	return nil
}

// return []string
func (this *Group) GetGroupAdminIds(groupId string) []string {
	return nil
}

//
// param GroupInfo
func (this *Group) GroupCreate(interface{}) error {
	return nil
}

func (this *Group) GroupDismiss(groupId string) error {
	return nil
}

func (this *Group) GroupMasterChange(groupId, userId string) error {
	return nil
}

func (this *Group) GroupAdminPromote(groupId, userId string) error {
	return nil
}

func (this *Group) GroupAdminDemote(groupId, userId string) error {
	return nil
}

// save to mongodb & mysql
func (this *Group) GroupMemberJoin(groupId, userId string) error {
	return nil
}

func (this *Group) GroupMemberLeave(groupId, userId string) error {
	return nil
}
