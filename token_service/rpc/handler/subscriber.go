package handler

import (
	proto "digicon/proto/rpc"
	"golang.org/x/net/context"
	log "github.com/sirupsen/logrus"
	"digicon/token_service/model"
)

type Subscriber struct{}

func (sub *Subscriber) Process(ctx context.Context, data *proto.CnyPriceResponse) error {
	log.Println("Picked up a new message")
	for _,v:=range data.Data  {
		model.CnyPriceMap[v.TokenId]=v
	}
	return nil
}