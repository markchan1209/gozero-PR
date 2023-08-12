package logic

import (
	"context"

	"apiTools/internal/svc"
	"apiTools/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryDTlistLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryDTlistLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryDTlistLogic {
	return &QueryDTlistLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryDTlistLogic) QueryDTlist(req *types.QueryDT_dailyTradeReq) (resp *types.QueryDT_dailyTradeResp, err error) {
	// todo: add your logic here and delete this line

	return
}
