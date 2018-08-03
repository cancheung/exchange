package model

type OutUserCurrencyHistory struct {
	Id          int    `json:"id"                  xorm:"not null pk autoincr comment('ID') INT(10)"`
	Uid         int32  `json:"uid"                 xorm:"not null default 0 INT(10)"`
	TradeUid    int32  `json:"trade_uid"           xorm:"not null default 0 INT(10)"`
	OrderId     string `json:"order_id"            xorm:"not null default '' comment('订单ID') VARCHAR(64)"`
	TokenId     int    `json:"token_id"            xorm:"not null default 0 comment('货币类型') INT(10)"`
	Num         int64  `json:"num"                 xorm:"not null default 0 comment('数量') BIGINT(64)"`
	Fee         int64  `json:"fee"                 xorm:"not null default 0 comment('手续费用') BIGINT(64)"`
	Surplus     int64  `json:"surplus"             xorm:"comment('账户余额') BIGINT(64)"`
	Operator    int    `json:"operator"            xorm:"not null default 0 comment('操作类型 1订单转入 2订单转出 3充币 4提币 5冻结') TINYINT(2)"`
	Address     string `json:"address"             xorm:"not null default '' comment('提币地址') VARCHAR(255)"`
	States      int    `json:"states"              xorm:"not null default 0 comment('订单状态: 0删除 1待支付 2待放行(已支付) 3确认支付(已完成) 4取消') TINYINT(2)"`
	CreatedTime string `json:"created_time"        xorm:"not null comment('创建时间') DATETIME"`
	UpdatedTime string `json:"updated_time"        xorm:"comment('修改时间') DATETIME"`
}

func (OutUserCurrencyHistory) TableName() string {
	return "user_currency_history"
}