package logic

import (
	"apiTools/internal/svc"
	"apiTools/internal/types"
	"context"
	"reflect"
	"testing"

	"github.com/zeromicro/go-zero/core/logx"
)

func TestSendMsgLogic_SendMsg(t *testing.T) {
	type fields struct {
		Logger logx.Logger
		ctx    context.Context
		svcCtx *svc.ServiceContext
	}
	type args struct {
		req *types.Request
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantResp *types.Response
		wantErr  bool
	}{
		// TODO: Add test cases.
		{
			name: "TestSendMsgLogic_SendMsg",
			fields: fields{
				Logger: logx.WithContext(context.Background()),
				ctx:    context.Background(),
				svcCtx: nil,
			},
			args: args{
				req: &types.Request{
					Name: "1",
				},
			},
			wantResp: nil,
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &SendMsgLogic{
				Logger: tt.fields.Logger,
				ctx:    tt.fields.ctx,
				svcCtx: tt.fields.svcCtx,
			}
			gotResp, err := l.SendMsg(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("SendMsgLogic.SendMsg() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("SendMsgLogic.SendMsg() = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}
