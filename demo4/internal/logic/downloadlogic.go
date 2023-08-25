package logic

import (
	"context"
	"io"
	"os"

	"demo4/internal/svc"
	"demo4/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DownloadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	writer io.Writer
}

func NewDownloadLogic(ctx context.Context, svcCtx *svc.ServiceContext, writer io.Writer) *DownloadLogic {
	return &DownloadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		writer: writer,
	}
}

func (l *DownloadLogic) Download(req *types.DownloadRequest) error {
	// todo: add your logic here and delete this line
	logx.Infof("download %s", req.File)
	body, err := os.ReadFile(req.File)
	if err != nil {
		return err
	}
	n, err := l.writer.Write(body) //body是读到的字节切片，n是写入的字节数
	if err != nil {
		return err
	}
	if n < len(body) {
		return io.ErrClosedPipe
	}
	return nil
}
