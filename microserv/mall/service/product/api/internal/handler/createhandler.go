package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"plan_go/microserv/mall/service/product/api/internal/logic"
	"plan_go/microserv/mall/service/product/api/internal/svc"
	"plan_go/microserv/mall/service/product/api/internal/types"
)

func CreateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewCreateLogic(r.Context(), svcCtx)
		resp, err := l.Create(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
