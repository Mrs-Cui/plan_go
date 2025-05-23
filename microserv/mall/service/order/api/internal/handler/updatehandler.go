package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"plan_go/microserv/mall/service/order/api/internal/logic"
	"plan_go/microserv/mall/service/order/api/internal/svc"
	"plan_go/microserv/mall/service/order/api/internal/types"
)

func UpdateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewUpdateLogic(r.Context(), svcCtx)
		resp, err := l.Update(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
