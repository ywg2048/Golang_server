package models

import (
	"github.com/astaxie/beego/orm"
	"labix.org/v2/mgo/bson"
	cspb "protocol"
)

type Messagecenter struct {
	Id      int32
	Title   string `orm:"size(100)"`
	Content string `orm:"size(256)"`

	IsActive int32
	Time     int64
}
type Userinfo struct {
	Id        int32
	Uid       int64
	userscore []Userscore
}
type Userscore struct {
	Id       int32
	Uid      int64
	Level    int32
	Score    int32
	Startnum int32
	Time     int64
}
type Player struct {
	Saccount           bson.ObjectId          `bson:"_id"`
	Caccount           string                 `bson:"c_account"`
	Uid                int64                  `bson:"uid"`
	WonderfulFriends   WonderfulFriendsData   `bson:"WonderfulFriends"`
	RechargeFlow       RechargeFlowData       `bson:"RechargeFlow"`
	RegistTime         int64                  `bson:"regist_time"`
	LastSignInTime     int64                  `bson:"last_signin_time"`
	FreeSignInOperTime int64                  `bson:"free_signin_oper_time"`
	Levels             []*cspb.CSStageNtf     `bson:"Levels"`
	Money              *cspb.CSMoneyReq       `bson:"Money"`
	FriendList         []*FriendListData      `bson:"FriendList"`
	ApplyFriendList    []*ApplyFriendListData `bson:"ApplyFriendList"`
}
type WonderfulFriendsData struct {
	LastSignInTime     int64 `bson:"last_signin_time"`
	FreeSignInOperTime int64 `bson:"free_signin_oper_time"`
}
type RechargeFlowData struct {
	Account      string `bson:"account"`
	Uid          int64  `bson:"uid"`
	Rmb          int32  `bson:"rmb"`
	GoodsType    int32  `bson:"goods_type"`
	GoodsSubType int32  `bson:"goods_sub_type"`
	GoodsNum     int32  `bson:"goods_num"`
	RechargeTime int64  `bson:"recharge_time"`
	Version      string `bson:"version"`
	Code         string `bson:"code"`
	Channel      string `bson:"channel"`
}
type FriendListData struct {
	Friendid   int32 `bson:"friendid"`
	IsActive   int32 `bson:"isActive"`
	Accepttime int64 `bson:"accepttime"`
}
type ApplyFriendListData struct {
	Applyuid     int32 `bson:"applyuid"`
	IsAccept     int32 `bson:"isAccept"`
	Isrefuse     int32 `bson:"isrefuse"`
	Applytime    int64 `bson:"applytime"`
	Oprationtime int64 `bson:"oprationtime"`
}

func init() {

	orm.RegisterModel(new(Messagecenter), new(Userscore), new(Userinfo))

}
