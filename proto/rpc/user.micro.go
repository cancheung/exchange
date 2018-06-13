// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: rpc/user.proto

package g2u

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "context"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for UserRPC service

type UserRPCService interface {
	Hello(ctx context.Context, in *HelloRequest, opts ...client.CallOption) (*HelloResponse, error)
	RegisterByPhone(ctx context.Context, in *RegisterPhoneRequest, opts ...client.CallOption) (*CommonErrResponse, error)
	RegisterByEmail(ctx context.Context, in *RegisterEmailRequest, opts ...client.CallOption) (*CommonErrResponse, error)
	Login(ctx context.Context, in *LoginRequest, opts ...client.CallOption) (*LoginResponse, error)
	ForgetPwd(ctx context.Context, in *ForgetRequest, opts ...client.CallOption) (*ForgetResponse, error)
	AuthSecurity(ctx context.Context, in *SecurityRequest, opts ...client.CallOption) (*SecurityResponse, error)
	ChangePwd(ctx context.Context, in *ChangePwdRequest, opts ...client.CallOption) (*CommonErrResponse, error)
	SendSms(ctx context.Context, in *SmsRequest, opts ...client.CallOption) (*CommonErrResponse, error)
	SendEmail(ctx context.Context, in *EmailRequest, opts ...client.CallOption) (*CommonErrResponse, error)
	NoticeList(ctx context.Context, in *NoticeListRequest, opts ...client.CallOption) (*NoticeListResponse, error)
	NoticeDetail(ctx context.Context, in *NoticeDetailRequest, opts ...client.CallOption) (*NoticeDetailResponse, error)
}

type userRPCService struct {
	c    client.Client
	name string
}

func NewUserRPCService(name string, c client.Client) UserRPCService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "g2u"
	}
	return &userRPCService{
		c:    c,
		name: name,
	}
}

func (c *userRPCService) Hello(ctx context.Context, in *HelloRequest, opts ...client.CallOption) (*HelloResponse, error) {
	req := c.c.NewRequest(c.name, "UserRPC.Hello", in)
	out := new(HelloResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userRPCService) RegisterByPhone(ctx context.Context, in *RegisterPhoneRequest, opts ...client.CallOption) (*CommonErrResponse, error) {
	req := c.c.NewRequest(c.name, "UserRPC.RegisterByPhone", in)
	out := new(CommonErrResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userRPCService) RegisterByEmail(ctx context.Context, in *RegisterEmailRequest, opts ...client.CallOption) (*CommonErrResponse, error) {
	req := c.c.NewRequest(c.name, "UserRPC.RegisterByEmail", in)
	out := new(CommonErrResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userRPCService) Login(ctx context.Context, in *LoginRequest, opts ...client.CallOption) (*LoginResponse, error) {
	req := c.c.NewRequest(c.name, "UserRPC.Login", in)
	out := new(LoginResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userRPCService) ForgetPwd(ctx context.Context, in *ForgetRequest, opts ...client.CallOption) (*ForgetResponse, error) {
	req := c.c.NewRequest(c.name, "UserRPC.ForgetPwd", in)
	out := new(ForgetResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userRPCService) AuthSecurity(ctx context.Context, in *SecurityRequest, opts ...client.CallOption) (*SecurityResponse, error) {
	req := c.c.NewRequest(c.name, "UserRPC.AuthSecurity", in)
	out := new(SecurityResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userRPCService) ChangePwd(ctx context.Context, in *ChangePwdRequest, opts ...client.CallOption) (*CommonErrResponse, error) {
	req := c.c.NewRequest(c.name, "UserRPC.ChangePwd", in)
	out := new(CommonErrResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userRPCService) SendSms(ctx context.Context, in *SmsRequest, opts ...client.CallOption) (*CommonErrResponse, error) {
	req := c.c.NewRequest(c.name, "UserRPC.SendSms", in)
	out := new(CommonErrResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userRPCService) SendEmail(ctx context.Context, in *EmailRequest, opts ...client.CallOption) (*CommonErrResponse, error) {
	req := c.c.NewRequest(c.name, "UserRPC.SendEmail", in)
	out := new(CommonErrResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userRPCService) NoticeList(ctx context.Context, in *NoticeListRequest, opts ...client.CallOption) (*NoticeListResponse, error) {
	req := c.c.NewRequest(c.name, "UserRPC.NoticeList", in)
	out := new(NoticeListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userRPCService) NoticeDetail(ctx context.Context, in *NoticeDetailRequest, opts ...client.CallOption) (*NoticeDetailResponse, error) {
	req := c.c.NewRequest(c.name, "UserRPC.NoticeDetail", in)
	out := new(NoticeDetailResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for UserRPC service

type UserRPCHandler interface {
	Hello(context.Context, *HelloRequest, *HelloResponse) error
	RegisterByPhone(context.Context, *RegisterPhoneRequest, *CommonErrResponse) error
	RegisterByEmail(context.Context, *RegisterEmailRequest, *CommonErrResponse) error
	Login(context.Context, *LoginRequest, *LoginResponse) error
	ForgetPwd(context.Context, *ForgetRequest, *ForgetResponse) error
	AuthSecurity(context.Context, *SecurityRequest, *SecurityResponse) error
	ChangePwd(context.Context, *ChangePwdRequest, *CommonErrResponse) error
	SendSms(context.Context, *SmsRequest, *CommonErrResponse) error
	SendEmail(context.Context, *EmailRequest, *CommonErrResponse) error
	NoticeList(context.Context, *NoticeListRequest, *NoticeListResponse) error
	NoticeDetail(context.Context, *NoticeDetailRequest, *NoticeDetailResponse) error
}

func RegisterUserRPCHandler(s server.Server, hdlr UserRPCHandler, opts ...server.HandlerOption) {
	type userRPC interface {
		Hello(ctx context.Context, in *HelloRequest, out *HelloResponse) error
		RegisterByPhone(ctx context.Context, in *RegisterPhoneRequest, out *CommonErrResponse) error
		RegisterByEmail(ctx context.Context, in *RegisterEmailRequest, out *CommonErrResponse) error
		Login(ctx context.Context, in *LoginRequest, out *LoginResponse) error
		ForgetPwd(ctx context.Context, in *ForgetRequest, out *ForgetResponse) error
		AuthSecurity(ctx context.Context, in *SecurityRequest, out *SecurityResponse) error
		ChangePwd(ctx context.Context, in *ChangePwdRequest, out *CommonErrResponse) error
		SendSms(ctx context.Context, in *SmsRequest, out *CommonErrResponse) error
		SendEmail(ctx context.Context, in *EmailRequest, out *CommonErrResponse) error
		NoticeList(ctx context.Context, in *NoticeListRequest, out *NoticeListResponse) error
		NoticeDetail(ctx context.Context, in *NoticeDetailRequest, out *NoticeDetailResponse) error
	}
	type UserRPC struct {
		userRPC
	}
	h := &userRPCHandler{hdlr}
	s.Handle(s.NewHandler(&UserRPC{h}, opts...))
}

type userRPCHandler struct {
	UserRPCHandler
}

func (h *userRPCHandler) Hello(ctx context.Context, in *HelloRequest, out *HelloResponse) error {
	return h.UserRPCHandler.Hello(ctx, in, out)
}

func (h *userRPCHandler) RegisterByPhone(ctx context.Context, in *RegisterPhoneRequest, out *CommonErrResponse) error {
	return h.UserRPCHandler.RegisterByPhone(ctx, in, out)
}

func (h *userRPCHandler) RegisterByEmail(ctx context.Context, in *RegisterEmailRequest, out *CommonErrResponse) error {
	return h.UserRPCHandler.RegisterByEmail(ctx, in, out)
}

func (h *userRPCHandler) Login(ctx context.Context, in *LoginRequest, out *LoginResponse) error {
	return h.UserRPCHandler.Login(ctx, in, out)
}

func (h *userRPCHandler) ForgetPwd(ctx context.Context, in *ForgetRequest, out *ForgetResponse) error {
	return h.UserRPCHandler.ForgetPwd(ctx, in, out)
}

func (h *userRPCHandler) AuthSecurity(ctx context.Context, in *SecurityRequest, out *SecurityResponse) error {
	return h.UserRPCHandler.AuthSecurity(ctx, in, out)
}

func (h *userRPCHandler) ChangePwd(ctx context.Context, in *ChangePwdRequest, out *CommonErrResponse) error {
	return h.UserRPCHandler.ChangePwd(ctx, in, out)
}

func (h *userRPCHandler) SendSms(ctx context.Context, in *SmsRequest, out *CommonErrResponse) error {
	return h.UserRPCHandler.SendSms(ctx, in, out)
}

func (h *userRPCHandler) SendEmail(ctx context.Context, in *EmailRequest, out *CommonErrResponse) error {
	return h.UserRPCHandler.SendEmail(ctx, in, out)
}

func (h *userRPCHandler) NoticeList(ctx context.Context, in *NoticeListRequest, out *NoticeListResponse) error {
	return h.UserRPCHandler.NoticeList(ctx, in, out)
}

func (h *userRPCHandler) NoticeDetail(ctx context.Context, in *NoticeDetailRequest, out *NoticeDetailResponse) error {
	return h.UserRPCHandler.NoticeDetail(ctx, in, out)
}
