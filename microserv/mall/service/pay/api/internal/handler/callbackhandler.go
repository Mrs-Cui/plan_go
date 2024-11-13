package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"plan_go/microserv/mall/service/pay/api/internal/logic"
	"plan_go/microserv/mall/service/pay/api/internal/svc"
	"plan_go/microserv/mall/service/pay/api/internal/types"
)

func CallbackHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CallbackRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewCallbackLogic(r.Context(), svcCtx)
		resp, err := l.Callback(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
