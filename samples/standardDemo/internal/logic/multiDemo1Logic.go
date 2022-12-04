package logic

import (
	"context"

	"github.com/chenkaiwei/crypto-m/samples/standardDemo/internal/svc"
	"github.com/chenkaiwei/crypto-m/samples/standardDemo/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MultiDemo1Logic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMultiDemo1Logic(ctx context.Context, svcCtx *svc.ServiceContext) *MultiDemo1Logic {
	return &MultiDemo1Logic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MultiDemo1Logic) MultiDemo1(req *types.SimpleMsg) (resp *types.SimpleMsg, err error) {
	logx.Info("MultiDemo1 called")

	resp = &types.SimpleMsg{
		Message: "your message is revieved --" + req.Message,
	}

	return
}
