// Code generated by goctl. DO NOT EDIT.
// Source: ucenter.proto

package ucenterclient

import (
	"context"

	"ucenter_srv/pb/ucenter"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AssetReq     = ucenter.AssetReq
	CaptchaReq   = ucenter.CaptchaReq
	CodeReq      = ucenter.CodeReq
	Coin         = ucenter.Coin
	ErrorDemoReq = ucenter.ErrorDemoReq
	ErrorDemoRes = ucenter.ErrorDemoRes
	LoginReq     = ucenter.LoginReq
	LoginRes     = ucenter.LoginRes
	MemberInfo   = ucenter.MemberInfo
	MemberReq    = ucenter.MemberReq
	MemberWallet = ucenter.MemberWallet
	MembersReq   = ucenter.MembersReq
	MembersRes   = ucenter.MembersRes
	NoRes        = ucenter.NoRes
	RegReq       = ucenter.RegReq
	RegRes       = ucenter.RegRes

	Ucenter interface {
		// 注册
		RegisterByPhone(ctx context.Context, in *RegReq, opts ...grpc.CallOption) (*RegRes, error)
		SendCode(ctx context.Context, in *CodeReq, opts ...grpc.CallOption) (*NoRes, error)
		// 会员查找
		FindMemberById(ctx context.Context, in *MemberReq, opts ...grpc.CallOption) (*MemberInfo, error)
		// 会员列表
		FindMembers(ctx context.Context, in *MembersReq, opts ...grpc.CallOption) (*MembersRes, error)
		// 登录
		Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginRes, error)
		FindWalletBySymbol(ctx context.Context, in *AssetReq, opts ...grpc.CallOption) (*MemberWallet, error)
		// error demo
		ErrorDemo(ctx context.Context, in *ErrorDemoReq, opts ...grpc.CallOption) (*ErrorDemoRes, error)
	}

	defaultUcenter struct {
		cli zrpc.Client
	}
)

func NewUcenter(cli zrpc.Client) Ucenter {
	return &defaultUcenter{
		cli: cli,
	}
}

// 注册
func (m *defaultUcenter) RegisterByPhone(ctx context.Context, in *RegReq, opts ...grpc.CallOption) (*RegRes, error) {
	client := ucenter.NewUcenterClient(m.cli.Conn())
	return client.RegisterByPhone(ctx, in, opts...)
}

func (m *defaultUcenter) SendCode(ctx context.Context, in *CodeReq, opts ...grpc.CallOption) (*NoRes, error) {
	client := ucenter.NewUcenterClient(m.cli.Conn())
	return client.SendCode(ctx, in, opts...)
}

// 会员查找
func (m *defaultUcenter) FindMemberById(ctx context.Context, in *MemberReq, opts ...grpc.CallOption) (*MemberInfo, error) {
	client := ucenter.NewUcenterClient(m.cli.Conn())
	return client.FindMemberById(ctx, in, opts...)
}

// 会员列表
func (m *defaultUcenter) FindMembers(ctx context.Context, in *MembersReq, opts ...grpc.CallOption) (*MembersRes, error) {
	client := ucenter.NewUcenterClient(m.cli.Conn())
	return client.FindMembers(ctx, in, opts...)
}

// 登录
func (m *defaultUcenter) Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginRes, error) {
	client := ucenter.NewUcenterClient(m.cli.Conn())
	return client.Login(ctx, in, opts...)
}

func (m *defaultUcenter) FindWalletBySymbol(ctx context.Context, in *AssetReq, opts ...grpc.CallOption) (*MemberWallet, error) {
	client := ucenter.NewUcenterClient(m.cli.Conn())
	return client.FindWalletBySymbol(ctx, in, opts...)
}

// error demo
func (m *defaultUcenter) ErrorDemo(ctx context.Context, in *ErrorDemoReq, opts ...grpc.CallOption) (*ErrorDemoRes, error) {
	client := ucenter.NewUcenterClient(m.cli.Conn())
	return client.ErrorDemo(ctx, in, opts...)
}
