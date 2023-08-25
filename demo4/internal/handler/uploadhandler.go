package handler

import (
	"net/http"

	"demo4/internal/logic"
	"demo4/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UploadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		logx.Infof("进入uploadHandler")
		l := logic.NewUploadLogic(r, svcCtx)
		resp, err := l.Upload()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp) 
			//以向客户端发送 JSON 数据作为成功的响应。该函数接受三个参数：
			//r.Context() 表示 HTTP 请求的上下文，用于提供额外的请求信息或取消请求。
			//w 是一个 http.ResponseWriter 对象，表示 HTTP 响应的写入器，用于将数据写入响应。
			//resp 是一个表示要发送给客户端的 JSON 数据。
		}

	}
}
