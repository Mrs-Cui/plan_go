package handler

import (
	"net/http"
	logic2 "plan_go/plan_go/microserv/mall/service/user/api/internal/logic"
	svc2 "plan_go/plan_go/microserv/mall/service/user/api/internal/svc"
	types2 "plan_go/plan_go/microserv/mall/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func LoginHandler(svcCtx *svc2.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types2.LoginRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic2.NewLoginLogic(r.Context(), svcCtx)
		resp, err := l.Login(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
