package user

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zlei/app/rest/user/internal/logic/user"
	"zlei/app/rest/user/internal/svc"
	"zlei/app/rest/user/internal/types"
)

func WxMiniAuthHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.WXMiniAuthReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := user.NewWxMiniAuthLogic(r.Context(), svcCtx)
		resp, err := l.WxMiniAuth(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
