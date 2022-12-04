package logic

import (
	"context"
	"errors"
	"github.com/chenkaiwei/crypto-m/cryptom"

	"github.com/chenkaiwei/crypto-m/samples/standardDemo/internal/svc"
	"github.com/chenkaiwei/crypto-m/samples/standardDemo/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ManualDemoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewManualDemoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ManualDemoLogic {
	return &ManualDemoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ManualDemoLogic) ManualDemo(req *types.StandardMsg) (resp *types.StandardMsg, err error) {

	cryptomManager := l.svcCtx.Multi1CryptomManager
	//手动解密
	decryptCustomMessage, err := cryptomManager.ContentDecrypt(l.ctx, req.EncryptedMessage)
	if err != nil {

		//复合型错误的判断示例
		var ce *cryptom.CryptomError
		if errors.As(err, &ce) && ce.ErrType() == cryptom.ErrTypeContentDecryptFailure {
			logx.Error("ErrType() == cryptom.ErrTypeContentDecryptFailure")
		}

		return nil, err
	}

	//局部加密
	encryptMsg, err := l.svcCtx.Multi1CryptomManager.ContentEncrypt(l.ctx, "来自服务端的加密消息4321")
	if err != nil {
		return nil, err
	}

	resp = &types.StandardMsg{
		Message:          "your message is revieved --" + req.Message,
		EncryptedMessage: encryptMsg,
		DecryptedMessage: "您发送的加密数据为：" + decryptCustomMessage,
	}

	return
}
