package model

import (
	"digicon/common/convert"
	"digicon/common/model"
	. "digicon/token_service/dao"
	"github.com/go-xorm/xorm"
	"github.com/liudng/godump"
	log "github.com/sirupsen/logrus"
	"time"
)

/*
const (
	TRADE_STATES_PART = 1 //部分成交
	TRADE_STATES_ALL  = 2 //全部成交
	TRADE_STATES_DEL  = 3 //撤单
)
*/
type Trade struct {
	TradeId          int    `xorm:"not null pk autoincr comment('交易表的id') INT(11)"`
	TradeNo          string `xorm:"comment('订单号') unique(uni_reade_no) VARCHAR(32)"`
	Uid              uint64 `xorm:"comment('买家uid') index unique(uni_reade_no) BIGINT(11)"`
	TokenId          int    `xorm:"comment('主货币id') index INT(11)"`
	TokenTradeId     int    `xorm:"comment('交易币种') INT(11)"`
	TokenAdmissionId int    `xorm:"comment('入账token_id') index INT(11)"`
	Symbol           string `xorm:"comment('交易队') VARCHAR(32)"`
	Price            int64  `xorm:"comment('价格') BIGINT(20)"`
	Num              int64  `xorm:"comment('数量') BIGINT(20)"`
	//Balance      int64  `xorm:"BIGINT(20)"`
	EntrustId string `xorm:"comment('委托ID')  VARCHAR(32)"`
	Fee       int64  `xorm:"comment('手续费数量') BIGINT(20)"`
	Opt       int    `xorm:"comment(' buy  1或sell 2') index TINYINT(4)"`
	DealTime  int64  `xorm:"comment('成交时间') BIGINT(20)"`
	//States    int    `xorm:"comment('0是挂单，1是部分成交,2成交， -1撤销') INT(11)"`
	FeeCny   int64 `xorm:"comment('手续费人民币') BIGINT(20)"`
	TotalCny int64 `xorm:"comment('成交额人民币') BIGINT(20)"`
}

func (s *Trade) Insert(session *xorm.Session, t ...*Trade) (err error) {
	defer func() {
		if err != nil {
			for _, v := range t {
				log.WithFields(log.Fields{
					"uid":      v.Uid,
					"opt":      v.Opt,
					"token_id": v.TokenId,
					"price":    v.Price,
					"fee":      v.Fee,
					"trade_no": v.TradeNo,
				}).Errorf("inset  money record error %s", err.Error())
			}
		}
	}()
	_, err = session.Insert(t)
	return
}

func (s *Trade) GetUserTradeList(pageIndex, pageSize int, uid uint64) (*model.ModelList, []*Trade, error) {
	engine := DB.GetMysqlConn()

	query := engine.Where("uid=?", uid).Desc("deal_time", "trade_id")
	tempQuery := *query

	count, err := tempQuery.Count(s)
	if err != nil {
		return nil, nil, err
	}

	// 获取分页
	offset, modelList := model.Paging(pageIndex, pageSize, int(count))

	var list []*Trade
	err = query.Select("*").Limit(modelList.PageSize, offset).Find(&list)
	if err != nil {
		return nil, nil, err
	}
	modelList.Items = list

	return modelList, list, nil
}

func GetUserTradeByEntrustId(entrust_id string) (g []*Trade, err error) {
	g = make([]*Trade, 0)
	err = DB.GetMysqlConn().Where("entrust_id=?", entrust_id).Find(&g)
	if err != nil {
		log.Errorln(err.Error())
		return
	}
	return
}

func CaluateAvgPrice(t []*Trade) int64 {
	var amount, sum int64

	for _, v := range t {
		amount += convert.Int64MulInt64By8Bit(v.Num, v.Price)
		sum += v.Num
	}
	tt := convert.Int64DivInt64By8Bit(amount, sum)

	godump.Dump(tt)
	return convert.Int64DivInt64By8Bit(amount, sum)
}

func Test2(beid, endid int64) {
	log.Infof("begin id=%d,endid=%d", beid, endid)
	g := make([]*Trade, 0)
	err := DB.GetMysqlConn().Where("trade_id>=? and trade_id<?", beid, endid).Find(&g)
	if err != nil {
		log.Fatalln(err.Error())
		return
	}
	if len(g) == 0 {
		return
	}

	for _, v := range g {
		log.Infof("process id %d", v.TradeId)
		if v.Opt == 1 {
			v.TokenAdmissionId = v.TokenTradeId
			_, err = DB.GetMysqlConn().Where("trade_id=?", v.TradeId).Cols("token_admission_id").Update(v)
			if err != nil {
				log.Fatalln(err.Error())
				return
			}
		} else {
			v.TokenAdmissionId = v.TokenId
			_, err = DB.GetMysqlConn().Where("trade_id=?", v.TradeId).Cols("token_admission_id").Update(v)
			if err != nil {
				log.Fatalln(err.Error())
				return
			}
		}
	}

	Test2(beid+1000, endid+1000)
	/*
		lastt := stime + 43200

		log.Infof("beigin %d",time.Now().Unix())
		_,err:=DB.GetMysqlConn().Exec("call statisticss_daily_fee()")
		if err!=nil {
			log.Fatalln(err.Error())
		}

		log.Infof("end %d",time.Now().Unix())

	*/
}

func Testj() {
	g := make([]string, 0)

	g = append(g, "1533285163_10")
	g = append(g, "1533283556_4")
	g = append(g, "1533284027_2")
	g = append(g, "1533284284_2")
	g = append(g, "1533284307_3")

	en2 := make(map[string]*EntrustDetail, 0)

	err := DB.GetMysqlConn().In("entrust_id", g).Cols().Find(&en2)
	if err != nil {
		log.Errorln(err)
		return
	}

	for _, v := range en2 {
		log.Infof("entrust_id %s", v.EntrustId)
	}

	en := make([]*EntrustDetail, 0)
	for _, k := range g {
		v, ok := en2[k]
		if ok {
			en = append(en, v)
		}
	}

	for _, v := range en {
		log.Infof("entrust_id %s", v.EntrustId)
	}

}

func ReverseTrade(index int) {
	//var ok bool
	var err error
	g := make([]*Trade, 0)
	err = DB.GetMysqlConn().Where("trade_id>? and trade_id<=?", (index-1)*1000, index*1000).Find(&g)
	if err != nil {
		log.Fatalln(err.Error())
	}
	if len(g) > 0 {

		log.Infof("begin process trade_id index  %d ,time %d", index, time.Now().Unix())
		for k, v := range g {
			if v.TradeId%2 == 1 {
				t := g[k+1]
				if v.TradeNo != t.TradeNo {
					log.Fatalln("err ")
				}

				trade_num := t.Num
				trade_fee := t.Fee
				trade_fee_cny := t.FeeCny

				main_num := v.Num
				main_fee := v.Fee
				main_fee_cny := v.FeeCny

				t.Num = main_num
				t.Fee = main_fee
				t.FeeCny = main_fee_cny

				v.Num = trade_num
				v.Fee = trade_fee
				v.FeeCny = trade_fee_cny

				_, err = DB.GetMysqlConn().Where("trade_id=?", t.TradeId).Cols("num", "fee", "fee_cny").Update(t)
				if err != nil {
					log.Fatalln(err.Error())
				}
				_, err = DB.GetMysqlConn().Where("trade_id=?", v.TradeId).Cols("num", "fee", "fee_cny").Update(v)
				if err != nil {
					log.Fatalln(err.Error())
				}
			}

		}

		log.Infof("end process trade_id index  %d ,time %d ", index, time.Now().Unix())
	} else {
		log.Infof("final process time %d ", time.Now().Unix())
		return
	}

	ReverseTrade(index + 1)
}
