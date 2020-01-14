package mongodb

import (
	"gitlab.com/pangold/goim/protocol"
	"gitlab.com/pangold/goimb/config"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"time"
)

type MongoDB struct {
	conn *mgo.Session
	c *mgo.Collection
}

func NewMongoDB(conf config.Host) *MongoDB {
	conn, err := mgo.Dial(conf.Address)
	if err != nil {
		panic(err)
	}
	conn.SetMode(mgo.Monotonic, true)
	return &MongoDB {
		conn: conn,
		c: conn.DB("history").C("message"),
	}
}

// cid could be cids
func (this *MongoDB) Add(message *protocol.Message, cids []string) {
	record := &Record{}
	if err := this.c.Find(bson.M{"id": message.Id}).One(record); err != nil {
		record.UpdatedTime = time.Now()
		for _, cid := range cids {
			this.setClient(record, cid)
		}
		this.c.Update(record.Id, record)
	} else {
		record.Message = *message
		record.CreatedTime = time.Now()
		record.UpdatedTime = time.Now()
		for _, cid := range cids {
			this.setClient(record, cid)
		}
		this.c.Insert(record)
	}
}

func (this *MongoDB) Find(uid, cid string) (res []*protocol.Message) {
	var records []Record
	if err := this.c.Find(this.getBson(uid, cid)).All(records); err != nil {
		log.Printf("find history(uid %s, cid %s) error: %v", uid, cid, err)
		return nil
	}
	for i := 0; i < len(records); i++ {
		res = append(res, &records[i].Message)
		this.setClient(&records[i], cid)
	}
	// FIXME:
	if err := this.c.Update(nil, records); err != nil {
		return nil
	}
	return res
}

func (this *MongoDB) Hit(uid, tid, gid string, action int32) bool {
	// some of uid/tid/gid is empty
	var condition = make(map[string]interface{})
	condition["action"] = action
	if uid != ""  {
		condition["user_id"] = tid
	}
	if tid != "" {
		condition["target_id"] = tid
	}
	if gid != "" {
		condition["group_id"] = gid
	}
	if err := this.c.Find(condition); err != nil {
		return true
	}
	return false
}

func (this *MongoDB) setClient(record *Record, cid string) {
	if cid == "pc" {
		record.PC = true
	} else if cid == "android" {
		record.Android = true
	} else if cid == "ios" {
		record.IOS = true
	} else if cid == "web" {
		record.Web = true
	}
}

func (this *MongoDB) getBson(uid, cid string) interface{} {
	if cid == "pc" {
		return bson.M{"user_id": uid, "pc": false}
	} else if cid == "android" {
		return bson.M{"user_id": uid, "android": false}
	} else if cid == "ios" {
		return bson.M{"user_id": uid, "ios": false}
	} else if cid == "web" {
		return bson.M{"user_id": uid, "web": false}
	}
	return nil
}
