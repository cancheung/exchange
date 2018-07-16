package controller

import (
	. "digicon/gateway/log"
	"digicon/gateway/rpc"
	. "digicon/proto/common"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ActionGroup struct{}

func (s *ActionGroup) Router(r *gin.Engine) {
	action := r.Group("/action")
	{
		action.POST("/get_google_key", s.GetGoogleAuthCode)
		action.POST("/auth_google_code", s.AuthGoogleCode)
		action.POST("/del_google_code", s.DelGoogleCode)
		action.POST("/reset_google_code", s.ResetGoogleCode)
		action.GET("/get_user_info", s.GetUserBaseInfo)
		action.GET("/get_user_real", s.GetUserRealName)
		action.GET("/get_user_invite", s.GetUserInvite)
		action.GET("/get_ip_record", s.GetIpRecord)
	}
}

func (s *ActionGroup) GetGoogleAuthCode(c *gin.Context) {
	ret := NewPublciError()
	defer func() {
		c.JSON(http.StatusOK, ret.GetResult())
	}()

	type GoogleAuthCodeParam struct {
		Uid uint64 `form:"uid" binding:"required"`
	}
	var param GoogleAuthCodeParam

	if err := c.ShouldBind(&param); err != nil {
		Log.Errorf(err.Error())
		ret.SetErrCode(ERRCODE_PARAM, err.Error())
		return
	}

	rsp, err := rpc.InnerService.UserSevice.CallGoogleSecretKey(param.Uid)
	if err != nil {
		ret.SetErrCode(ERRCODE_UNKNOWN, err.Error())
		return
	}
	ret.SetErrCode(rsp.Err, rsp.Message)
	ret.SetDataSection("key", rsp.SecretKey)
	ret.SetDataSection("url", rsp.Url)
}

func (s *ActionGroup) AuthGoogleCode(c *gin.Context) {
	ret := NewPublciError()
	defer func() {
		c.JSON(http.StatusOK, ret.GetResult())
	}()

	type GoogleAuthCodeParam struct {
		Uid  uint64 `form:"uid" binding:"required"`
		Code uint32 `form:"code" binding:"required"`
	}
	var param GoogleAuthCodeParam

	if err := c.ShouldBind(&param); err != nil {
		Log.Errorf(err.Error())
		ret.SetErrCode(ERRCODE_PARAM, err.Error())
		return
	}

	rsp, err := rpc.InnerService.UserSevice.CallAuthGoogleSecretKey(param.Uid, param.Code)
	if err != nil {
		ret.SetErrCode(ERRCODE_UNKNOWN, err.Error())
		return
	}
	ret.SetErrCode(rsp.Err, rsp.Message)
}

func (s *ActionGroup) DelGoogleCode(c *gin.Context) {
	ret := NewPublciError()
	defer func() {
		c.JSON(http.StatusOK, ret.GetResult())
	}()
	type GoogleDelCodeParam struct {
		Uid  uint64 `form:"uid" binding:"required"`
		Code uint32 `form:"code" binding:"required"`
	}
	var param GoogleDelCodeParam

	if err := c.ShouldBind(&param); err != nil {
		Log.Errorf(err.Error())
		ret.SetErrCode(ERRCODE_PARAM, err.Error())
		return
	}
	rsp, err := rpc.InnerService.UserSevice.CallDelGoogleSecretKey(param.Uid, param.Code)
	if err != nil {
		ret.SetErrCode(ERRCODE_UNKNOWN, err.Error())
		return
	}
	ret.SetErrCode(rsp.Err, rsp.Message)
}

func (s *ActionGroup) ResetGoogleCode(c *gin.Context) {
	ret := NewPublciError()
	defer func() {
		c.JSON(http.StatusOK, ret.GetResult())
	}()
	type GoogleDelCodeParam struct {
		Uid  uint64 `form:"uid" binding:"required"`
		Auth uint32 `form:"auth" binding:"required"`
		Sms  string `form:"sms" binding:"required"`
	}
	var param GoogleDelCodeParam

	if err := c.ShouldBind(&param); err != nil {
		Log.Errorf(err.Error())
		ret.SetErrCode(ERRCODE_PARAM, err.Error())
		return
	}

	rsp, err := rpc.InnerService.UserSevice.CallResetGoogleSecretKey(param.Uid, param.Auth, param.Sms)
	if err != nil {
		ret.SetErrCode(ERRCODE_UNKNOWN, err.Error())
		return
	}
	ret.SetErrCode(rsp.Err, rsp.Message)
}

func (s *ActionGroup) GetUserBaseInfo(c *gin.Context) {
	ret := NewPublciError()
	defer func() {
		c.JSON(http.StatusOK, ret.GetResult())
	}()
	type GetUserBaseInfoParam struct {
		Uid uint64 `form:"uid" binding:"required"`
	}
	var param GetUserBaseInfoParam

	if err := c.ShouldBindQuery(&param); err != nil {
		Log.Errorf(err.Error())
		ret.SetErrCode(ERRCODE_PARAM, err.Error())
		return
	}
	rsp, d, err := rpc.InnerService.UserSevice.CallGetUserBaseInfo(param.Uid)
	if err != nil {
		ret.SetErrCode(ERRCODE_UNKNOWN, err.Error())
		return
	}
	ret.SetErrCode(rsp.Err, rsp.Message)
	ret.SetDataSection("data", d)
}

func (s *ActionGroup) GetUserRealName(c *gin.Context) {
	ret := NewPublciError()
	defer func() {
		c.JSON(http.StatusOK, ret.GetResult())
	}()
	type GetUserBaseInfoParam struct {
		Uid uint64 `form:"uid" binding:"required"`
	}
	var param GetUserBaseInfoParam

	if err := c.ShouldBindQuery(&param); err != nil {
		Log.Errorf(err.Error())
		ret.SetErrCode(ERRCODE_PARAM, err.Error())
		return
	}
	rsp, d, err := rpc.InnerService.UserSevice.CallGetUserRealName(param.Uid)
	if err != nil {
		ret.SetErrCode(ERRCODE_UNKNOWN, err.Error())
		return
	}
	ret.SetErrCode(rsp.Err, rsp.Message)
	ret.SetDataSection("data", d)
}

func (s *ActionGroup) GetUserInvite(c *gin.Context) {
	ret := NewPublciError()
	defer func() {
		c.JSON(http.StatusOK, ret.GetResult())
	}()
	type GetUserBaseInfoParam struct {
		Uid uint64 `form:"uid" binding:"required"`
	}
	var param GetUserBaseInfoParam

	if err := c.ShouldBindQuery(&param); err != nil {
		Log.Errorf(err.Error())
		ret.SetErrCode(ERRCODE_PARAM, err.Error())
		return
	}
	rsp, d, err := rpc.InnerService.UserSevice.CallGetUserInvite(param.Uid)
	if err != nil {
		ret.SetErrCode(ERRCODE_UNKNOWN, err.Error())
		return
	}
	ret.SetErrCode(rsp.Err, rsp.Message)
	ret.SetDataSection("data", d)
}

func (s *ActionGroup) GetIpRecord(c *gin.Context) {
	ret := NewPublciError()
	defer func() {
		c.JSON(http.StatusOK, ret.GetResult())
	}()
	type GetUserBaseInfoParam struct {
		Uid   uint64 `form:"uid" binding:"required"`
		Limit int32  `form:"limit" `
		Page  int32  `form:"page" `
	}
	var param GetUserBaseInfoParam

	if err := c.ShouldBindQuery(&param); err != nil {
		Log.Errorf(err.Error())
		ret.SetErrCode(ERRCODE_PARAM, err.Error())
		return
	}

	if param.Limit == 0 {
		param.Limit = 5
	}
	if param.Page == 0 {
		param.Page = 1
	}

	rsp, err := rpc.InnerService.UserSevice.CallGetIpRecord(param.Uid)
	if err != nil {
		ret.SetErrCode(ERRCODE_UNKNOWN, err.Error())
		return
	}
	ret.SetErrCode(rsp.Err, rsp.Message)
	ret.SetDataSection("list", rsp.Data)
}
