package cs_handle

import (
	// "fmt"
	"github.com/astaxie/beego"
	// "time"
	// models "tuojie.com/piggo/quickstart.git/models"
)
import cspb "protocol"

// import proto "code.google.com/p/goprotobuf/proto"

// import db "tuojie.com/piggo/quickstart.git/db/collection"
// import "labix.org/v2/mgo/bson"

// import db_session "tuojie.com/piggo/quickstart.git/db/session"

func FriendmessageHandle(
	req *cspb.CSPkg,
	res_list *cspb.CSPkgList) int32 {
	beego.Info("*********FriendmessageHandle Start**********")
	req_data := req.GetBody().GetFriendmessageReq()
	beego.Info(req_data)
	ret := int32(1)
	isGive := int32(1)
	uid := int32(2885377)
	res_data := new(cspb.CSFriendmessageRes)
	messageId := req_data.GetMessageId()
	friendListId := req_data.GetFriendListId()
	elementType := req_data.GetElementType()
	*res_data = cspb.CSFriendmessageRes{
		IsGive:       &isGive,
		Uid:          &uid,
		MessageId:    &messageId,
		FriendListId: &friendListId,
		ElementType:  &elementType,
	}

	res_pkg_body := new(cspb.CSBody)
	*res_pkg_body = cspb.CSBody{

		FriendmessageRes: res_data,
	}
	res_list = makeCSPkgList(int32(cspb.Command_kFriendmessageRes),
		res_pkg_body, res_list)
	return ret

}
