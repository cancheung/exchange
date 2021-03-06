package dao

import (
	"digicon/config_service/conf"
	log "github.com/sirupsen/logrus"
	"github.com/hashicorp/consul/api"
)

type ConsulCli struct {
	ccon *api.Client
}

func NewConsulCli() *ConsulCli {
	addr := conf.Cfg.MustValue("consul", "addr")
	token := conf.Cfg.MustValue("consul", "token")

	client, err := api.NewClient(&api.Config{
		Token:   token,
		Address: addr,
	})
	if err != nil {
		log.Fatal("new consul client error!")
	}

	return &ConsulCli{
		ccon: client,
	}
}

func (s *Dao) GetConsulCli() *api.Client {
	return s.consul.ccon
}
