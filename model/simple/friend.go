package simple

import (
	"fmt"
	"gitlab.com/pangold/goimb/proto"
)

var (
	users []*imb.UserInfo
	relationships [][2]*imb.UserInfo
)

func init() {
	// registered users
	users = []*imb.UserInfo {
		{UserId: "10000", Url: "0.png", NickName: "panjin"},
		{UserId: "10001", Url: "0.png", NickName: "pandora"},
		{UserId: "10002", Url: "0.png", NickName: "pan"},
		{UserId: "10003", Url: "0.png", NickName: "dora"},
		{UserId: "10004", Url: "0.png", NickName: "baba"},
		{UserId: "10005", Url: "0.png", NickName: "mama"},
		{UserId: "10006", Url: "0.png", NickName: "gege"},
		{UserId: "10007", Url: "0.png", NickName: "jiejie"},
		{UserId: "10008", Url: "0.png", NickName: "didi"},
		{UserId: "10009", Url: "0.png", NickName: "meimei"},
	}
	relationships = [][2]*imb.UserInfo {
		{users[0], users[1]},
		{users[0], users[2]},
		{users[0], users[3]},
		{users[1], users[2]},
	}
}

func uget(userId string) *imb.UserInfo {
	for i := 0; i < len(users); i++ {
		if users[i].UserId == userId {
			return users[i]
		}
	}
	return nil
}

type Friend struct {

}

func (this *Friend) GetFriends(userId string) interface{} {
	friends := &imb.Friends{}
	for i := 0; i < len(relationships); i++ {
		if relationships[i][0].UserId == userId {
			friends.Friends = append(friends.Friends, relationships[i][1])
		} else if relationships[i][1].UserId == userId {
			friends.Friends = append(friends.Friends, relationships[i][0])
		}
	}
	return friends
}

func (this *Friend) GetFriendInfo(uid string) interface{} {
	return uget(uid)
}

func (this *Friend) FriendCreate(uid1, uid2 string) error {
	user1, user2 := uget(uid1), uget(uid2)
	if user1 == nil {
		return fmt.Errorf("uid(%s) is not exist ", uid1)
	}
	if user2 == nil {
		return fmt.Errorf("uid(%s) is not exist ", uid2)
	}
	relationships = append(relationships, [2]*imb.UserInfo{user1, user2})
	return nil
}

func (this *Friend) FriendDelete(uid1, uid2 string) error {
	for i := 0; i < len(relationships); i++ {
		if (relationships[i][0].UserId == uid1 && relationships[i][1].UserId == uid2) ||
			(relationships[i][0].UserId == uid2 && relationships[i][1].UserId == uid1) {
			relationships = append(relationships[:i], relationships[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("uid(%s) and uid(%s) are not friend yet", uid1, uid2)
}