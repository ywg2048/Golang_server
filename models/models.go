package models

import (
	"github.com/astaxie/beego/orm"
	"labix.org/v2/mgo/bson"
	cspb "protocol"
)

/*所有的数据库模型都在这*/

/*mysql*/
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

//好友消息
type Messages struct {
	Id          int32
	Uid         int32
	Fromuid     int32
	Touid       int32
	Messagetype int32 /*1：赠送 2,接收*/
	Fromname    string
	FromStarId  int32
	Fromstar    string
	ElementType int32 //1:小红花，2：卡片，3：加好友
	CardId      int32
	Number      int32
	IsFinish    int32
	Time        int64
}

//排名
type Ranking struct {
	Id           int32
	Uid          int32
	Name         string
	Medal        int32
	Level        int32
	StarId       int32
	MedalLevelId int32 /*头衔的等级*/
	Time         int64
}
type FriendRank struct {
	Id           int32
	Uid          int32
	Name         string
	Medal        int32
	Level        int32
	StarId       int32
	MedalLevelId int32 /*头衔的等级*/
	Time         int64
}

/*成就*/
type Achievement struct {
	Id             int32
	Uid            int32
	AchievementSum int32 //总共多少成就
	AchievementId  int32 /*完成成就的Id*/
	Time           int64
}

/*mongo*/
/*人物的属性*/
type Player struct {
	Saccount           bson.ObjectId          `bson:"_id"`
	Caccount           string                 `bson:"c_account"`
	Uid                int64                  `bson:"uid"`
	Name               string                 `bson:"name"`
	Gold               int32                  `bson:"gold"`
	Flower             int32                  `bson:"flower"`
	Diamond            int32                  `bson:"diamond"`
	StarId             int32                  `bson:"starid"`
	Star               []*StarDate            `bson:"star"`
	SolutionPool       int32                  `bson:"experience_pool"`
	RegistTime         int64                  `bson:"regist_time"`
	FriendList         []*FriendListData      `bson:"FriendList"`
	ApplyFriendList    []*ApplyFriendListData `bson:"ApplyFriendList"`
	Recharge           []*RechargeData        `bson:"recharge"`
	WonderfulFriends   WonderfulFriendsData   `bson:"WonderfulFriends"`
	RechargeFlow       RechargeFlowData       `bson:"RechargeFlow"`
	LastSignInTime     int64                  `bson:"last_signin_time"`
	FreeSignInOperTime int64                  `bson:"free_signin_oper_time"`
	Levels             []*cspb.CSStageNtf     `bson:"Levels"`
	Money              *cspb.CSMoneyReq       `bson:"Money"`
}
type StarDate struct {
	StarId       int32       `bson:"starid"`
	Starname     string      `bson:"starname"`
	Level        int32       `bson:"level"`
	Solution     int32       `bson:"experience"`
	Dress        int32       `bson:"dress"`
	Dressname    string      `bson:"dressname"`
	Fighting     int32       `bson:"fighting"`
	IsActive     int32       `bson:"is_active"` //是否选择
	Medal        int32       `bson:"medal"`
	MedalLevelId int32       `bson:"medal_level_id"`
	Cards        []*CardData `bson:"cards"`
}
type CardData struct {
	CardId  int32 `bson:"card_id"`
	CardNum int32 `bson:"card_num"`
}
type RechargeData struct {
	Uid          int64  `bson:"uid"`
	Rmb          int32  `bson:"rmb"`
	Dollar       int32  `bson:"dollar"`
	GoodsType    int32  `bson:"goods_type"`
	GoodsSubType int32  `bson:"goods_sub_type"`
	GoodsNum     int32  `bson:"goods_num"`
	RechargeTime int64  `bson:"recharge_time"`
	Version      string `bson:"version"`
	Code         string `bson:"code"`
	Channel      string `bson:"channel"`
}

/*系统*/
type System struct {
	Version string `bson:"version"`
}

/*老版*/
// type Player struct {
// 	Saccount           bson.ObjectId          `bson:"_id"`
// 	Caccount           string                 `bson:"c_account"`
// 	Uid                int64                  `bson:"uid"`
// 	WonderfulFriends   WonderfulFriendsData   `bson:"WonderfulFriends"`
// 	RechargeFlow       RechargeFlowData       `bson:"RechargeFlow"`
// 	RegistTime         int64                  `bson:"regist_time"`
// 	LastSignInTime     int64                  `bson:"last_signin_time"`
// 	FreeSignInOperTime int64                  `bson:"free_signin_oper_time"`
// 	Levels             []*cspb.CSStageNtf     `bson:"Levels"`
// 	Money              *cspb.CSMoneyReq       `bson:"Money"`
// 	FriendList         []*FriendListData      `bson:"FriendList"`
// 	ApplyFriendList    []*ApplyFriendListData `bson:"ApplyFriendList"`
// }
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
	Friendid   int32  `bson:"friendid"`
	FriendName string `bson:"friendname"`
	IsActive   int32  `bson:"isActive"`
	Accepttime int64  `bson:"accepttime"`
}
type ApplyFriendListData struct {
	Applyuid     int32 `bson:"applyuid"`
	IsAccept     int32 `bson:"isAccept"`
	Isrefuse     int32 `bson:"isrefuse"`
	Applytime    int64 `bson:"applytime"`
	Oprationtime int64 `bson:"oprationtime"`
}
type Pet struct {
	Account       string     `bson:"account"`
	PetId         int32      `bson:"pet_id"`
	PetLevel      int32      `bson:"pet_level"`
	PetCurExp     int32      `bson:"pet_cur_exp"`
	PetTotalExp   int32      `bson:"pet_total_exp"`
	PetStarLevel  int32      `bson:"pet_star_level"`
	Petmedallevel int32      `bson:"pet_medal_level"`
	PetmedalNum   int32      `bson:"pet_medal_num"`
	Petcard       PetcardNtf `bson:"pet_card"`
	DressId       int32      `bson:"dress_id"`
}
type PetcardNtf []struct {
	CardId  int32 `bson:"cardid"`
	CardNum int32 `bson:"cardnum"`
}

type Chip struct {
	Account  string `bson:"account"`
	ChipId   int32  `bson:"chip_id"`
	ChipType int32  `bson:"chip_type"`
	ChipNum  int32  `bson:"chip_num"`
}

func init() {

	orm.RegisterModel(new(Messagecenter), new(Userscore), new(Userinfo), new(Ranking), new(Messages), new(FriendRank))

}
