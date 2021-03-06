package model

import (
	. "digicon/token_service/dao"
	log "github.com/sirupsen/logrus"
)

/*
type Price struct {
	Id     int64 `xorm:"BIGINT(20)"`
	Open   int64 `xorm:"comment('开盘价') BIGINT(20)"`
	Close  int64 `xorm:"comment('收盘价') BIGINT(20)"`
	Low    int64 `xorm:"comment('最低价') BIGINT(20)"`
	High   int64 `xorm:"comment('最高价') BIGINT(20)"`
	Amount int64 `xorm:"comment('成交量') BIGINT(20)"`
	Vol    int64 `xorm:"comment('成交额') BIGINT(20)"`
	Count  int64 `xorm:"BIGINT(20)"`
}

*/

type Price struct {
	Id          int64  `xorm:"index(keep) BIGINT(20)"`
	Symbol      string `xorm:"index(keep) VARCHAR(32)"`
	Price       int64  `xorm:"BIGINT(20)"`
	CreatedTime int64  `xorm:"BIGINT(20)"`
	Amount      int64  `xorm:"BIGINT(20)"`
	Vol         int64  `xorm:"BIGINT(20)"`
	Count       int64  `xorm:"BIGINT(20)"`
	UsdVol      string  `xorm:"varchar(96)"`
}

func GetPrice(symbol string) (*Price, bool) {
	m := &Price{}
	ok, err := DB.GetMysqlConn2().Where("symbol=?", symbol).Desc("created_time").Limit(1, 0).Get(m)
	if err != nil {
		log.Fatalln("err data price")
	}

	log.Infof("init price is %d", m.Price)
	return m, ok
}
