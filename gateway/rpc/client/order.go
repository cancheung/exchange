package client

import (
	"context"
	proto "digicon/proto/rpc"
	log "github.com/sirupsen/logrus"
)

func (s *CurrencyRPCCli) CallOrdersList(req *proto.OrdersListRequest) (rsp *proto.OrdersListResponse, err error) {
	rsp, err = s.conn.OrdersList(context.TODO(), req)
	if err != nil {
		log.Errorln(err.Error())
		return
	}
	return
}

func (s *CurrencyRPCCli) CallDeleteOrder(req *proto.OrderRequest) (rsp *proto.OrderResponse, err error) {
	rsp, err = s.conn.DeleteOrder(context.TODO(), req)
	if err != nil {
		log.Errorln(err.Error())
		return
	}
	return
}

func (s *CurrencyRPCCli) CallCancelOrder(req *proto.CancelOrderRequest) (rsp *proto.OrderResponse, err error) {
	rsp, err = s.conn.CancelOrder(context.TODO(), req)
	if err != nil {
		log.Errorln(err.Error())
		return
	}
	return
}

func (s *CurrencyRPCCli) CallConfirmOrder(req *proto.ConfirmOrderRequest) (rsp *proto.OrderResponse, err error) {

	rsp, err = s.conn.ConfirmOrder(context.TODO(), req)
	if err != nil {
		log.Errorln(err.Error())
		return
	}
	return
}

func (s *CurrencyRPCCli) CallReadyOrder(req *proto.OrderRequest) (rsp *proto.OrderResponse, err error) {
	rsp, err = s.conn.ReadyOrder(context.TODO(), req)
	if err != nil {
		log.Errorln(err.Error())
		return
	}
	return
}

func (s *CurrencyRPCCli) CallAddOrder(req *proto.AddOrderRequest) (rsp *proto.OrderResponse, err error) {
	rsp, err = s.conn.AddOrder(context.TODO(), req)
	if err != nil {
		log.Errorln(err.Error())
		return
	}
	return
}

func (s *CurrencyRPCCli) CallGetTradeDetail(req *proto.TradeDetailRequest) (rsp *proto.TradeDetailResponse, err error) {
	rsp, err = s.conn.TradeDetail(context.TODO(), req)
	if err != nil {
		log.Errorln(err.Error())
		return
	}
	return
}

func (s *CurrencyRPCCli) CallGetTradeHistory(req *proto.GetTradeHistoryRequest) (rsp *proto.OtherResponse, err error) {
	rsp, err = s.conn.GetTradeHistory(context.TODO(), req)
	if err != nil {
		log.Errorln(err.Error())
		return
	}
	return
}

func (s *CurrencyRPCCli) CallGetRecentTransactionPrice(req *proto.GetRecentTransactionPriceRequest) (rsp *proto.OtherResponse, err error) {
	rsp, err = s.conn.GetRecentTransactionPrice(context.TODO(), req)
	if err != nil {
		log.Errorln(err.Error())
		return
	}
	return
}


func (s *CurrencyRPCCli) CallGetUserBalanceUids(req *proto.GetUserBalanceUids)(rsp *proto.UserBalancesResponse, err error)  {
	rsp, err = s.conn.GetUsersBalance(context.TODO(), req)
	if err != nil {
		log.Errorln(err)
		return
	}
	return
}

func (s *CurrencyRPCCli) CallAdminConfirm(req *proto.ConfirmOrderRequest) ( rsp *proto.OrderResponse, err error) {
	rsp, err = s.conn.AdminConfirm(context.TODO(), req)
	if err != nil {
		log.Errorln(err)
		return
	}
	return
}

func (s *CurrencyRPCCli) CallAdminCancel(req *proto.CancelOrderRequest)(rsp *proto.OrderResponse, err error) {
	rsp, err = s.conn.AdminCancel(context.TODO(), req)
	if err != nil {
		log.Errorln(err)
		return
	}
	return
}