package logic

import (
	"context"

	"github.com/chenkaiwei/crypto-m/samples/simpleDemo/internal/svc"
	"github.com/chenkaiwei/crypto-m/samples/simpleDemo/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CryptionTestLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCryptionTestLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CryptionTestLogic {
	return &CryptionTestLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CryptionTestLogic) CryptionTest(req *types.SimpleMsg) (resp *types.SimpleMsg, err error) {
	resp = &types.SimpleMsg{
		Message: "加密信息已收到，内容为--" + req.Message,
	}
	return
}
