package logic

import (
	"apiTools/internal/svc"
	"apiTools/internal/types"
	"context"
	"reflect"
	"testing"

	"github.com/zeromicro/go-zero/core/logx"
)

func TestQueryDTlistLogic_QueryDTlist(t *testing.T) {
	type fields struct {
		Logger logx.Logger
		ctx    context.Context
		svcCtx *svc.ServiceContext
	}
	type args struct {
		req *types.QueryDT_dailyTradeReq
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantResp *types.QueryDT_dailyTradeResp
		wantErr  bool
	}{
		// TODO: Add test cases.
		{
			name: "TestQueryDTlistLogic_QueryDTlist",
			fields: fields{
				Logger: logx.WithContext(context.Background()),
				ctx:    context.Background(),
				svcCtx: nil,
			},
			args: args{
				req: &types.QueryDT_dailyTradeReq{
					StartDate: "2020-01-01",
					EndDate:   "2020-01-02",
					Domain:    "1",
				},
			},
			wantResp: nil,
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &QueryDTlistLogic{
				Logger: tt.fields.Logger,
				ctx:    tt.fields.ctx,
				svcCtx: tt.fields.svcCtx,
			}
			gotResp, err := l.QueryDTlist(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("QueryDTlistLogic.QueryDTlist() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("QueryDTlistLogic.QueryDTlist() = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}
