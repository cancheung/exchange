package handler

import (
	. "digicon/common/constant"
	. "digicon/proto/common"
	proto "digicon/proto/rpc"
	//. "digicon/user_service/dao"
	"golang.org/x/net/context"

	"digicon/common/constant"
	. "digicon/user_service/log"
	"digicon/user_service/model"
	"time"

	"github.com/go-redis/redis"

	"fmt"

	"strconv"

	"github.com/gin-gonic/gin/json"
)

type RPCServer struct{}

func (s *RPCServer) Hello(ctx context.Context, req *proto.HelloRequest, rsp *proto.HelloResponse) error {
	Log.Print("Received Say.Hello request")
	rsp.Greeting = "Hello " + req.Name
	return nil
}

//注册
func (s *RPCServer) Register(ctx context.Context, req *proto.RegisterRequest, rsp *proto.CommonErrResponse) error {
	if req.Type == 1 { //手机注册
		/*
			r := model.RedisOp{}
			code, err := r.GetSmsCode(req.Ukey, model.SMS_REGISTER)
			if err == redis.Nil {
				rsp.Err = ERRCODE_SMS_CODE_NIL
				rsp.Message = GetErrorMessage(rsp.Err)
				return nil
			} else if err != nil {
				rsp.Err = ERRCODE_UNKNOWN
				rsp.Message = err.Error()
				return nil
			}

			if req.Code == code {
				u := &model.User{}
				rsp.Err = u.Register(req, "phone")
				rsp.Message = GetErrorMessage(rsp.Err)
				return nil
			}

			rsp.Err = ERRCODE_SMS_CODE_DIFF
			rsp.Message = GetErrorMessage(rsp.Err)

		*/
		ret, err := model.AuthSms(req.Ukey, model.SMS_REGISTER, req.Code)
		if err != nil {
			rsp.Err = ERRCODE_UNKNOWN
			rsp.Message = err.Error()
			return nil
		}
		if ret != ERRCODE_SUCCESS {
			rsp.Err = ret
			return nil
		}
		u := &model.User{}
		rsp.Err = u.Register(req, "phone")
		rsp.Message = GetErrorMessage(rsp.Err)
		return nil
	} else if req.Type == 2 {
		ret, err := model.AuthEmail(req.Ukey, model.SMS_REGISTER, req.Code)
		if err != nil {
			rsp.Err = ERRCODE_UNKNOWN
			rsp.Message = err.Error()
			return nil
		}
		if ret != ERRCODE_SUCCESS {
			rsp.Err = ret
			return nil
		}
		u := &model.User{}
		rsp.Err = u.Register(req, "email")
		rsp.Message = GetErrorMessage(rsp.Err)
		return nil
	}

	rsp.Err = ERRCODE_SMS_CODE_DIFF
	rsp.Message = GetErrorMessage(rsp.Err)
	return nil
}

//注册by email
func (s *RPCServer) RegisterByEmail(ctx context.Context, req *proto.RegisterEmailRequest, rsp *proto.CommonErrResponse) error {
	return nil
}

//登陆
func (s *RPCServer) Login(ctx context.Context, req *proto.LoginRequest, rsp *proto.LoginResponse) error {
	u := &model.User{}
	var token string
	var ret int32
	if req.Type == 1 { //手机登陆
		token, ret = u.LoginByPhone(req.Ukey, req.Pwd)
	} else if req.Type == 2 { //邮箱登陆
		token, ret = u.LoginByEmail(req.Ukey, req.Pwd)
	}

	if ret == ERRCODE_SUCCESS {
		new(model.LoginRecord).AddLoginRecord(u.Uid, req.Ip)

		var p proto.LoginUserBaseData
		u.GetLoginUser(&p)
		p.Token = []byte(token)
		rsp.Data = &p
	}
	rsp.Err = ret
	rsp.Message = GetErrorMessage(rsp.Err)
	return nil
}

//忘记密码
func (s *RPCServer) ForgetPwd(ctx context.Context, req *proto.ForgetRequest, rsp *proto.ForgetResponse) error {
	var ret int32
	var err error

	u := &model.User{}
	if req.Type == 1 { //电话找回
		ret, err = u.GetUserByPhone(req.Ukey)
	} else if req.Type == 2 {
		ret, err = u.GetUserByEmail(req.Ukey)
	} else {
		ret = ERRCODE_PARAM
		return nil
	}

	if err != nil {
		rsp.Err = ret
		rsp.Message = err.Error()
		return err
	}

	if ret != ERRCODE_SUCCESS {
		rsp.Err = ret
		return nil
	}
	ret, err = u.AuthCodeByAl(req.Ukey, req.Code, model.SMS_FORGET, false)
	if err == redis.Nil {
		rsp.Err = ERRCODE_SMS_CODE_NIL
		rsp.Message = GetErrorMessage(rsp.Err)
		return nil
	} else if err != nil {
		rsp.Err = ret
		rsp.Message = err.Error()
		return err
	}
	if ret != ERRCODE_SUCCESS {
		rsp.Err = ret
		return nil
	}

	err = u.ModifyPwd(req.Pwd)
	if err != nil {
		rsp.Err = ERRCODE_UNKNOWN
		rsp.Message = err.Error()
		return nil
	}
	rsp.Err = ERRCODE_SUCCESS
	return nil

	/*
		if req.Type == 1 { //电话找回
			ret, err := model.AuthSms(req.Ukey, req.Type, req.Code)
			if err != nil {
				rsp.Err = ERRCODE_UNKNOWN
				rsp.Message = err.Error()
				return nil
			}

			if ret != ERRCODE_SUCCESS {
				rsp.Err = ret
				rsp.Message = GetErrorMessage(ret)
			}

			u := model.User{}
			ret, err = u.GetUserByPhone(req.Ukey)
			if err != nil {
				rsp.Err = ret
				rsp.Message = err.Error()
				return err
			}

			if ret != ERRCODE_SUCCESS {
				rsp.Err = ret
				rsp.Message = GetErrorMessage(rsp.Err)
				return nil
			}
			err = u.ModifyPwd(req.Pwd)
			if err != nil {
				rsp.Err = ERRCODE_UNKNOWN
				rsp.Message = err.Error()
				return nil
			}
			rsp.Err = ERRCODE_SUCCESS
			rsp.Message = GetErrorMessage(rsp.Err)
			return nil

		} else if req.Type == 2 { //邮箱找回

		}
	*/
	rsp.Err = ERRCODE_PARAM
	rsp.Message = GetErrorMessage(rsp.Err)
	return nil
}

//安全认证
func (s *RPCServer) AuthSecurity(ctx context.Context, req *proto.SecurityRequest, rsp *proto.SecurityResponse) error {
	/*
		security_key, err := DB.GenSecurityKey(req.Phone)
		if err != nil {
			return nil
		}
		rsp.Err = ERRCODE_SUCCESS
		rsp.Message = GetErrorMessage(rsp.Err)
		rsp.SecurityKey = security_key
	*/
	return nil
}

//发生短信验证码
func (s *RPCServer) SendSms(ctx context.Context, req *proto.SmsRequest, rsp *proto.CommonErrResponse) error {
	ret, err := model.ProcessSmsLogic(req.Type, req.Phone, req.Region)
	if err != nil {
		rsp.Err = ret
		rsp.Message = err.Error()
		return nil
	}
	rsp.Err = ret
	rsp.Message = GetErrorMessage(rsp.Err)
	return nil
}

//发送邮箱验证码
func (s *RPCServer) SendEmail(ctx context.Context, req *proto.EmailRequest, rsp *proto.CommonErrResponse) error {
	ret, err := model.ProcessEmailLogic(req.Type, req.Email)
	if err != nil {
		rsp.Err = ret
		rsp.Message = err.Error()
		return nil
	}
	rsp.Err = ret
	rsp.Message = GetErrorMessage(rsp.Err)
	return nil
}

//改变密码
func (s *RPCServer) ChangePwd(ctx context.Context, req *proto.EmailRequest, rsp *proto.CommonErrResponse) error {
	/*
			security_key, err := DB.GetSecurityKeyByPhone(req.Phone)
			if err != nil {
				return nil
			}
			if string(security_key) == string(req.SecurityKey) {
				u := model.User{}
				ret := u.GetUserByPhone(req.Phone)
				if ret != ERRCODE_SUCCESS {
					rsp.Err = ret
					rsp.Message = GetErrorMessage(rsp.Err)
					return nil
				}

				err = u.ModifyPwd(req.Pwd)
				if err != nil {
					rsp.Err = ERRCODE_UNKNOWN
					rsp.Message = err.Error()
					return nil
				}
				rsp.Err = ERRCODE_SUCCESS
				rsp.Message = GetErrorMessage(rsp.Err)
			} else {
				rsp.Err = ERRCODE_SECURITY_KEY
				rsp.Message = GetErrorMessage(rsp.Err)
			}
	// 	*/
	return nil
}

//获取登陆记录
func (s *RPCServer) GetIpRecord(ctx context.Context, req *proto.CommonPageRequest, rsp *proto.IpRecordResponse) error {
	g := new(model.LoginRecord).GetLoginRecord(req.Uid, int(req.Page), int(req.Limit))
	for _, v := range g {
		rsp.Data = append(rsp.Data, &proto.IpRecordBaseData{
			Ip:          v.Ip,
			CreatedTime: time.Unix(v.CreatedTime, 0).Format("2006-01-02 15:04:05"),
		})
	}

	return nil
}

func (this *RPCServer) TokenList(ctx context.Context, req *proto.NullRequest, rsp *proto.TokenListResponse) error {
	g := new(model.Tokens).GetTokens()
	for _, v := range g {
		rsp.Data = append(rsp.Data, &proto.TokenMarkBaseData{
			TokenId: int32(v.Id),
			Mark:    v.Mark,
		})
	}
	return nil
}

func (this *RPCServer) CheckSecurity(ctx context.Context, req *proto.CheckSecurityRequest, rsp *proto.CheckSecurityResponse) error {
	u := &model.User{}
	var ret int32
	var err error
	var uid int64
	if req.Type == 3 {
		uid, err = strconv.ParseInt(req.Ukey, 10, 64)
		if err != nil {
			rsp.Err = ERRCODE_UNKNOWN
			rsp.Message = err.Error()
			return nil
		}
		ret, err = u.GetUser(uint64(uid))

	} else if req.Type == 2 {
		ret, err = u.GetUserByEmail(req.Ukey)

	} else if req.Type == 1 {
		ret, err = u.GetUserByPhone(req.Ukey)
	} else {
		rsp.Err = ERRCODE_PARAM
		return nil
	}
	if err != nil {
		rsp.Err = ret
		rsp.Message = err.Error()
		return nil
	}

	if ret != ERRCODE_SUCCESS {
		rsp.Err = ret
		return nil
	}

	if req.Type == 3 {
		rsp.Auth = u.GetAuthMethodExpectGoogle()
	} else {
		rsp.Auth = u.GetAuthMethod()
	}
	//rsp.Auth = u.GetAuthMethod()
	if rsp.Auth == constant.AUTH_PHONE {
		rsp.Region = u.Country
	}
	return nil
}

/*
	// bind user email
*/
func (this *RPCServer) BindEmail(ctx context.Context, req *proto.BindEmailRequest, rsp *proto.BindPhoneEmailResponse) error {
	u := new(model.User)
	u.GetUser(req.Uid)
	phone := u.Phone
	var err error
	code, err := model.AuthEmail(req.Email, model.SMS_BIND_EMAIL, req.EmailCode)
	fmt.Println("code:", code, err)

	if err != nil {
		Log.Errorln("auth code by email error!")
		rsp.Code = ERRCODE_UNKNOWN
		return err
	}
	if req.VerifyType == 1 { // 3: 短信校验
		rsp.Code, err = model.AuthSms(phone, model.SMS_BIND_EMAIL, req.VerifyCode)
		fmt.Println(rsp.Code, err)
		//rsp.Code, err = u.AuthCodeByAl(phone, req.VerifyCode, model.SMS_BIND_EMAIL)
		if err != nil {
			rsp.Code = ERRCODE_UNKNOWN
			return err
		}
	} else if req.VerifyType == 2 { // 4 谷歌验证
		rsp.Code, err = u.AuthCodeByAl(u.GoogleVerifyId, req.VerifyCode, model.SMS_BIND_EMAIL, false)
		if err != nil {
			return err
		}
	} else {
		Log.Errorln(" not found verifyType!")
		rsp.Code = ERRCODE_UNKNOWN
		return nil
	}
	has, err := u.BindUserEmail(req.Email, req.Uid)

	if err != nil {
		Log.Errorln("bind user email err!", err.Error())
		rsp.Code = ERRCODE_UNKNOWN
		return nil
	}
	if has {
		rsp.Code = ERRCODE_EMAIL_EXIST
		rsp.Message = "邮箱已经存在"
		return nil
	}

	err = u.SecurityChmod(AUTH_EMAIL)
	fmt.Println("security chmod :", err)
	if err != nil {
		msg := "after bind user email, security chmod error!"
		Log.Errorln(msg)
		rsp.Code = ERRCODE_UNKNOWN
		rsp.Message = msg
	}
	u.RefreshCache(req.Uid)
	rsp.Code = ERRCODE_SUCCESS
	return nil
}

func (this *RPCServer) BindPhone(ctx context.Context, req *proto.BindPhoneRequest, rsp *proto.BindPhoneEmailResponse) error {
	u := new(model.User)
	u.GetUser(req.Uid)
	email := u.Email
	var err error
	//rsp.Code, err = u.AuthCodeByAl(req.Phone, req.PhoneCode, model.SMS_BIND_PHONE)
	rsp.Code, err = model.AuthSms(req.Phone, model.SMS_BIND_PHONE, req.PhoneCode)

	if err != nil {
		rsp.Code = ERRCODE_UNKNOWN
		return err
	}
	if req.VerifyType == 1 { //  1. email verify
		//rsp.Code, err = u.AuthCodeByAl(phone, req.VerifyCode, model.SMS_BIND_PHONE)
		rsp.Code, err = model.AuthEmail(email, model.SMS_BIND_PHONE, req.VerifyCode)
		if err != nil {
			rsp.Code = ERRCODE_UNKNOWN
			return err
		}
	} else if req.VerifyType == 2 { // 2. google verify
		rsp.Code, err = u.AuthCodeByAl(u.GoogleVerifyId, req.VerifyCode, model.SMS_BIND_PHONE, false)
		if err != nil {
			rsp.Code = ERRCODE_UNKNOWN
			return err
		}
	} else {
		Log.Errorln(" not found verifyType!")
		rsp.Code = ERRCODE_UNKNOWN
		return nil
	}
	has, err := u.BindUserPhone(req.Phone, req.Uid)
	if err != nil {
		Log.Errorln("bind user phone err!", err.Error())
		rsp.Code = ERRCODE_UNKNOWN
		return nil
	}
	if has {
		rsp.Code = ERRCODE_PHONE_EXIST
		rsp.Message = "电话已经存在"
		return nil
	}
	err = u.SecurityChmod(AUTH_PHONE)
	if err != nil {
		msg := "after bind user phone, security chmod error!"
		Log.Errorln(msg)
		rsp.Code = ERRCODE_UNKNOWN
		rsp.Message = msg
		return nil
	}
	u.RefreshCache(req.Uid)
	rsp.Code = ERRCODE_SUCCESS
	return nil
}

/*
	获取认证信息
*/

func (this *RPCServer) GetAuthInfo(ctx context.Context, req *proto.GetAuthInfoRequest, rsp *proto.GetAuthInfoResponse) error {
	u := new(model.User)

	code, err := u.GetUser(req.Uid)
	if err != nil {
		fmt.Println(err)
		rsp.Data = ""
		rsp.Code = code
		return err
	}
	//fmt.Println("uid:", req.Uid)
	securityCode := u.SecurityAuth

	type AuthInfo struct {
		EmailAuth    int32 `json:"email_auth"`     //
		PhoneAuth    int32 `json:"phone_auth"`     //
		RealName     int32 `json:"real_name"`      //
		TwoLevelAuth int32 `json:"two_level_auth"` //
	}
	authInfo := new(AuthInfo)
	if securityCode-(securityCode^constant.AUTH_PHONE) == constant.AUTH_PHONE {
		authInfo.PhoneAuth = 1
	}
	if securityCode-(securityCode^constant.AUTH_EMAIL) == constant.AUTH_EMAIL {
		authInfo.EmailAuth = 1
	}
	if securityCode-(securityCode^constant.AUTH_TWO) == constant.AUTH_TWO {
		authInfo.TwoLevelAuth = 1
	}
	if securityCode-(securityCode^constant.AUTH_FIRST) == constant.AUTH_FIRST {
		authInfo.RealName = 1
	}
	data, err := json.Marshal(authInfo)
	if err != nil {
		fmt.Println(err.Error())
		Log.Errorln(err)
		rsp.Code = ERRCODE_UNKNOWN
		return err
	}
	rsp.Data = string(data)
	rsp.Code = ERRCODE_SUCCESS
	return nil
}
