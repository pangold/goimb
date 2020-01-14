package db

import (
	"database/sql"
	"gitlab.com/pangold/goimb/proto"
	"log"
)

type Friend struct {
	DB *sql.DB
}

func (this *Friend) GetFriends(userId string) interface{} {
	rows, err := this.DB.Query("SELECT u.user_id, u.nickname, u.url " +
		"FROM users AS u " +
		"WHERE u.user_id = (SELECT user1_id FROM friends WHERE user1_id = ?) " +
		"OR " +
		"WHERE u.user_id = (SELECT user2_id FROM friends WHERE user2_id = ?)", userId, userId)
	if err != nil {
		log.Printf(err.Error())
		return nil
	}
	var res *imb.Friends
	for rows.Next() {
		u := &imb.UserInfo{}
		if err := rows.Scan(&u.UserId, &u.NickName, &u.Url); err != nil {
			res.Friends = append(res.Friends, u)
		}
	}
	return res
}

func (this *Friend) GetFriendInfo(uid string) interface{} {
	// FIXME: auth.service
	var res *imb.UserInfo
	row := this.DB.QueryRow("SELECT u.user_id, u.nickname, u.url FROM users AS u WHERE user_id = ?", uid)
	if err := row.Scan(&res.UserId, &res.NickName, &res.Url); err != nil {
		log.Printf(err.Error())
		return res
	}
	return res
}

func (this *Friend) FriendCreate(uid1, uid2 string) error {
	stmt, err := this.DB.Prepare("INSERT friends(user1_id, user2_id) VALUES(?, ?)")
	defer stmt.Close()
	if err != nil {
		return err
	}
	if _, err := stmt.Exec(uid1, uid2); err != nil {
		return err
	}
	return nil
}

func (this *Friend) FriendDelete(uid1, uid2 string) error {
	stmt, err := this.DB.Prepare("DELETE FROM friends WHERE user1_id = ? AND user2_id = ?")
	defer stmt.Close()
	if err != nil {
		return err
	}
	if _, err := stmt.Exec(uid1, uid2); err != nil {
		return err
	}
	return nil
}