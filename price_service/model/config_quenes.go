package model

import (
	. "digicon/price_service/dao"
	proto "digicon/proto/rpc"
	log "github.com/sirupsen/logrus"
)

type ConfigQuenes struct {
	Id           int64  `xorm:"pk autoincr BIGINT(20)"`
	TokenId      int    `xorm:"comment('交易币') unique(union_quene_id) INT(11)"`
	TokenTradeId int    `xorm:"comment('实际交易币') unique(union_quene_id) INT(11)"`
	Switch       int    `xorm:"comment('开关0关1开') TINYINT(4)"`
	Name         string `xorm:"comment('USDT/BTC') VARCHAR(32)"`
	Price        int64  `xorm:"BIGINT(20)"`
}

/*
type ConfigQuenes struct {
	Id           int64  `xorm:"pk autoincr BIGINT(20)"`
	TokenId      int    `xorm:"comment('交易币') unique(union_quene_id) INT(11)"`
	TokenTradeId int    `xorm:"comment('实际交易币') unique(union_quene_id) INT(11)"`
	Switch       int    `xorm:"comment('开关0关1开') TINYINT(4)"`
	Price        int64  `xorm:"comment('初始价格') BIGINT(20)"`
	Name         string `xorm:"comment('USDT/BTC') VARCHAR(32)"`
	Scope        string `xorm:"comment('振幅') DECIMAL(6,2)"`
	Low          int64  `xorm:"comment('最低价') BIGINT(20)"`
	High         int64  `xorm:"comment('最高价') BIGINT(20)"`
	Amount       int64  `xorm:"comment('成交量') BIGINT(20)"`
}
*/
func GetAllQuenes() []*ConfigQuenes {
	t := make([]*ConfigQuenes, 0)
	err := DB.GetMysqlConn2().Where("switch=1").Find(&t)
	if err != nil {
		log.Errorln(err.Error())
		return nil
	}
	return t
}

func InitAllQuene() {
	k := make([]*ConfigQuenes, 0)
	err := DB.GetMysqlConn2().Where("switch=1").Find(&k)
	if err != nil {
		log.Errorln(err.Error())
		return
	}
	for _, v := range k {
		ConfigQueneInit[v.Name] = v

	}

	return
}

func InitConfig() {
	g := GetAllQuenes()
	InitAllQuene()
	for _, v := range g {
		meta := &proto.ConfigQueneBaseData{
			TokenId:      int32(v.TokenId),
			TokenTradeId: int32(v.TokenTradeId),
			Name:         v.Name,
		}

		ConfigQueneData[v.Name] = meta

		k, ok := ConfigQueneArr[int32(v.TokenId)]
		if ok {
			k = append(k, meta)
			ConfigQueneArr[int32(v.TokenId)] = k
		} else {
			a := make([]*proto.ConfigQueneBaseData, 0)
			a = append(a, meta)
			ConfigQueneArr[int32(v.TokenId)] = a
		}

		//m := GetConfigQuenesByType(1)
	}
	//m := GetConfigQuenesByType(1)

	/*
		h := GetCnyQuenes()
		for _, v := range h {
			ConfigCny[int32(v.TokenId)] = &proto.CnyPriceBaseData{
				TokenId:  int32(v.TokenId),
				CnyPrice: v.Price,
			}
		}
	*/
}

var ConfigQueneData map[string]*proto.ConfigQueneBaseData

var ConfigQueneArr map[int32][]*proto.ConfigQueneBaseData

//var ConfigCny map[int32]*proto.CnyPriceBaseData

var ConfigQueneInit map[string]*ConfigQuenes

var IsFinish bool

/*
func GetTokenCnyPrice2(token_id int32) int64 {
	g, ok := ConfigCny[token_id]
	if ok {
		return g.CnyPrice
	}
	return 0
}
*/
func GetConfigQuenesByType(token_id int32) []*proto.ConfigQueneBaseData {
	g, ok := ConfigQueneArr[token_id]
	if ok {
		return g
	}
	return nil
}

func init() {
	ConfigQueneData = make(map[string]*proto.ConfigQueneBaseData, 0)
	ConfigQueneArr = make(map[int32][]*proto.ConfigQueneBaseData, 0)

	//	ConfigCny = make(map[int32]*proto.CnyPriceBaseData, 0)

	ConfigQueneInit = make(map[string]*ConfigQuenes)
	IsFinish = false
}
