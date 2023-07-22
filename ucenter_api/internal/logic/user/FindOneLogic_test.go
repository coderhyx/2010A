package user

import (
	common "2010A/common/ctl"
	"context"
	"reflect"
	"testing"
	"ucenter_api/internal/svc"
	"ucenter_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

func TestFindOneLogic_FindOne(t *testing.T) {
	type fields struct {
		Logger logx.Logger
		ctx    context.Context
		svcCtx *svc.ServiceContext
	}
	type args struct {
		req *types.FindOneRequest
	}
	//填充下面的测试用例
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantRes *common.Result
		wantErr bool
	}{
		{
			name: "根据id查找用户",
			fields: fields{
				Logger: logx.WithContext(context.Background()),
				ctx:    context.Background(),
				svcCtx: &svc.ServiceContext{},
			},
			args: args{
				req: &types.FindOneRequest{
					MemberId: 2,
				},
			},
			wantRes: &common.Result{
				Code:    0,
				Message: "success",
				Data: &types.FindOneResponse{
					Username:     "小米",
					Password:     "cd02b1e2eee0664d97e47fe897fd5d955ba1651496434c86368ed4017ac4f4bea8565256715767dfc1fedb273b0d7213a4096f685b09adc34b5f65d93a6132c29ce8cdf8687cf5f46472bedc05eacac995e4072f705fc6370bc8041d76c245e6d647b5cf2d9601b47aad308055bfdbd5495d7ea4af8443b5ebf500fd1d4a9ecd",
					Captcha:      nil,
					Phone:        "",
					Promotion:    "",
					Code:         "",
					Country:      "1",
					SuperPartner: "11",
					Ip:           "",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &FindOneLogic{
				Logger: tt.fields.Logger,
				ctx:    tt.fields.ctx,
				svcCtx: tt.fields.svcCtx,
			}
			gotRes, err := l.FindOne(tt.args.req)
			//fmt.Println("=============>", gotRes)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindOne() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("FindOne() gotRes = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

func TestFindOneLogic_FindOne1(t *testing.T) {
	type fields struct {
		Logger logx.Logger
		ctx    context.Context
		svcCtx *svc.ServiceContext
	}
	type args struct {
		req *types.FindOneRequest
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantRes *common.Result
		wantErr bool
	}{ //生成测试用例
		{
			name: "Valid member ID",
			fields: fields{
				Logger: logx.WithContext(context.Background()),
				ctx:    context.Background(),
				svcCtx: &svc.ServiceContext{},
			},
			args: args{
				req: &types.FindOneRequest{
					MemberId: 2,
				},
			},
			wantRes: &common.Result{
				Code:    0,
				Message: "success",
				Data: &types.FindOneResponse{
					Username:     "小米",
					Password:     "cd02b1e2eee0664d97e47fe897fd5d955ba1651496434c86368ed4017ac4f4bea8565256715767dfc1fedb273b0d7213a4096f685b09adc34b5f65d93a6132c29ce8cdf8687cf5f46472bedc05eacac995e4072f705fc6370bc8041d76c245e6d647b5cf2d9601b47aad308055bfdbd5495d7ea4af8443b5ebf500fd1d4a9ecd",
					Captcha:      nil,
					Phone:        "",
					Promotion:    "",
					Code:         "",
					Country:      "1",
					SuperPartner: "11",
					Ip:           "",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &FindOneLogic{
				Logger: tt.fields.Logger,
				ctx:    tt.fields.ctx,
				svcCtx: tt.fields.svcCtx,
			}
			gotRes, err := l.FindOne(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindOne() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("FindOne() gotRes = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
