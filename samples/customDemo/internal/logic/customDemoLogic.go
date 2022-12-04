package logic

import (
	"context"

	"github.com/chenkaiwei/crypto-m/samples/customDemo/internal/svc"
	"github.com/chenkaiwei/crypto-m/samples/customDemo/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CustomDemoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCustomDemoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CustomDemoLogic {
	return &CustomDemoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CustomDemoLogic) CustomDemo(req *types.SimpleMsg) (resp *types.SimpleMsg, err error) {
	resp = &types.SimpleMsg{
		Message: "已收到加密信息--" + req.Message,
	}

	return
}
