package handler

import (
	"net/http"

	"demo4/internal/logic"
	"demo4/internal/svc"
	"demo4/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func DownloadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DownloadRequest
		if err := httpx.Parse(r, &req); err != nil { //httpx.Parse(r, &req) 是一个调用 httpx 包中的 Parse 函数，用于解析 HTTP 请求，并将解析结果存储到 req 变量中
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewDownloadLogic(r.Context(), svcCtx, w)
		err := l.Download(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
