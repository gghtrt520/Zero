package user

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zlei/app/rest/user/internal/logic/user"
	"zlei/app/rest/user/internal/svc"
	"zlei/app/rest/user/internal/types"
)

func RegisterHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RegisterReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := user.NewRegisterLogic(r.Context(), svcCtx)
		resp, err := l.Register(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
