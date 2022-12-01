package cryptom

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
)

const ContentEncryptionKey = "CEK"

type CryptomManager interface {
	RequestHandle(next http.HandlerFunc) http.HandlerFunc
	ManualHandle(next http.HandlerFunc) http.HandlerFunc
	ResponseHandle(next http.HandlerFunc) http.HandlerFunc
	ContentDecrypt(ctx context.Context, encryptedMsg string) (res string, err error)
	ContentEncrypt(ctx context.Context, msg string) (res string, err error)
}

func getCekFromContext(ctx context.Context) (cek []byte, err error) {
	//
	value := ctx.Value(ContentEncryptionKey)

	if value == nil {
		err = ErrCEKNotFoundInContext
		return
	}
	cek = value.([]byte)
	logx.Info("成功获取已经解密的内容密钥--", cek)
	return
}
