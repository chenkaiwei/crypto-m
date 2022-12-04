package handler

import (
	"net/http"

	"github.com/chenkaiwei/crypto-m/samples/standardDemo/internal/logic"
	"github.com/chenkaiwei/crypto-m/samples/standardDemo/internal/svc"
	"github.com/chenkaiwei/crypto-m/samples/standardDemo/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func MultiDemo2Handler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SimpleMsg
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewMultiDemo2Logic(r.Context(), svcCtx)
		resp, err := l.MultiDemo2(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
