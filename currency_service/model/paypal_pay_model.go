package model

import (
	"digicon/currency_service/dao"
	. "digicon/proto/common"
	proto "digicon/proto/rpc"
	"time"

	"digicon/common/constant"
	"digicon/currency_service/rpc/client"
	"fmt"
	log "github.com/sirupsen/logrus"
)

type UserCurrencyPaypalPay struct {
	Uid        uint64 `xorm:"not null pk default 0 comment('用户uid') INT(10)"        json:"uid"`
	Paypal     string `xorm:"not null default '' comment('paypal 账号') VARCHAR(20)"  json:"paypal"`
	CreateTime string `xorm:"not null comment('创建时间') DATETIME"`
	UpdateTime string `xorm:"not null comment('修改时间') DATETIME"`
}

func (pal *UserCurrencyPaypalPay) SetPaypal(req *proto.PaypalRequest) (int32, error) {

	/////////////////  1.  验证  验证码 /////////////////////////
	rsp, err := client.InnerService.UserSevice.CallAuthVerify(&proto.AuthVerifyRequest{
		Uid:      req.Uid,
		Code:     req.Verify,
		AuthType: constant.SMS_PAYPAL_PAY, // 设置银行卡支付 10
	})
	fmt.Println("=========================", rsp)
	if err != nil {
		log.Errorln(err.Error())
		return ERRCODE_SMS_CODE_DIFF, err
	}
	if rsp.Code != ERRCODE_SUCCESS {
		return ERRCODE_SMS_CODE_DIFF, err
	}

	//调用实名接口

	//检查数据库是否存在该条记录
	engine := dao.DB.GetMysqlConn()
	has, err := engine.Exist(&UserCurrencyPaypalPay{
		Uid: req.Uid,
	})
	if err != nil {
		return ERRCODE_UNKNOWN, err
	}
	current := time.Now().Format("2006-01-02 15:04:05")
	if has {
		pal.UpdateTime = current
		_, err := engine.Update(pal)
		if err != nil {
			return ERRCODE_UNKNOWN, err
		}
		//return ERRCODE_ACCOUNT_EXIST, errors.New("account already exist!!")
	} else {
		_, err := engine.InsertOne(&UserCurrencyPaypalPay{
			Uid:        req.Uid,
			Paypal:     req.Paypal,
			CreateTime: current,
			UpdateTime: current,
		})
		if err != nil {
			return ERRCODE_ACCOUNT_EXIST, err
		}
	}
	return ERRCODE_SUCCESS, nil
}

func (pal *UserCurrencyPaypalPay) GetByUid(uid uint64) (err error) {
	engine := dao.DB.GetMysqlConn()
	_, err = engine.Where("uid =?", uid).Get(pal)
	return
}
