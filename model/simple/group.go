package simple

import (
	"errors"
	"fmt"
	"gitlab.com/pangold/goimb/proto"
	"strconv"
	"sync/atomic"
)

var (
	groups []*imb.GroupInfo
	lastId int32
)

func init() {
	group := &imb.GroupInfo{GroupId: groupId(), GroupName: "Family", Script: "We are family"}
	group.Members = append(group.Members, &imb.GroupMember{GroupId: group.GroupId, Role: imb.Role_ROLE_MASTER,   User: users[0]})
	group.Members = append(group.Members, &imb.GroupMember{GroupId: group.GroupId, Role: imb.Role_ROLE_ORDINARY, User: users[1]})
	group.Members = append(group.Members, &imb.GroupMember{GroupId: group.GroupId, Role: imb.Role_ROLE_ORDINARY, User: users[2]})
	group.Members = append(group.Members, &imb.GroupMember{GroupId: group.GroupId, Role: imb.Role_ROLE_ORDINARY, User: users[3]})
	group.Members = append(group.Members, &imb.GroupMember{GroupId: group.GroupId, Role: imb.Role_ROLE_ORDINARY, User: users[4]})
	groups = append(groups, group)
}

func groupId() string {
	id := atomic.AddInt32(&lastId, 1)
	return strconv.Itoa(int(id))
}

// get pos of where groupId located
func gpos(groupId string) int {
	for i := 0; i < len(groups); i++ {
		if groups[i].GroupId == groupId {
			return i
		}
	}
	return -1
}

// get pos of where memberId located at (groups.Members)
func mpos(groupId, userId string) (int, int) {
	if pos := gpos(groupId); pos != -1 {
		for i := 0; i < len(groups[pos].Members); i++ {
			if groups[pos].Members[i].User.UserId == userId {
				return pos, i
			}
		}
		return pos, -1
	}
	return -1, -1
}

func gget(groupId string) *imb.GroupInfo {
	gpos := gpos(groupId)
	if gpos != -1 {
		return groups[gpos]
	}
	return nil
}

func mget(groupId, userId string) *imb.GroupMember {
	gpos, mpos := mpos(groupId, userId)
	if gpos != -1 && mpos != -1 {
		return groups[gpos].Members[mpos]
	}
	return nil
}

//
func ugets(userId string) (res []*imb.GroupInfo) {
	for i := 0; i < len(groups); i++ {
		for j := 0; j < len(groups[i].Members); j++ {
			if groups[i].Members[j].User.UserId == userId {
				res = append(res, groups[i])
			}
		}
	}
	return res
}


//////
type Group struct {

}

func (this *Group) CheckAdminRole(groupId, userId string) bool {
	if member := mget(groupId, userId); member != nil {
		return member.Role > imb.Role_ROLE_ORDINARY
	}
	return false
}

func (this *Group) CheckMasterRole(groupId, userId string) bool {
	if member := mget(groupId, userId); member != nil {
		return member.Role > imb.Role_ROLE_MASTER
	}
	return false
}

// return imb.Groups
func (this *Group) GetGroups(userId string) interface{} {
	res := &imb.Groups{}
	res.Groups = ugets(userId)
	return res
}

// return imb.GroupInfo
func (this *Group) GetGroupInfo(groupId string) interface{} {
	return gget(groupId)
}

// return imb.Members
func (this *Group) GetGroupMembers(groupId string) interface{} {
	if group := gget(groupId); group != nil {
		return group.Members
	}
	return nil
}

// return []string
func (this *Group) GetGroupMemberIds(groupId string) (res []string) {
	if group := gget(groupId); group != nil {
		for _, m := range group.Members {
			res = append(res, m.User.UserId)
		}
	}
	return res
}

// return []string
func (this *Group) GetGroupAdminIds(groupId string) (res []string) {
	if group := gget(groupId); group != nil {
		for _, m := range group.Members {
			if m.Role > imb.Role_ROLE_ORDINARY {
				res = append(res, m.User.UserId)
			}
		}
	}
	return res
}

//
// param GroupInfo
func (this *Group) GroupCreate(group interface{}) error {
	g := group.(*imb.GroupInfo)
	// fill users(members)
	for i := 0; i < len(g.Members); i++ {
		g.Members[i].User = uget(g.Members[i].User.UserId)
		if g.Members[i].User == nil {
			g.Members = append(g.Members[:i], g.Members[i+1:]...)
		}
	}
	// a group should has member(s), at least mater
	if len(g.Members) < 2 {
		return errors.New("requires at least 2 members")
	}
	// generate group Id
	g.GroupId = groupId()
	// first member is master
	g.Members[0].Role = imb.Role_ROLE_MASTER
	groups = append(groups, g)
	return nil
}

func (this *Group) GroupDismiss(groupId string) error {
	if pos := gpos(groupId); pos != -1 {
		groups = append(groups[:pos], groups[pos+1:]...)
		return nil
	}
	return fmt.Errorf("group %s is not exist", groupId)
}

func (this *Group) GroupMasterChange(groupId, userId string) error {
	var pre, now *imb.GroupMember
	pos := gpos(groupId)
	if pos == -1 {
		return fmt.Errorf("group %s is not exist", groupId)
	}
	for i := 0; i < len(groups[pos].Members); i++ {
		if groups[pos].Members[i].Role == imb.Role_ROLE_MASTER {
			pre = groups[pos].Members[i]
		}
		if groups[pos].Members[i].User.UserId == userId {
			now = groups[pos].Members[i]
		}
	}
	if pre != nil {
		pre.Role = imb.Role_ROLE_ORDINARY
	}
	if now == nil {
		return fmt.Errorf("no such user %s", userId)
	}
	now.Role = imb.Role_ROLE_MASTER
	return nil
}

func (this *Group) GroupAdminPromote(groupId, userId string) error {
	gpos, mpos := mpos(groupId, userId)
	if gpos == -1 || mpos == -1 {
		return fmt.Errorf("no such group %s or member %s", groupId, userId)
	}
	groups[gpos].Members[mpos].Role = imb.Role_ROLE_ADMIN
	return nil
}

func (this *Group) GroupAdminDemote(groupId, userId string) error {
	gpos, mpos := mpos(groupId, userId)
	if gpos == -1 || mpos == -1 {
		return fmt.Errorf("no such group %s or member %s", groupId, userId)
	}
	groups[gpos].Members[mpos].Role = imb.Role_ROLE_ORDINARY
	return nil
}

func (this *Group) GroupMemberJoin(groupId, userId string) error {
	group := gget(groupId)
	if group == nil {
		return fmt.Errorf("no such group %s", groupId)
	}
	user := uget(userId)
	if user == nil {
		return fmt.Errorf("no such user %s", userId)
	}
	member := &imb.GroupMember{GroupId: groupId, User: user, Role: imb.Role_ROLE_ORDINARY}
	group.Members = append(group.Members, member)
	return nil
}

func (this *Group) GroupMemberLeave(groupId, userId string) error {
	gp, mp := mpos(groupId, userId)
	if gp == -1 {
		return fmt.Errorf("no such group %s", groupId)
	}
	if mp == -1 {
		return fmt.Errorf("no such member %s", userId)
	}
	groups[gp].Members = append(groups[gp].Members[:mp], groups[gp].Members[mp+1:]...)
	return nil
}
