package logic

import (
	"context"

	"demo3/internal/svc"
	"demo3/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type Demo3Logic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDemo3Logic(ctx context.Context, svcCtx *svc.ServiceContext) *Demo3Logic {
	return &Demo3Logic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Demo3Logic) Demo3(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line
	resp = new(types.Response)
	resp.Message = req.Name
	return
}
