package model

import "github.com/go-xorm/xorm"

const (
	TRADE_STATES_PART = 1
	TRADE_STATES_ALL  = 2
	TRADE_STATES_DEL  = 3
)

type Trade struct {
	TradeId      int    `xorm:"not null pk autoincr comment('交易表的id') INT(11)"`
	TradeNo      string `xorm:"comment('订单号') unique(uni_reade_no) VARCHAR(32)"`
	Uid          uint64 `xorm:"comment('买家uid') index unique(uni_reade_no) BIGINT(11)"`
	TokenId      int    `xorm:"comment('主货币id') index INT(11)"`
	TokenTradeId int    `xorm:"comment('交易币种') INT(11)"`
	Price        int64  `xorm:"comment('价格') BIGINT(20)"`
	Num          int64  `xorm:"comment('数量') BIGINT(20)"`
	Money        int64  `xorm:"BIGINT(20)"`
	Fee          int64  `xorm:"comment('手续费') BIGINT(20)"`
	Opt          int    `xorm:"comment(' buy  1或sell 2') index TINYINT(4)"`
	DealTime     int64  `xorm:"comment('成交时间') BIGINT(20)"`
	States       int    `xorm:"comment('0是挂单，1是部分成交,2成交， -1撤销') INT(11)"`
}

func (s *Trade) Insert(session *xorm.Session, t ...*Trade) (err error) {
	_, err = session.Insert(t)
	return
}
