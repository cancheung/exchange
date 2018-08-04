package cron

import (
	"digicon/common/constant"
	"digicon/currency_service/dao"
	"digicon/currency_service/model"
	proto "digicon/proto/rpc"
	"encoding/json"
)

//代币划转到法币事件处理
func HandlerTransferFromToken() {
	rdsClient := dao.DB.GetRedisConn()
	userCurrencyMD := new(model.UserCurrency)

	for {
		msgBody, err := rdsClient.LPop(constant.RDS_TOKEN_TO_CURRENCY_TODO).Bytes()
		if err != nil {
			continue
		}

		msg := &proto.TransferToCurrencyTodoMessage{}
		err = json.Unmarshal(msgBody, msg)
		if err != nil {
			continue
		}

		userCurrencyMD.TransferFromToken(msg)
	}
}
