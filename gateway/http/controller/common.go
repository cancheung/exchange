package controller

import (
	"digicon/gateway/rpc"
	. "digicon/proto/common"
	log "github.com/sirupsen/logrus"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CommonGroup struct{}

func (s *CommonGroup) Router(r *gin.Engine) {
	user := r.Group("/common")
	{
		user.GET("/token_list", s.GetToknesList)
	}
}

func (s *CommonGroup) GetToknesList(c *gin.Context) {
	ret := NewPublciError()
	defer func() {
		c.JSON(http.StatusOK, ret.GetResult())
	}()

	rsp, err := rpc.InnerService.UserSevice.CallTokensList()
	if err != nil {
		log.Errorf(err.Error())
		ret.SetErrCode(ERRCODE_UNKNOWN, err.Error())
		return
	}
	ret.SetErrCode(rsp.Err)
	ret.SetDataSection("list", rsp.Data)
}
