// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: rpc/currency.proto

/*
Package g2u is a generated protocol buffer package.

It is generated from these files:
	rpc/currency.proto
	rpc/token.proto
	rpc/user.proto
	rpc/waller.proto

It has these top-level messages:
	AdminRequest
	AdminResponse
	CommonErrResponse
	HelloRequest
	HelloResponse
	RegisterPhoneRequest
	RegisterEmailRequest
	LoginRequest
	LoginResponse
	ForgetRequest
	ForgetResponse
	SecurityRequest
	SecurityResponse
	SmsRequest
	EmailRequest
	ChangePwdRequest
	NoticeListRequest
	NoticeListResponse
	NoticeDetailRequest
	NoticeDetailResponse
	HelloRequest2
	HelloResponse2
*/
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

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for CurrencyRPC service

type CurrencyRPCService interface {
	AdminCmd(ctx context.Context, in *AdminRequest, opts ...client.CallOption) (*AdminResponse, error)
}

type currencyRPCService struct {
	c    client.Client
	name string
}

func NewCurrencyRPCService(name string, c client.Client) CurrencyRPCService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "g2u"
	}
	return &currencyRPCService{
		c:    c,
		name: name,
	}
}

func (c *currencyRPCService) AdminCmd(ctx context.Context, in *AdminRequest, opts ...client.CallOption) (*AdminResponse, error) {
	req := c.c.NewRequest(c.name, "CurrencyRPC.AdminCmd", in)
	out := new(AdminResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for CurrencyRPC service

type CurrencyRPCHandler interface {
	AdminCmd(context.Context, *AdminRequest, *AdminResponse) error
}

func RegisterCurrencyRPCHandler(s server.Server, hdlr CurrencyRPCHandler, opts ...server.HandlerOption) {
	type currencyRPC interface {
		AdminCmd(ctx context.Context, in *AdminRequest, out *AdminResponse) error
	}
	type CurrencyRPC struct {
		currencyRPC
	}
	h := &currencyRPCHandler{hdlr}
	s.Handle(s.NewHandler(&CurrencyRPC{h}, opts...))
}

type currencyRPCHandler struct {
	CurrencyRPCHandler
}

func (h *currencyRPCHandler) AdminCmd(ctx context.Context, in *AdminRequest, out *AdminResponse) error {
	return h.CurrencyRPCHandler.AdminCmd(ctx, in, out)
}
