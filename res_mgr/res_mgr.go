//带>>>的注释不要删掉,自动添加代码的工具需要
//本文件不要修改最好!

package res_mgr

import "os"
import "io/ioutil"
import log "code.google.com/p/log4go"
import proto "code.google.com/p/goprotobuf/proto"
import "resource"

func resLoadFile(file string, pb proto.Message) {
	f, err := os.Open("./res/" + file)
	if err != nil {
		log.Error("Load res file %s error %v\n", file, err)
		os.Exit(1)
	}
	defer f.Close()
	content, err := ioutil.ReadAll(f)
	if err != nil {
		log.Error("Read content error %v\n", err)
		os.Exit(1)
	}
	err = proto.Unmarshal(content, pb)
	if err != nil {
		log.Error("Unmarshal content error %v\n", err)
		os.Exit(1)
	}

	log.Fine("load res %s success info:%s\n", file, pb.String())
}

//带>>>的注释不要删掉,自动添加代码的工具需要
func Init() {

	resLoadFile("pet.bytes", &PetData)
	resLoadFile("chip.bytes", &ChipData)
	resLoadFile("randomgoods.bytes", &RandomgoodsData)
	resLoadFile("petlevel.bytes", &PetlevelData)
	resLoadFile("dailylogin.bytes", &DailyloginData)
	resLoadFile("randomitem.bytes", &RandomitemData)
	resLoadFile("keyvalue.bytes", &KeyvalueData)
	resLoadFile("star.bytes", &StarData)
	resLoadFile("loadingtiplist.bytes", &LoadingtiplistData)
	resLoadFile("updateconfigs.bytes", &UpdateconfigsData)
} //>>>func

var PetData resource.PetArray
var ChipData resource.ChipArray
var RandomgoodsData resource.RandomgoodsArray
var PetlevelData resource.PetlevelArray
var DailyloginData resource.DailyloginArray
var RandomitemData resource.RandomitemArray
var KeyvalueData resource.KeyvalueArray

var StarData resource.StarArray
var LoadingtiplistData resource.LoadingtiplistArray
var UpdateconfigsData resource.UpdateconfigsArray

//>>>var
