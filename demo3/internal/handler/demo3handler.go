package handler

import (
	"net/http"

	"demo3/internal/logic"
	"demo3/internal/svc"
	"demo3/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func Demo3Handler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewDemo3Logic(r.Context(), svcCtx)
		resp, err := l.Demo3(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
